#!/bin/bash

xterm -hold -e go run websockets/goServer.go &

cd webserv/POS

xterm -hold -e npm start &
sleep 2
echo $! > pgrep -f xterm
