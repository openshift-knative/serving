#!/usr/bin/env bash

root="$(dirname "${BASH_SOURCE[0]}")"

source $(dirname $0)/resolve.sh

release=$1
output_file="openshift/release/knative-serving-${release}.yaml"

resolve_resources "config/core/ config/hpa-autoscaling/" "$output_file"

if [[ "$release" != "ci" ]]; then
  # Drop the "knative-" suffix, which is added in upstream branch.
  # e.g. knative-v1.7.0 => v1.7.0
  release=${release#"knative-"}
  ${root}/download_release_artifacts.sh $release
fi
