#!/bin/sh

APP_NAME="check-aws-ip-region"
PLATFORMS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm windows/amd64"

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_NAME=$APP_NAME'-'$GOOS'-'$GOARCH
  if [[ "${GOOS}" == "windows" ]]; then BIN_NAME+='.exe'; fi
  echo "Building ${BIN_NAME}"
  env GOOS=$GOOS GOARCH=$GOARCH go build -o ${BIN_NAME} .
  if [ $? -ne 0 ]; then
    echo 'An error has occurred! Aborting the script execution...'
    exit 1
  fi
done

