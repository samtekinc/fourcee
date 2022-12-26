#!/bin/bash -e
BUILD_DIR="./build/"

mkdir -p $BUILD_DIR

BUILD_FOLDER="workflow-handler"
EXECUTABLE_NAME="workflow-handler"
FUNCTION_NAME="tfom-workflow-handler"

cd "$(dirname "$0")"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${BUILD_DIR}${EXECUTABLE_NAME}" "./lambda/workflow-handler"
cd "${BUILD_DIR}"
zip "${FUNCTION_NAME}.zip" "${EXECUTABLE_NAME}"
cd $OLDPWD

aws lambda update-function-code --function-name $FUNCTION_NAME --zip-file "fileb://${BUILD_DIR}${FUNCTION_NAME}.zip" --no-cli-pager