#!/usr/bin/env bash

set -euo pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

tmp_dir=$(mktemp -d)
git clone --branch main https://github.com/openshift-knative/hack "$tmp_dir"

pushd "$tmp_dir"
go install github.com/openshift-knative/hack/cmd/generate
popd

rm -rf "$tmp_dir"

$(go env GOPATH)/bin/generate \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --app-file-fmt "/ko-app/%s" \
  --dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-9-release-golang-%s-nofips-openshift-4.19"
