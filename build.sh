#!/usr/bin/env bash

go build -o musicom \
    duration.go \
    dynamic.go \
    effect.go \
    fluidsynth.go \
    measure.go  \
    midicodes.go \
    note.go \
    sequencer.go \
    song.go \
    timesig.go \
    track.go  \
    main.go
