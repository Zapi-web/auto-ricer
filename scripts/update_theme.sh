#!/bin/bash

if [ $# -eq 0 ]; then
	echo "Usage: update_theme.sh <path to picture>"
	exit 1
fi

if [ ! -f "$1" ]; then
	echo "Error: file "$1" not found"
	exit 2
fi

/usr/bin/pipx run pywal -i "$1"
