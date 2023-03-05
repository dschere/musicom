#!/usr/bin/env bash



export PORT=$1
export SF_FILE=$2

exec fluidsynth -a alsa -i -s \
    -o synth.ladspa.active=1 \
    -o "shell.port=$PORT" \
    $SF_FILE > /dev/null 2>&1
