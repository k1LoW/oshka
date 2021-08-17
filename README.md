# oshka

`oshka` is a tool for extracting nested CI/CD supply chains and executing commands.

## Concept

Security checks should be performed not only on the source code of the repository, but also on the code of the third-party actions of GitHub Actions and Docker images that compose the CI/CD.

The primary purpose of `oshka` is for the continuous security check of the nested CI/CD supply chains ( So the default execution `--command` is `trivy fs --exit-code 1 .`. ).

Because most tools can be run on the filesystem, oshka has a strategy of deploying the supply chains in temporary directories.

### Behavior

1. Analyze the nested supply chains that compose the CI/CD
  - Detect repositories of third-party actions.
  - Detect docker images.
2. Extract the resources of supply chains to local temporary directory.
2. Execute any commands to the extracted resources.

## Usage

### Scan local filesystem and supply chains for vulnerabilities using Trivy

``` console
$ oshka fs .
```

### Scan action of GitHub Actions and supply chains for vulnerabilities using Trivy

``` console
$ oshka action actions/cache@v2
```

## Supported supply chains

- GitHub Actions
  - Workflow file (ex. `.github/workflows/*.yml` )
  - Action file (ex. `action.yml` )
- Docker image

## References

- [aquasecurity/trivy](https://github.com/aquasecurity/trivy): Scanner for vulnerabilities in container images, file systems, and Git repositories, as well as for configuration issues
- [Security hardening for GitHub Actions "Using third-party actions"](https://docs.github.com/en/actions/learn-github-actions/security-hardening-for-github-actions#using-third-party-actions)
