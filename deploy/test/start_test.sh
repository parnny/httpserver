#!/bin/bash
rm -rf nohup-pid.txt
rm -rf nohup-out.log
nohup python test.py > ./nohup-out.log 2>&1 & echo $! > nohup-pid.txt
