# oshka

`oshka`【oʊʃkə】is a tool for extracting nested CI/CD supply chains and executing commands.

## Concept

Security checks should be performed not only on the source code of the repository, but also on the code of the third-party actions of GitHub Actions and Docker images that compose the CI/CD.

The primary purpose of `oshka` is for the continuous security check of the nested CI/CD supply chains ( So the default execution `--command` is `trivy fs --exit-code 1 .` ).

Because most tools can be run on the filesystem, oshka has a strategy of deploying the supply chains in temporary directories.

### Behavior

1. Analyze the nested supply chains that compose the CI/CD
    - Detect repositories of third-party actions.
    - Detect docker images.
2. Extract the resources of supply chains to local temporary directory.
2. Execute any commands to the extracted resources.

## Usage

### Scan local filesystem and supply chains for vulnerabilities using Trivy

( The default execution `--command` is `trivy fs --exit-code 1 .` )

``` console
$ oshka run fs .
```

<details>

<summary>Result</summary>

``` console
$ oshka run fs .
2021-08-18T08:12:18+09:00 [INFO] Create temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/
2021-08-18T08:12:18+09:00 [INFO] Extract local . to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:18+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:20.995+0900    INFO    Using your github token
2021-08-18T08:12:20.997+0900    INFO    Need to update DB
2021-08-18T08:12:20.997+0900    INFO    Downloading DB...
2.59 MiB / 23.00 MiB [-------------------->___________________________________________________________________________________________________________________________________________________________________] 11.27% ? p/s ?8.23 MiB / 23.00 MiB [----------------------------------------------------------------->______________________________________________________________________________________________________________________] 35.80% ? p/s ?15.66 MiB / 23.00 MiB [---------------------------------------------------------------------------------------------------------------------------->__________________________________________________________] 68.11% ? p/s ?21.46 MiB / 23.00 MiB [-------------------------------------------------------------------------------------------------------------------------------------------------------------->___________] 93.31% 31.40 MiB p/s ETA 0s23.00 MiB / 23.00 MiB [-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 32.04 MiB p/s 1s2021-08-18T08:12:22.942+0900    INFO    Number of language-specific files: 1
2021-08-18T08:12:22.942+0900    INFO    Detecting gomod vulnerabilities...

go.sum (gomod)
==============
Total: 3 (UNKNOWN: 2, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+------------------+------------------+----------+-------------------+---------------+---------------------------------------+
|     LIBRARY      | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION |                 TITLE                 |
+------------------+------------------+----------+-------------------+---------------+---------------------------------------+
| gopkg.in/yaml.v2 | CVE-2019-11254   | MEDIUM   | 2.2.2             | v2.2.8        | kubernetes: Denial of                 |
|                  |                  |          |                   |               | service in API server via             |
|                  |                  |          |                   |               | crafted YAML payloads by...           |
|                  |                  |          |                   |               | -->avd.aquasec.com/nvd/cve-2019-11254 |
+                  +------------------+----------+                   +---------------+---------------------------------------+
|                  | GMS-2019-2       | UNKNOWN  |                   | v2.2.3        | XML Entity Expansion                  |
+                  +------------------+          +                   +               +---------------------------------------+
|                  | GO-2021-0061     |          |                   |               |                                       |
+------------------+------------------+----------+-------------------+---------------+---------------------------------------+
2021-08-18T08:12:22+09:00 [INFO] Detect action actions/setup-go@v1 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:22+09:00 [INFO] Detect action actions/checkout@v1 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:22+09:00 [INFO] Detect action codecov/codecov-action@v1 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:22+09:00 [INFO] Detect action actions/checkout@v2 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-18T08:12:22+09:00 [INFO] Extract action actions/setup-go@v1 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/setup-go@v1
Enumerating objects: 1017, done.
Counting objects: 100% (15/15), done.
Compressing objects: 100% (14/14), done.
Total 1017 (delta 8), reused 2 (delta 1), pack-reused 1002
2021-08-18T08:12:24+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/setup-go@v1
2021-08-18T08:12:24.761+0900    INFO    Using your github token
2021-08-18T08:12:24.832+0900    INFO    Number of language-specific files: 1
2021-08-18T08:12:24.832+0900    INFO    Detecting npm vulnerabilities...

package-lock.json (npm)
=======================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 1, CRITICAL: 0)

+------------+------------------+----------+-------------------+---------------+---------------------------------------+
|  LIBRARY   | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION |                 TITLE                 |
+------------+------------------+----------+-------------------+---------------+---------------------------------------+
| underscore | CVE-2021-23358   | HIGH     | 1.8.3             | 1.12.1        | nodejs-underscore: Arbitrary code     |
|            |                  |          |                   |               | execution via the template function   |
|            |                  |          |                   |               | -->avd.aquasec.com/nvd/cve-2021-23358 |
+------------+------------------+----------+-------------------+---------------+---------------------------------------+
2021-08-18T08:12:24+09:00 [INFO] Extract action actions/checkout@v1 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v1
Enumerating objects: 997, done.
Counting objects: 100% (27/27), done.
Compressing objects: 100% (20/20), done.
Total 997 (delta 10), reused 11 (delta 6), pack-reused 970
2021-08-18T08:12:25+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v1
2021-08-18T08:12:25.945+0900    INFO    Using your github token
2021-08-18T08:12:26.000+0900    INFO    Number of language-specific files: 0
2021-08-18T08:12:26+09:00 [INFO] Extract action codecov/codecov-action@v1 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/codecov/codecov-action@v1
Enumerating objects: 3873, done.
Counting objects: 100% (820/820), done.
Compressing objects: 100% (324/324), done.
Total 3873 (delta 601), reused 653 (delta 493), pack-reused 3053
2021-08-18T08:12:28+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/codecov/codecov-action@v1
2021-08-18T08:12:28.460+0900    INFO    Using your github token
2021-08-18T08:12:28.535+0900    INFO    Number of language-specific files: 1
2021-08-18T08:12:28.535+0900    INFO    Detecting npm vulnerabilities...

package-lock.json (npm)
=======================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 1, CRITICAL: 0)

+------------+------------------+----------+-------------------+---------------+---------------------------------------+
|  LIBRARY   | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION |                 TITLE                 |
+------------+------------------+----------+-------------------+---------------+---------------------------------------+
| path-parse | CVE-2021-23343   | HIGH     | 1.0.6             | 1.0.7         | nodejs-path-parse:                    |
|            |                  |          |                   |               | ReDoS via splitDeviceRe,              |
|            |                  |          |                   |               | splitTailRe and splitPathRe           |
|            |                  |          |                   |               | -->avd.aquasec.com/nvd/cve-2021-23343 |
+------------+------------------+----------+-------------------+---------------+---------------------------------------+
2021-08-18T08:12:28+09:00 [INFO] Extract action actions/checkout@v2 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v2
Enumerating objects: 997, done.
Counting objects: 100% (28/28), done.
Compressing objects: 100% (21/21), done.
Total 997 (delta 11), reused 11 (delta 6), pack-reused 969
2021-08-18T08:12:29+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v2
2021-08-18T08:12:29.810+0900    INFO    Using your github token
2021-08-18T08:12:29.889+0900    INFO    Number of language-specific files: 1
2021-08-18T08:12:29.889+0900    INFO    Detecting npm vulnerabilities...

package-lock.json (npm)
=======================
Total: 3 (UNKNOWN: 0, LOW: 0, MEDIUM: 2, HIGH: 1, CRITICAL: 0)

+---------------+------------------+----------+-------------------+---------------------+---------------------------------------+
|    LIBRARY    | VULNERABILITY ID | SEVERITY | INSTALLED VERSION |    FIXED VERSION    |                 TITLE                 |
+---------------+------------------+----------+-------------------+---------------------+---------------------------------------+
| @actions/core | CVE-2020-15228   | MEDIUM   | 1.1.3             | 1.2.6               | Environment Variable                  |
|               |                  |          |                   |                     | Injection in GitHub Actions           |
|               |                  |          |                   |                     | -->avd.aquasec.com/nvd/cve-2020-15228 |
+---------------+------------------+          +-------------------+---------------------+---------------------------------------+
| node-fetch    | CVE-2020-15168   |          | 2.6.0             | 3.0.0-beta.9, 2.6.1 | node-fetch: size of data after        |
|               |                  |          |                   |                     | fetch() JS thread leads to DoS        |
|               |                  |          |                   |                     | -->avd.aquasec.com/nvd/cve-2020-15168 |
+---------------+------------------+----------+-------------------+---------------------+---------------------------------------+
| underscore    | CVE-2021-23358   | HIGH     | 1.8.3             | 1.12.1              | nodejs-underscore: Arbitrary code     |
|               |                  |          |                   |                     | execution via the template function   |
|               |                  |          |                   |                     | -->avd.aquasec.com/nvd/cve-2021-23358 |
+---------------+------------------+----------+-------------------+---------------------+---------------------------------------+
2021-08-18T08:12:29+09:00 [INFO] Cleanup temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/

Run results
===========
+---------------------------+--------+--------------------------+-----------+
|           NAME            |  TYPE  |         COMMAND          | EXIT CODE |
+---------------------------+--------+--------------------------+-----------+
| .                         | local  | trivy fs --exit-code 1 . | 1         |
| actions/setup-go@v1       | action | trivy fs --exit-code 1 . | 1         |
| actions/checkout@v1       | action | trivy fs --exit-code 1 . | 0         |
| codecov/codecov-action@v1 | action | trivy fs --exit-code 1 . | 1         |
| actions/checkout@v2       | action | trivy fs --exit-code 1 . | 1         |
+---------------------------+--------+--------------------------+-----------+
$
```

