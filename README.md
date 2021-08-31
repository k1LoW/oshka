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
2021-08-31T20:40:24+09:00 [INFO] Create temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/
2021-08-31T20:40:24+09:00 [INFO] Extract local . to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-31T20:40:25+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-31T20:40:25.362+0900    INFO    Using your github token
2021-08-31T20:40:25.540+0900    INFO    Number of language-specific files: 5
2021-08-31T20:40:25.540+0900    INFO    Detecting gobinary vulnerabilities...
2021-08-31T20:40:25.542+0900    INFO    Detecting gomod vulnerabilities...

dist/goreleaserdocker077582362/oshka (gobinary)
===============================================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
|             LIBRARY              | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION  |                 TITLE                 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
| github.com/containerd/containerd | CVE-2021-32760   | MEDIUM   | v1.5.3            | v1.4.8, v1.5.4 | containerd: pulling and               |
|                                  |                  |          |                   |                | extracting crafted container          |
|                                  |                  |          |                   |                | image may result in Unix file...      |
|                                  |                  |          |                   |                | -->avd.aquasec.com/nvd/cve-2021-32760 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+

dist/oshka-darwin-windows_darwin_amd64/oshka (gobinary)
=======================================================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
|             LIBRARY              | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION  |                 TITLE                 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
| github.com/containerd/containerd | CVE-2021-32760   | MEDIUM   | v1.5.3            | v1.4.8, v1.5.4 | containerd: pulling and               |
|                                  |                  |          |                   |                | extracting crafted container          |
|                                  |                  |          |                   |                | image may result in Unix file...      |
|                                  |                  |          |                   |                | -->avd.aquasec.com/nvd/cve-2021-32760 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+

dist/oshka-darwin-windows_windows_amd64/oshka.exe (gobinary)
============================================================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
|             LIBRARY              | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION  |                 TITLE                 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
| github.com/containerd/containerd | CVE-2021-32760   | MEDIUM   | v1.5.3            | v1.4.8, v1.5.4 | containerd: pulling and               |
|                                  |                  |          |                   |                | extracting crafted container          |
|                                  |                  |          |                   |                | image may result in Unix file...      |
|                                  |                  |          |                   |                | -->avd.aquasec.com/nvd/cve-2021-32760 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+

dist/oshka-linux_linux_amd64/oshka (gobinary)
=============================================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
|             LIBRARY              | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION  |                 TITLE                 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+
| github.com/containerd/containerd | CVE-2021-32760   | MEDIUM   | v1.5.3            | v1.4.8, v1.5.4 | containerd: pulling and               |
|                                  |                  |          |                   |                | extracting crafted container          |
|                                  |                  |          |                   |                | image may result in Unix file...      |
|                                  |                  |          |                   |                | -->avd.aquasec.com/nvd/cve-2021-32760 |
+----------------------------------+------------------+----------+-------------------+----------------+---------------------------------------+

go.sum (gomod)
==============
Total: 18 (UNKNOWN: 3, LOW: 0, MEDIUM: 7, HIGH: 8, CRITICAL: 0)

