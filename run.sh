#!/bin/bash

go run websockets/server.go &
cd webserv
npm start &

