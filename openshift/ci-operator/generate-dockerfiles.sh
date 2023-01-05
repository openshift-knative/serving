#!/usr/bin/env bash

set -x

function generate_dockerfiles() {
  local target_dir=$1; shift
  for img in $@; do
    local image_base=$(basename $img)
    mkdir -p $target_dir/$image_base
    bin=$image_base envsubst < openshift/ci-operator/Dockerfile.in > $target_dir/$image_base/Dockerfile
  done
}

generate_dockerfiles $@