+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
|              LIBRARY               | VULNERABILITY ID | SEVERITY |         INSTALLED VERSION         |             FIXED VERSION             |                  TITLE                  |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/apache/thrift           | CVE-2019-0205    | HIGH     | 0.12.0                            | 0.13.0                                | thrift: Endless loop when               |
|                                    |                  |          |                                   |                                       | feed with specific input data           |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-0205    |
+                                    +------------------+          +                                   +                                       +-----------------------------------------+
|                                    | CVE-2019-0210    |          |                                   |                                       | thrift: Out-of-bounds read              |
|                                    |                  |          |                                   |                                       | related to TJSONProtocol                |
|                                    |                  |          |                                   |                                       | or TSimpleJSONProtocol                  |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-0210    |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2020-13949   |          |                                   | v0.14.0                               | libthrift: potential DoS when           |
|                                    |                  |          |                                   |                                       | processing untrusted payloads           |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-13949   |
+------------------------------------+------------------+          +-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/buger/jsonparser        | CVE-2020-10675   |          | 0.0.0-20180808090653-f4dd9f5a6b44 | v0.0.0-20200321185410-91ac96899e49    | golang-github-buger-jsonparser:         |
|                                    |                  |          |                                   |                                       | infinite loop via a Delete call         |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-10675   |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2020-35381   |          |                                   | v1.1.1                                | jsonparser: GET call can lead to        |
|                                    |                  |          |                                   |                                       | a slice bounds out of range...          |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-35381   |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/containerd/containerd   | CVE-2021-32760   | MEDIUM   | 1.5.3                             | v1.4.8, v1.5.4                        | containerd: pulling and                 |
|                                    |                  |          |                                   |                                       | extracting crafted container            |
|                                    |                  |          |                                   |                                       | image may result in Unix file...        |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2021-32760   |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/dgrijalva/jwt-go        | CVE-2020-26160   | HIGH     | 3.2.0+incompatible                |                                       | jwt-go: access restriction              |
|                                    |                  |          |                                   |                                       | bypass vulnerability                    |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-26160   |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/gorilla/handlers        | GO-2020-0020     | UNKNOWN  | 0.0.0-20150720190736-60c7bfde3e33 | v1.3.0                                |                                         |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/miekg/dns               | CVE-2019-19794   | MEDIUM   | 1.0.14                            | v1.1.25-0.20191211073109-8ebf2e419df7 | golang-github-miekg-dns: predictable    |
|                                    |                  |          |                                   |                                       | TXID can lead to response forgeries     |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-19794   |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/sassoftware/go-rpmutils | CVE-2020-7667    | HIGH     | 0.0.0-20190420191620-a8f1baeba37b | v0.1.0                                | In package                              |
|                                    |                  |          |                                   |                                       | github.com/sassoftware/go-rpmutils/cpio |
|                                    |                  |          |                                   |                                       | before version 0.1.0, the               |
|                                    |                  |          |                                   |                                       | CPIO extraction functionality           |
|                                    |                  |          |                                   |                                       | doesn't sanitize...                     |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-7667    |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/satori/go.uuid          | GO-2020-0018     | UNKNOWN  | 1.2.0                             | v1.2.1-0.20181016170032-d91630c85102  |                                         |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| github.com/ulikunitz/xz            | CVE-2021-29482   | HIGH     | 0.5.7                             | v0.5.8                                | ulikunitz/xz: Infinite                  |
|                                    |                  |          |                                   |                                       | loop in readUvarint allows              |
|                                    |                  |          |                                   |                                       | for denial of service                   |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2021-29482   |
+                                    +------------------+----------+                                   +                                       +-----------------------------------------+
|                                    | GO-2020-0016     | UNKNOWN  |                                   |                                       |                                         |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
| k8s.io/kubernetes                  | CVE-2019-1002101 | MEDIUM   | 1.13.0                            | 1.11.9, 1.12.7, 1.13.5,               | kubernetes: Mishandling of              |
|                                    |                  |          |                                   | 1.14.1-beta.0                         | symlinks allows for arbitrary           |
|                                    |                  |          |                                   |                                       | file write via `kubectl cp`...          |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-1002101 |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2019-11250   |          |                                   | v1.16.0-beta.1                        | kubernetes: Bearer tokens               |
|                                    |                  |          |                                   |                                       | written to logs at high                 |
|                                    |                  |          |                                   |                                       | verbosity levels (>= 7)...              |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-11250   |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2020-8554    |          |                                   |                                       | kubernetes: MITM using                  |
|                                    |                  |          |                                   |                                       | LoadBalancer or ExternalIPs             |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-8554    |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2020-8564    |          |                                   | v1.20.0-alpha.1                       | kubernetes: Docker config               |
|                                    |                  |          |                                   |                                       | secrets leaked when file is             |
|                                    |                  |          |                                   |                                       | malformed and loglevel >=...            |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-8564    |
+                                    +------------------+          +                                   +---------------------------------------+-----------------------------------------+
|                                    | CVE-2020-8565    |          |                                   | v1.20.0-alpha.2                       | kubernetes: Incomplete fix              |
|                                    |                  |          |                                   |                                       | for CVE-2019-11250 allows for           |
|                                    |                  |          |                                   |                                       | token leak in logs when...              |
|                                    |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-8565    |
+------------------------------------+------------------+----------+-----------------------------------+---------------------------------------+-----------------------------------------+
2021-08-31T20:40:25+09:00 [INFO] Detect action actions/setup-go@v2 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-31T20:40:25+09:00 [INFO] Detect action actions/checkout@v2 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-31T20:40:25+09:00 [INFO] Detect action golangci/golangci-lint-action@v2 from /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/local-cdb4ee2
2021-08-31T20:40:25+09:00 [INFO] Extract action actions/setup-go@v2 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/setup-go@v2
Enumerating objects: 1035, done.
Counting objects: 100% (31/31), done.
Compressing objects: 100% (29/29), done.
Total 1035 (delta 12), reused 8 (delta 0), pack-reused 1004
2021-08-31T20:40:26+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/setup-go@v2
2021-08-31T20:40:26.931+0900    INFO    Using your github token
2021-08-31T20:40:26.988+0900    INFO    Number of language-specific files: 1
2021-08-31T20:40:26.988+0900    INFO    Detecting npm vulnerabilities...

