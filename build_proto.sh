#!/usr/bin/env bash

mkdir -p proto_gen/onebot

protoc -I onebot_idl --gofast_out=proto_gen/onebot onebot_idl/*.proto
