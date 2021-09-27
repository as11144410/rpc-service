#!/usr/bin/env bash

set -euo pipefail


API_SPEC_PATH='api/protobuf-spec'

if [ $# -ge 1 ] && [ -n "$1" ]
then
  eval "protoc -I . --go_out=plugins=grpc,paths=source_relative:. ${API_SPEC_PATH}/$1pb/$1.proto"
  echo "generate $1 api successed."
else
  eval "protoc -I . --go_out=plugins=grpc,paths=source_relative:. ${API_SPEC_PATH}/helloworld/helloworld.proto"


  eval 'sed -i "" -e "s/,omitempty//g" ${API_SPEC_PATH}/**/*.go'

  echo "generate all api successed."
fi