package-lock.json (npm)
=======================
Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)

2021-08-31T20:40:26+09:00 [INFO] Extract action actions/checkout@v2 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v2
Enumerating objects: 997, done.
Counting objects: 100% (28/28), done.
Compressing objects: 100% (21/21), done.
Total 997 (delta 11), reused 11 (delta 6), pack-reused 969
2021-08-31T20:40:29+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/actions/checkout@v2
2021-08-31T20:40:29.201+0900    INFO    Using your github token
2021-08-31T20:40:29.261+0900    INFO    Number of language-specific files: 1
2021-08-31T20:40:29.261+0900    INFO    Detecting npm vulnerabilities...

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
2021-08-31T20:40:29+09:00 [INFO] Extract action golangci/golangci-lint-action@v2 to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/golangci/golangci-lint-action@v2
Enumerating objects: 1342, done.
Counting objects: 100% (431/431), done.
Compressing objects: 100% (241/241), done.
Total 1342 (delta 329), reused 268 (delta 188), pack-reused 911
2021-08-31T20:40:31+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/golangci/golangci-lint-action@v2
2021-08-31T20:40:31.226+0900    INFO    Using your github token
2021-08-31T20:40:31.281+0900    INFO    Number of language-specific files: 2
2021-08-31T20:40:31.281+0900    INFO    Detecting npm vulnerabilities...
2021-08-31T20:40:31.282+0900    INFO    Detecting gomod vulnerabilities...

package-lock.json (npm)
=======================
Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)


sample-go-mod/go.sum (gomod)
============================
Total: 4 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 3, CRITICAL: 0)

