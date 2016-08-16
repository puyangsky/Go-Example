#!/bin/bash
for i in `seq 1 15`;
do 
	sh ./qps.sh &
done
