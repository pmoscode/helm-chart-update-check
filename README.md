# helm-chart-update-check

Checks Docker hub for Docker images tag update. It compares the given Helm chart "appVersion" (which must be used as default version for the container image tag of the deployment) with the image tags on DockerHub.

## Usage

```text
hcuc <flags>

Flags:
  --debug
        Enable debug outputs
  --docker-hub-repo string
        DockHub repo to check tag versions
  --fail-on-update
        Return exit code 1, if update is available
  --helm-chart-path string
        Helm chart to check for updates (default ".")
```

Flags `--docker-hub-repo string` and `--helm-chart-path` are required.

Per default, it shows a small table with the current version and all available versions.
When `--fail-on-update` is set, the app is exiting with code 1.
