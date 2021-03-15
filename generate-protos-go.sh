#!/usr/bin/env bash

set -ex

PROTOS="mlflow/protos"
OUT_DIR="clients/go/mlflow"

mkdir -p $OUT_DIR

protoc -I="$PROTOS" \
    --go_out=$OUT_DIR \
    --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR \
    --go-grpc_opt=paths=source_relative \
    "$PROTOS"/databricks.proto \
    "$PROTOS"/service.proto \
    "$PROTOS"/model_registry.proto \
    "$PROTOS"/databricks_artifacts.proto \
    "$PROTOS"/scalapb/scalapb.proto
