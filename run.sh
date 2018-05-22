#!/bin/bash

go run websockets/goServer.go &
cd webserv
npm start &

