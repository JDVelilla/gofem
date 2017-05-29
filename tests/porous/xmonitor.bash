#!/bin/bash

FILES="../../fem/*.go ../../ele/*.go ../*.go *.go"

while true; do
    inotifywait -q -e modify $FILES
    echo
    echo
    echo
    echo
    go test -test.run="upp01b"
done
