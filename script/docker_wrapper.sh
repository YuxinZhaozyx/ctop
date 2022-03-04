docker_command=docker
ctop_command=ctop

if [ $1 == "run" ]
then
    $docker_command run -e CONTAINER_CREATED_USER=$(whoami) ${@:2}
elif [ $1 == "ctop" ]
then
    $ctop_command ${@:2}
else
    $docker_command $@
fi
