#!/usr/bin/env bash

root="$(dirname "${BASH_SOURCE[0]}")"/../..

source $(dirname $0)/resolve.sh

release=$(yq r openshift/project.yaml project.tag)
release=${release/knative-/}

echo "Release: $release"

# Reconcile dependencies in case of dependabot updates
"${root}"/hack/update-deps.sh
# Re-apply patches that touch vendor/ dir
git apply "${root}"/openshift/patches/001-object.patch
git apply "${root}"/openshift/patches/002-mutemetrics.patch
git apply "${root}"/openshift/patches/003-routeretry.patch

./openshift/generate.sh

readonly YAML_OUTPUT_DIR="openshift/release/artifacts/"

# Clean up
rm -rf "$YAML_OUTPUT_DIR"
mkdir -p "$YAML_OUTPUT_DIR"

readonly SERVING_CRD_YAML=${YAML_OUTPUT_DIR}/serving-crds.yaml
readonly SERVING_CORE_YAML=${YAML_OUTPUT_DIR}/serving-core.yaml
readonly SERVING_HPA_YAML=${YAML_OUTPUT_DIR}/serving-hpa.yaml
readonly SERVING_POST_INSTALL_JOBS_YAML=${YAML_OUTPUT_DIR}/serving-post-install-jobs.yaml

# Generate Knative component YAML files
resolve_resources "config/core/300-resources/ config/core/300-imagecache.yaml" "$SERVING_CRD_YAML"                    "$release"
resolve_resources "config/core/"                                               "$SERVING_CORE_YAML"                   "$release"
resolve_resources "config/hpa-autoscaling/"                                    "$SERVING_HPA_YAML"                    "$release"
resolve_resources "config/post-install/storage-version-migration.yaml"         "$SERVING_POST_INSTALL_JOBS_YAML"      "$release"
