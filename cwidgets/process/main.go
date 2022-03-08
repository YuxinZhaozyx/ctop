package process

import (
	"regexp"
	"sort"
	"strconv"
	"time"
	"fmt"

	gopsutil_cpu "github.com/shirou/gopsutil/v3/cpu"
	gopsutil_process "github.com/shirou/gopsutil/v3/process"
	"github.com/bcicen/ctop/logging"

	ui "github.com/gizak/termui"
	api "github.com/fsouza/go-dockerclient"
)

var (
	log       = logging.Init()
	sizeError = termSizeError()
	colWidth  = [2]int{65, 0}
	numShow   = 20
)




type ProcessManager struct {
	client   *api.Client
}

func NewProcessManager() (*ProcessManager, error) {
	client, err := api.NewClientFromEnv()
	if err != nil {
		return nil, err
	}

	pm := &ProcessManager{
		client:    client,
	}

	info, err := client.Info()
	if err != nil {
		return nil, err
	}

	log.Debugf("docker-connector ID: %s", info.ID)
	log.Debugf("docker-connector Driver: %s", info.Driver)
	log.Debugf("docker-connector Images: %d", info.Images)
	log.Debugf("docker-connector Name: %s", info.Name)
	log.Debugf("docker-connector ServerVersion: %s", info.ServerVersion)

	return pm, nil
}

var envPattern = regexp.MustCompile(`(?P<KEY>[^=]+)=(?P<VALUJE>.*)`)

func (pm *ProcessManager) GetProcessMetas() (ps []Meta) {
	activeContainers, err := pm.client.ListContainers(api.ListContainersOptions{
		All: false,
	})
	if err != nil {
		log.Errorf("%s (%T)", err.Error(), err)
		return
	}

	for _, container := range activeContainers {
		container_name := container.Names[0]

		insp, err := pm.client.InspectContainer(container.ID)
		if err != nil {
			continue
		}

		created_user := "undefined"
		for _, env := range insp.Config.Env {
			match := envPattern.FindStringSubmatch(env)
			key := match[1]
			value := match[2]
			if key == "CONTAINER_CREATED_USER" {
				created_user = value
				break
			}
		}

		topResult, err := pm.client.TopContainer(container.ID, "aux")
		if err != nil {
			continue
		}

		for _, data := range topResult.Processes {
			processMeta := NewMeta(
				"pid", data[1],
				"user", created_user,
				"name", container_name,
				"vsz", data[4],
				"rss", data[5],
				"cpu", data[2],
				"mem", data[3],
				"start", data[8],
				"command", data[10],
			)
			ps = append(ps, processMeta)
		}
	}

	// calculate realtime cpu usage
	numProcess := len(ps)
	numValidProcess := numProcess
	interval := time.Second / 2
	realtimeCpuUsageSuccess := true
	processes := make([]*gopsutil_process.Process, numProcess)
	startCpuTimesStats := make([]*gopsutil_cpu.TimesStat, numProcess)
	endCpuTimesStats := make([]*gopsutil_cpu.TimesStat, numProcess)
	cpuTimes := make([]float64, numProcess)
	for i, processMeta := range ps {
		pid, err := strconv.ParseInt(processMeta["pid"], 10, 32)
		if err != nil {
			continue
		}
		processes[i], err = gopsutil_process.NewProcess(int32(pid))
		if err != nil {
			numValidProcess--;
		}
	}
	if numValidProcess == 0 {
		realtimeCpuUsageSuccess = false
	}
	startTime := time.Now()
	for i, process := range processes {
		if process != nil {
			startCpuTimesStats[i], err = process.Times()
			if err != nil {
				processes[i] = nil
			}
		}
	}
	time.Sleep(interval)
	duration := time.Since(startTime).Seconds()
	if duration <= 0 {
		realtimeCpuUsageSuccess = false
	} else {
		for i, process := range processes {
			if process != nil {
				endCpuTimesStats[i], err = process.Times()
				if err != nil {
					processes[i] = nil
				}
			}
		}
		for i := 0; i < numProcess; i++ {
			if processes[i] != nil {
				cpuTimes[i] = 100 * (endCpuTimesStats[i].Total() - startCpuTimesStats[i].Total()) / duration
			}
		}
	}

	if realtimeCpuUsageSuccess {
		for i := 0; i < numProcess; i++ {
			ps[i]["cpu"] = fmt.Sprintf("%.2f", cpuTimes[i])
		}
	}


	sort.SliceStable(ps, func(i, j int) bool {
		i_cpu, err := strconv.ParseFloat(ps[i]["cpu"], 32)
		if err != nil {
			return false
		}

		j_cpu, err := strconv.ParseFloat(ps[j]["cpu"], 32)
		if err != nil {
			return false
		}

		if i_cpu != j_cpu {
			return i_cpu > j_cpu
		}

		i_mem, err := strconv.ParseFloat(ps[i]["mem"], 32)
		if err != nil {
			return false
		}

		j_mem, err := strconv.ParseFloat(ps[j]["mem"], 32)
		if err != nil {
			return false
		}

		return i_mem > j_mem
	})

	return ps
}

func termSizeError() *ui.Par {
	p := ui.NewPar("screen too small!")
	p.Height = 1
	p.Width = 20
	p.Border = false
	return p
}
