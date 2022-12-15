#!/usr/bin/env bash

root="$(dirname "${BASH_SOURCE[0]}")"

source $(dirname $0)/resolve.sh

readonly YAML_OUTPUT_DIR="openshift/release/artifacts/"

readonly SERVING_CRD_YAML=${YAML_OUTPUT_DIR}/1-serving-crds.yaml
readonly SERVING_CORE_YAML=${YAML_OUTPUT_DIR}/2-serving-core.yaml
readonly SERVING_HPA_YAML=${YAML_OUTPUT_DIR}/3-serving-hpa.yaml
readonly SERVING_POST_INSTALL_JOBS_YAML=${YAML_OUTPUT_DIR}/4-serving-post-install-jobs.yaml

# Generate Knative component YAML files
resolve_resources "config/core/300-resources/ config/core/300-imagecache.yaml" "$SERVING_CRD_YAML"
resolve_resources "config/core/"                                               "$SERVING_CORE_YAML"
resolve_resources "config/hpa-autoscaling/"                                    "$SERVING_HPA_YAML"
resolve_resources "config/post-install/storage-version-migration.yaml"         "$SERVING_POST_INSTALL_JOBS_YAML"

# Generate a single yaml for the CI.
release=$1
output_file="openshift/release/knative-serving-${release}.yaml"
resolve_resources "$YAML_OUTPUT_DIR" "$output_file"
