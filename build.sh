#!/usr/bin/env bash

go build -o misicom \
    duration.go \
    effect.go \
    measure.go  \
    midicodes.go \
    note.go \
    sequencer.go \
    song.go \
    timesig.go \
    track.go  
