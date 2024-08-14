#!/bin/bash

OUT=${1:?"output path"}
shift

set -e

export BUILD_GOOS=${GOOS}
export BUILD_GOARCH=${GOARCH}

GOOS=${BUILD_GOOS} GOARCH=${BUILD_GOARCH} go build -o "${OUT}" "${@}"