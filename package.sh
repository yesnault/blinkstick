#!/bin/bash

cd cli/blink
for GOOS in windows darwin linux freebsd; do
    for GOARCH in 386 amd64 arm; do
        if [[ $GOARCH == "arm" && $GOOS != "linux" ]]; then
          continue;
        fi;
        architecture="${GOOS}-${GOARCH}"
        echo "Building ${architecture} ${path}"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go build -ldflags "-X github.com/yesnault/blinkstick/cli/blink/update.architecture=${architecture}" -o bin/blink-${architecture}
    done
done
