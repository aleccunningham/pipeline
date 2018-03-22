# pipectl

`pipectl` is a command-line tool of the Pipeline Operator for creating, listing, checking status of, getting logs of,
and deleting `Pipeline`s. Each function is implemented as a sub-command of `sparkctl`.

To build `pipectl`, run the following command from within `pipectl/`:

```bash
$ go build -o pipectl
```

## Flags

The following global flags are available for all the sub commands:
* `--namespace`: the Kubernetes namespace of the `Pipeline`(s). Defaults to `default`.
* `--kubeconfig`: the path to the file storing configuration for accessing the Kubernetes API server. Defaults to
`$HOME/.kube/config`

## Available Commands

### List

`list` is a sub command of `pipectl` for listing `Pipeline` objects in the namespace specified by
`--namespace`.

Usage:
```bash
$ pipectl list --namespace monitoring
```

### Status

`status` is a sub command of `pipectl` for checking and printing the status of a `Pipeline` in the namespace
specified by `--namespace`.

Usage:
```bash
$ pipectl status <Pipeline name> --namespace monitoring
```

### Exec

`exec` is a sub command of `pipectl` for creating an agent to run the steps defined in a `Pipeline` with the given name in the
namespace specified by `--namespace`. The command by default fetches the logs of the driver pod. To make it fetch logs
of an executor pod instead, use the flag `--executor` to specify the ID of the executor whose logs should be fetched.

Usage:
```bash
$ pipectl log <Pipeline name> [--executor <executor ID, e.g., 1>]
```

### Delete

`delete` is a sub command of `pipectl` for delete `Pipeline` with the given name in the namespace
specified by `--namespace`.

Usage:
```bash
$ pipectl delete <SparkApplication name>
```