+-----------------------------+------------------+----------+-----------------------------------+---------------------------------------+---------------------------------------+
|           LIBRARY           | VULNERABILITY ID | SEVERITY |         INSTALLED VERSION         |             FIXED VERSION             |                 TITLE                 |
+-----------------------------+------------------+----------+-----------------------------------+---------------------------------------+---------------------------------------+
| github.com/dgrijalva/jwt-go | CVE-2020-26160   | HIGH     | 3.2.0+incompatible                |                                       | jwt-go: access restriction            |
|                             |                  |          |                                   |                                       | bypass vulnerability                  |
|                             |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-26160 |
+-----------------------------+------------------+          +-----------------------------------+---------------------------------------+---------------------------------------+
| github.com/gogo/protobuf    | CVE-2021-3121    |          | 1.2.1                             | v1.3.2                                | gogo/protobuf:                        |
|                             |                  |          |                                   |                                       | plugin/unmarshal/unmarshal.go         |
|                             |                  |          |                                   |                                       | lacks certain index validation        |
|                             |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2021-3121  |
+-----------------------------+------------------+----------+-----------------------------------+---------------------------------------+---------------------------------------+
| github.com/miekg/dns        | CVE-2019-19794   | MEDIUM   | 1.0.14                            | v1.1.25-0.20191211073109-8ebf2e419df7 | golang-github-miekg-dns: predictable  |
|                             |                  |          |                                   |                                       | TXID can lead to response forgeries   |
|                             |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2019-19794 |
+-----------------------------+------------------+----------+-----------------------------------+---------------------------------------+---------------------------------------+
| golang.org/x/crypto         | CVE-2020-29652   | HIGH     | 0.0.0-20200622213623-75b288015ac9 | v0.0.0-20201216223049-8b5274cf687f    | golang: crypto/ssh: crafted           |
|                             |                  |          |                                   |                                       | authentication request can            |
|                             |                  |          |                                   |                                       | lead to nil pointer dereference       |
|                             |                  |          |                                   |                                       | -->avd.aquasec.com/nvd/cve-2020-29652 |
+-----------------------------+------------------+----------+-----------------------------------+---------------------------------------+---------------------------------------+
2021-08-31T20:40:31+09:00 [INFO] Cleanup temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/

Run results
===========
+----------------------------------+--------+--------------------------+-----------+------------------------------------------------------------------+
|               NAME               |  TYPE  |         COMMAND          | EXIT CODE |                               HASH                               |
+----------------------------------+--------+--------------------------+-----------+------------------------------------------------------------------+
| .                                | local  | trivy fs --exit-code 1 . | 1         | 2d978015910831113e1895405220d3c83dfd2e49316d282ac07b7746a00e4234 |
|                                  |        |                          |           | (dir hash)                                                       |
| actions/setup-go@v2              | action | trivy fs --exit-code 1 . | 0         | 331ce1d993939866bb63c32c6cbbfd48fa76fc57 (commit hash)           |
| actions/checkout@v2              | action | trivy fs --exit-code 1 . | 1         | 5a4ac9002d0be2fb38bd78e4b4dbde5606d7042f (commit hash)           |
| golangci/golangci-lint-action@v2 | action | trivy fs --exit-code 1 . | 1         | 5c56cd6c9dc07901af25baab6f2b0d9f3b7c3018 (commit hash)           |
+----------------------------------+--------+--------------------------+-----------+------------------------------------------------------------------+
$
```

</details>

### Scan remote Git repository and supply chains for vulnerabilities using Trivy

``` console
$ oshka run repo github.com/rails/rails
```

<details>

<summary>Result</summary>

``` console
$ oshka run repo github.com/cli/cli
2021-08-31T00:46:39+09:00 [INFO] Create temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/
2021-08-31T00:46:39+09:00 [INFO] Extract repo github.com/cli/cli to /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/github.com/cli/cli
Enumerating objects: 26086, done.
Counting objects: 100% (743/743), done.
Compressing objects: 100% (579/579), done.
Total 26086 (delta 256), reused 412 (delta 160), pack-reused 25343
2021-08-31T00:46:46+09:00 [INFO] Run `trivy fs --exit-code 1 .` on /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/github.com/cli/cli
2021-08-31T00:46:47.049+0900    INFO    Using your github token
2021-08-31T00:46:47.130+0900    INFO    Number of language-specific files: 1
2021-08-31T00:46:47.130+0900    INFO    Detecting gomod vulnerabilities...

[...]

2021-08-31T00:48:05+09:00 [INFO] Cleanup temporary directory for extracting supply chains: /var/folders/fp/hk95_wsj7s18mmc9drvrxdp1tt294n/T/