</details>

### Scan action of GitHub Actions and supply chains for vulnerabilities using Trivy

``` console
$ oshka run action actions/cache@v2
```

<details>

<summary>Result</summary>

``` console
$ oshka run action actions/cache@v2
2021-08-18T02:17:18+09:00 [INFO] Create temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/
2021-08-18T02:17:18+09:00 [INFO] Extract action actions/cache@v2 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/cache@v2

[...]

Run results
===========
+-----------------------------------+--------+--------------------------+-----------+
|               NAME                |  TYPE  |         COMMAND          | EXIT CODE |
+-----------------------------------+--------+--------------------------+-----------+
| actions/cache@v2                  | action | trivy fs --exit-code 1 . | 0         |
| ubuntu:latest                     | image  | trivy fs --exit-code 1 . | 1         |
| datadog/squid:latest              | image  | trivy fs --exit-code 1 . | 1         |
| actions/checkout@v2               | action | trivy fs --exit-code 1 . | 1         |
| github/codeql-action/init@v1      | action | trivy fs --exit-code 1 . | 1         |
| github/codeql-action/autobuild@v1 | action | trivy fs --exit-code 1 . | 1         |
| github/codeql-action/analyze@v1   | action | trivy fs --exit-code 1 . | 1         |
| actions/setup-node@v1             | action | trivy fs --exit-code 1 . | 1         |
+-----------------------------------+--------+--------------------------+-----------+
$
```

</details>

### Scan more deep supply chains

``` console
$ oshka run fs . --depth 3
```

## Supported supply chains

- GitHub Actions
    - Workflow file (ex. `.github/workflows/*.yml` )
    - Action file (ex. `action.yml` )
- Docker image

## References

- [aquasecurity/trivy](https://github.com/aquasecurity/trivy): Scanner for vulnerabilities in container images, file systems, and Git repositories, as well as for configuration issues
- [Security hardening for GitHub Actions](https://docs.github.com/en/actions/learn-github-actions/security-hardening-for-github-actions#using-third-party-actions): "Using third-party actions"
