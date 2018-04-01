# pipeline
Go package implements a container-based pipeline engine

Pipeline is a kubernetes native, in cluster continuous integration and delivery platform that runs defined pipelines in reaction to event triggers (i.e. updated code in a GitHub repository).

- `piped` is a runtime executor that builds docker images inside a cluster; considered a daemon
- `pipectl` is a CLI tool that triggers the pipeline API in the same fashion as event triggers
- `pipeline-gateway` normalizes external events to comply to `pipeline`s event schema
- `piped-stash` is a pluggable backend for the `pipeline` for perisitent storage
- `pipeline`s are a series of related jobs that are run in parallel/consecutively inside a cluster
- `pipeline-controller` is a kubernetes custom resource definition that configures the runtime daemon

### Gettin things workin

The `Pipeline` platform requires two `ConfigMaps`; one for `piped` configurations, and another that defines scripts to run against a `Deployment`'s container image, via using `spec.labels.selectors`. 

### More details

- For a given `Deployment`, a git-sidecar is run to monitor the current docker image, and compares it to the most up-to-date docker image that exists in `x` registry.

### Flow

1. Trigger
2. Event occurs
3. `pipeline-controller` parses the event and hands off instructions to a worker via the `piped` runtime
4. the worker runs a `Pipeline`, followed by building the new image and updating a selected `Deployment`
