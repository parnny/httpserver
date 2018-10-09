#!/bin/bash
kill -9 `cat nohup-pid.txt`
rm -rf nohup-pid.txt
rm -rf nohup-out.log
nohup ./httpserver -config=config/server.toml > ./nohup-out.log 2>&1 & echo $! > nohup-pid.txt
