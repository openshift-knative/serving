# Openshift Knative Serving

This repository holds Openshift's fork of
[`knative/serving`](https://github.com/knative/serving) with additions and
fixes needed only for the OpenShift side of things.

## How this repository works ?

The default branch holds up-to-date specific [openshift files](./openshift) 
that are necessary for CI setups and maintaining it. This includes:

- Scripts to create a new release branch from `upstream`
- CI setup files
  - operator configuration (for Openshift's CI setup)
  - tests scripts
- Operator's base configurations

Each release branch holds the upstream code for that release and our
openshift's specific files.

## CI Setup

For the CI setup, three repositories are of importance:

- This repository
- [openshift/release](https://github.com/openshift/release) which contains the configuration of CI jobs that are run on this repository
- [openshift-knative/hack](https://github.com/openshift-knative/hack) which is used to generate CI job definitions
  
All of the following is based on OpenShift’s CI operator
configs. General understanding of that mechanism is assumed in the
following documentation.

The job manifests for the CI jobs are generated via [openshift-knative/hack](https://github.com/openshift-knative/hack). The
basic configuration lives in the
[ci-operator/config/openshift-knative/serving](https://github.com/openshift/release/tree/master/ci-operator/config/openshift-knative/serving) folder of the
[openshift/release](https://github.com/openshift/release) repository. These files include which version to
build against (OCP 4.x), which images to build
(this includes all the images needed to run Knative and also all the
images required for running e2e tests) and which command to execute
for the CI jobs to run (more on this later).

Before we can create the ci-operator configs mentioned above, we need
to make sure there are Dockerfiles for all images that we need
(they’ll be referenced by the ci-operator config hence we need to
create them first). The [generate-dockerfiles.sh](https://github.com/openshift-knative/serving/blob/master/openshift/ci-operator/generate-dockerfiles.sh) script takes care of
creating all the Dockerfiles needed automatically. The files now need
to be committed to the branch that CI is being setup for.

The Knative Serving release files are generated via the
generate-release.sh file in the openshift-knative/serving
repository. 

After the file has been added to the folder as mentioned above, the
job manifests itself will need to be generated as is described in the
corresponding [ci-operator documentation](https://docs.google.com/document/d/1SQ_qlkcplqhe8h6ONXdgBr7YUVbs4oRSj4ISl3gpLW4/edit#heading=h.8w7nj9363nsd). 
This process is automated using [openshift-knative/hack](https://github.com/openshift-knative/hack).

Once all of this is done (Dockerfiles committed, ci-operator config
created and job manifests generated and pushed), the CI setup for that branch is active.

## Create a new release

Refer to the [release manual](https://docs.google.com/document/d/18HVtCbvOuUpunORixcVdGNK6ZPjA9NybcGgVSP0g1LM/edit#heading=h.mam6uyjn1dzv).
