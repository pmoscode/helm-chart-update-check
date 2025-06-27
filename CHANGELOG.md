# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

<!-- insertion marker -->
## Unreleased

<small>[Compare with latest](https://github.com/pmoscode/helm-chart-update-check/compare/v0.2.0...HEAD)</small>

### Features

- update semver library to v3.4.0 ([a6f5f09](https://github.com/pmoscode/helm-chart-update-check/commit/a6f5f0949b260c17a7bc1966ceebca4870b37c1c) by Peter Motzko).

<!-- insertion marker -->
## [v0.2.0](https://github.com/pmoscode/helm-chart-update-check/releases/tag/v0.2.0) - 2024-11-17

<small>[Compare with v0.1.3](https://github.com/pmoscode/helm-chart-update-check/compare/v0.1.3...v0.2.0)</small>

### Features

- add "exclude-versions" flag and do small code refactor ([7a66347](https://github.com/pmoscode/helm-chart-update-check/commit/7a66347b308355a7021dcad1a21493fa18e0adc7) by Peter Motzko).
- add dependency version update ([ca97c70](https://github.com/pmoscode/helm-chart-update-check/commit/ca97c70c13f94524408a664a6d0a8dcf5da5f662) by Peter Motzko).

### Bug Fixes

- update module github.com/masterminds/semver/v3 to v3.3.0 ([7cafed9](https://github.com/pmoscode/helm-chart-update-check/commit/7cafed92010a43b099d6ada8817e00eb35f4e596) by renovate[bot]).

## [v0.1.3](https://github.com/pmoscode/helm-chart-update-check/releases/tag/v0.1.3) - 2023-12-11

<small>[Compare with v0.1.2](https://github.com/pmoscode/helm-chart-update-check/compare/v0.1.2...v0.1.3)</small>

### Bug Fixes

- increase patch version to get correct check ([69f9912](https://github.com/pmoscode/helm-chart-update-check/commit/69f9912be734f546d6ca4cfb9438c5ba98ff7723) by Peter Motzko).

## [v0.1.2](https://github.com/pmoscode/helm-chart-update-check/releases/tag/v0.1.2) - 2023-12-11

<small>[Compare with v0.1.1](https://github.com/pmoscode/helm-chart-update-check/compare/v0.1.1...v0.1.2)</small>

### Features

- add CHANGELOG.md ([f4d3902](https://github.com/pmoscode/helm-chart-update-check/commit/f4d390269f88523db566d64be82b0bfb777d5783) by Peter Motzko).
- add build and test workflow ([2efe0b1](https://github.com/pmoscode/helm-chart-update-check/commit/2efe0b1e8f42072e07ff49375f7ce611ff21b3fb) by Peter Motzko).
- add test task and add "go get" to build task ([137f7f5](https://github.com/pmoscode/helm-chart-update-check/commit/137f7f57c6bc2758c911b2ec0791f0aba063f3d6) by Peter Motzko).
- add tests ([f278e97](https://github.com/pmoscode/helm-chart-update-check/commit/f278e972e512ba5a45876a0490f5f56bab4b3a3a) by Peter Motzko).

### Bug Fixes

- versions with meta information are handled correct now ([e34aad6](https://github.com/pmoscode/helm-chart-update-check/commit/e34aad672fa874bcf03445058812022a1d1df579) by Peter Motzko).

## [v0.1.1](https://github.com/pmoscode/helm-chart-update-check/releases/tag/v0.1.1) - 2023-11-28

<small>[Compare with v0.1.0](https://github.com/pmoscode/helm-chart-update-check/compare/v0.1.0...v0.1.1)</small>

### Bug Fixes

- remove strict version parsing to avoid loosing prefixed version numbers ([beed12a](https://github.com/pmoscode/helm-chart-update-check/commit/beed12ab24f3160709c08f9ae8cbee19ff9f40b4) by Peter Motzko).

## [v0.1.0](https://github.com/pmoscode/helm-chart-update-check/releases/tag/v0.1.0) - 2023-11-27

<small>[Compare with first commit](https://github.com/pmoscode/helm-chart-update-check/compare/2f7b9d6761bf31f171b01ef1477518068f30e96f...v0.1.0)</small>

### Features

- update README.md ([de468e7](https://github.com/pmoscode/helm-chart-update-check/commit/de468e7c1e580bb5f04f73029b816abd72524861) by Peter Motzko).
- update git-changelog parameters to 2.4.0 ([138078c](https://github.com/pmoscode/helm-chart-update-check/commit/138078ce5f36197381b74d34f0499b4754201fb5) by Peter Motzko).
- replace environment variables config with cli parameters and add fail-on-update parameter ([48d804d](https://github.com/pmoscode/helm-chart-update-check/commit/48d804dfbe55ce256343874b5643001bfb039857) by Peter Motzko).
- extend error message ([43af8f7](https://github.com/pmoscode/helm-chart-update-check/commit/43af8f7c8f3d102f0ca9488f7424a4692cdd5915) by Peter Motzko).
- outsource DockerHub code to own module and glue all together ([9fe054b](https://github.com/pmoscode/helm-chart-update-check/commit/9fe054b395ed643a58d551724d0775ec77b51bad) by Peter Motzko).
- add release workflow ([7083b49](https://github.com/pmoscode/helm-chart-update-check/commit/7083b49329367ac48eab5e1e9b2fe977b36a56fe) by Peter Motzko).
- add Taskfile (with build and changelog tasks) ([260bf65](https://github.com/pmoscode/helm-chart-update-check/commit/260bf654b276663caf885918a304489d958c0853) by Peter Motzko).
- Initial commit ([e0003dc](https://github.com/pmoscode/helm-chart-update-check/commit/e0003dc4ab03d028a1cf30f472c7227e11095a3c) by Peter Motzko).

