#!/bin/bash
cd "`dirname $0`" || exit 1

readonly G_PROC_NAME='xagent'

function check()
{
    ps -ef | grep ${G_PROC_NAME} | grep -v 'grep'
    return $?
}

function help()
{
    echo "${0} <start|stop|restart|status>"
    exit 1
}

function kill_proc()
{
    all_possible_id=$(ps -eo pid,command | grep ${G_PROC_NAME} | grep -v 'grep' | awk '{print $1}')

    for proc_id in ${all_possible_id}
    do
        kill -INT ${proc_id}
    done
}

function start()
{
    nohup ./bin/${G_PROC_NAME} >/dev/null 2>&1 &
}


function stop()
{
    kill_proc
}

function restart()
{
    stop
    start
    return 0
}

function status()
{
    check
    if [ $? -eq 0 ]; then
        echo 'Running'
        return 0
    else
        echo 'Not running'
        return 1
    fi
}

case "${1}" in
start)
    start
    ;;
stop)
    stop
    ;;
restart)
    restart
    ;;
status)
    status
    ;;
*)
    help
    ;;
esac
