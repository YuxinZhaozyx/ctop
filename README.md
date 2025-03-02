<p align="center"><img width="200px" src="/_docs/img/logo.png" alt="ctop"/></p>

#

![release][release] ![homebrew][homebrew] ![macports][macports]

Top-like interface for container metrics

`ctop` provides a concise and condensed overview of real-time metrics for multiple containers:
<p align="center"><img src="_docs/img/grid.gif" alt="ctop"/></p>

as well as a [single container view][single_view] for inspecting a specific container.

`ctop` comes with built-in support for Docker and runC; connectors for other container and cluster systems are planned for future releases.

This version add additional

## Build and Install (modified verison)

The modified version has the following additional features:

+ Support "container created user"

  + Automatically record the user who created the container and display in container-monitor menu

+ Process-level monitor

  + Monitor all processes created in containers. Press <kbd>TAB</kbd> to open process-monitor.

### Linux (Generic)

```shell
git clone https://github.com/YuxinZhaozyx/ctop.git ctop-src
cd ctop-src
docker run --rm -v $(pwd):/app quay.io/vektorcloud/go:1.15 sh -c "apk add --no-cache make && make build"

cp ./ctop /usr/local/bin/ctop
cp ./script/docker_wrapper.sh /usr/local/bin/docker_wrapper.sh
chmod a+x /usr/local/bin/ctop
chmod a+x /usr/local/bin/docker_wrapper.sh
echo "alias docker=\"docker_wrapper.sh\"" >> /etc/bash.bashrc
```

Restart terminal. Now the `ctop` will display the user creating the container at additional column `USER`. You can also use `docker ctop` to trigger ctop menu.


## Install (origin version)

Fetch the [latest release](https://github.com/bcicen/ctop/releases) for your platform:

#### Debian/Ubuntu

Maintained by a [third party](https://packages.azlux.fr/)
```bash
echo "deb http://packages.azlux.fr/debian/ buster main" | sudo tee /etc/apt/sources.list.d/azlux.list
wget -qO - https://azlux.fr/repo.gpg.key | sudo apt-key add -
sudo apt update
sudo apt install docker-ctop
```

#### Arch

`ctop` is available for Arch in the [AUR](https://aur.archlinux.org/packages/ctop-bin/)

#### Linux (Generic)

```bash
sudo wget https://github.com/bcicen/ctop/releases/download/0.7.6/ctop-0.7.6-linux-amd64 -O /usr/local/bin/ctop
sudo chmod +x /usr/local/bin/ctop
```

#### OS X

```bash
brew install ctop
```
or
```bash
sudo port install ctop
```
or
```bash
sudo curl -Lo /usr/local/bin/ctop https://github.com/bcicen/ctop/releases/download/0.7.6/ctop-0.7.6-darwin-amd64
sudo chmod +x /usr/local/bin/ctop
```

#### Docker

```bash
docker run --rm -ti \
  --name=ctop \
  --volume /var/run/docker.sock:/var/run/docker.sock:ro \
  quay.io/vektorlab/ctop:latest
```

## Building (origin version)

Build steps can be found [here][build].

## Usage

`ctop` requires no arguments and uses Docker host variables by default. See [connectors][connectors] for further configuration options.

### Config file

While running, use `S` to save the current filters, sort field, and other options to a default config path (`~/.config/ctop/config` on XDG systems, else `~/.ctop`).

Config file values will be loaded and applied the next time `ctop` is started.

### Options

Option | Description
--- | ---
`-a`	| show active containers only
`-f <string>` | set an initial filter string
`-h`	| display help dialog
`-i`  | invert default colors
`-r`	| reverse container sort order
`-s`  | select initial container sort field
`-v`	| output version information and exit

### Keybindings

|           Key            | Action                                                     |
| :----------------------: | ---------------------------------------------------------- |
| <kbd>&lt;ENTER&gt;</kbd> | Open container menu                                        |
|       <kbd>a</kbd>       | Toggle display of all (running and non-running) containers |
|       <kbd>f</kbd>       | Filter displayed containers (`esc` to clear when open)     |
|       <kbd>H</kbd>       | Toggle ctop header                                         |
|       <kbd>h</kbd>       | Open help dialog                                           |
|       <kbd>s</kbd>       | Select container sort field                                |
|       <kbd>r</kbd>       | Reverse container sort order                               |
|       <kbd>o</kbd>       | Open single view                                           |
|       <kbd>l</kbd>       | View container logs (`t` to toggle timestamp when open)    |
|       <kbd>e</kbd>       | Exec Shell                                                 |
|       <kbd>c</kbd>       | Configure columns                                          |
|       <kbd>S</kbd>       | Save current configuration to file                         |
|  <kbd>&lt;TAB&gt;</kbd>  | Open process-monitor menu                                  |
|       <kbd>q</kbd>       | Quit ctop                                                  |

[build]: _docs/build.md
[connectors]: _docs/connectors.md
[single_view]: _docs/single.md
[release]: https://img.shields.io/github/release/bcicen/ctop.svg "ctop"
[homebrew]: https://img.shields.io/homebrew/v/ctop.svg "ctop"
[macports]: https://repology.org/badge/version-for-repo/macports/ctop.svg?header=macports "ctop"

## Alternatives

See [Awesome Docker list](https://github.com/veggiemonk/awesome-docker/blob/master/README.md#terminal) for similar tools to work with Docker.