Run results
===========
+----------------------------------------+--------+--------------------------+-----------+------------------------------------------+
|                  NAME                  |  TYPE  |         COMMAND          | EXIT CODE |                   HASH                   |
+----------------------------------------+--------+--------------------------+-----------+------------------------------------------+
| github.com/cli/cli                     | repo   | trivy fs --exit-code 1 . | 1         | e6ff77ce73c201b0ee36d2b802ea45e9e1ad1822 |
|                                        |        |                          |           | (commit hash)                            |
| github/codeql-action/analyze@v1        | action | trivy fs --exit-code 1 . | 1         | 33f3438c1d59883f5e769fdf2b6adb6794d91d0f |
|                                        |        |                          |           | (commit hash)                            |
| actions/setup-go@v2                    | action | trivy fs --exit-code 1 . | 0         | 331ce1d993939866bb63c32c6cbbfd48fa76fc57 |
|                                        |        |                          |           | (commit hash)                            |
| goreleaser/goreleaser-action@v2        | action | trivy fs --exit-code 1 . | 0         | 5a54d7e660bda43b405e8463261b3d25631ffe86 |
|                                        |        |                          |           | (commit hash)                            |
| mislav/bump-homebrew-formula-action@v1 | action | trivy fs --exit-code 1 . | 0         | d631ddd46015c5c3c4e3f0da275c15d99475d760 |
|                                        |        |                          |           | (commit hash)                            |
| actions/checkout@v2                    | action | trivy fs --exit-code 1 . | 1         | 5a4ac9002d0be2fb38bd78e4b4dbde5606d7042f |
|                                        |        |                          |           | (commit hash)                            |
| github/codeql-action/init@v1           | action | trivy fs --exit-code 1 . | 1         | 33f3438c1d59883f5e769fdf2b6adb6794d91d0f |
|                                        |        |                          |           | (commit hash)                            |
+----------------------------------------+--------+--------------------------+-----------+------------------------------------------+
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
+-----------------------------------+--------+--------------------------+-----------+-------------------------------------------------------------------------+
|               NAME                |  TYPE  |         COMMAND          | EXIT CODE |                                  HASH                                   |
+-----------------------------------+--------+--------------------------+-----------+-------------------------------------------------------------------------+
| actions/cache@v2                  | action | trivy fs --exit-code 1 . | 0         | c64c572235d810460d0d6876e9c705ad5002b353                                |
|                                   |        |                          |           | (commit hash)                                                           |
| github/codeql-action/analyze@v1   | action | trivy fs --exit-code 1 . | 1         | 33f3438c1d59883f5e769fdf2b6adb6794d91d0f                                |
|                                   |        |                          |           | (commit hash)                                                           |
| ubuntu:latest                     | image  | trivy fs --exit-code 1 . | 1         | sha256:10cbddb6cf8568f56584ccb6c866203e68ab8e621bb87038e254f6f27f955bbe |
|                                   |        |                          |           | (digest)                                                                |
| datadog/squid:latest              | image  | trivy fs --exit-code 1 . | 1         | sha256:f7d19d5e3f4163771291d91de393ce667f2327a3e080c39b9b7ea9e19f91488f |
|                                   |        |                          |           | (digest)                                                                |
| actions/setup-node@v1             | action | trivy fs --exit-code 1 . | 1         | f1f314fca9dfce2769ece7d933488f076716723e (commit hash)                  |
| actions/checkout@v2               | action | trivy fs --exit-code 1 . | 1         | 5a4ac9002d0be2fb38bd78e4b4dbde5606d7042f (commit hash)                  |
| github/codeql-action/init@v1      | action | trivy fs --exit-code 1 . | 1         | 33f3438c1d59883f5e769fdf2b6adb6794d91d0f (commit hash)                  |
| github/codeql-action/autobuild@v1 | action | trivy fs --exit-code 1 . | 1         | 33f3438c1d59883f5e769fdf2b6adb6794d91d0f (commit hash)                  |
+-----------------------------------+--------+--------------------------+-----------+-------------------------------------------------------------------------+
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
        - When using Dockerfile, require `docker` for building image.
- Docker image

## References

- [aquasecurity/trivy](https://github.com/aquasecurity/trivy): Scanner for vulnerabilities in container images, file systems, and Git repositories, as well as for configuration issues
- [Security hardening for GitHub Actions](https://docs.github.com/en/actions/learn-github-actions/security-hardening-for-github-actions#using-third-party-actions): "Using third-party actions"
