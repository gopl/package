#!/bin/bash

N=$1
if [[ -z "${N}" ]]; then
	N=10
fi

for(( i=0; i<${N}; i++ )); do
	date
	echo "stdout $i"
	echo "stderr $i" >&2
	sleep 1
done

exit 2
