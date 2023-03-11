#!/usr/bin/env bash



export PORT=$1
export SF_FILE=$2
export NUM_MIDI_CHANNELS=16

exec fluidsynth -a alsa -i -s \
    -o synth.ladspa.active=1 \
    -o "synth.effects-groups=$NUM_MIDI_CHANNELS" \
    -o "synth.midi-channels=$NUM_MIDI_CHANNELS" \
    -o "shell.port=$PORT" \
    $SF_FILE 
