# Changelog

## [v0.7.0](https://github.com/moznion/gonstructor/compare/v0.6.0...v0.7.0) - 2026-10-01

- Add `-getterPrefix` command line flag to customize the Getter prefix. by @ilovelinux in https://github.com/moznion/gonstructor/pull/53
- Update module github.com/stretchr/testify to v1.12.1 by @renovate[bot] in https://github.com/moznion/gonstructor/pull/51
- Update go toolchain directive to v1.27.1 by @renovate[bot] in https://github.com/moznion/gonstructor/pull/43
- Introduce Songmu/tagpr for release management by @moznion in https://github.com/moznion/gonstructor/pull/54
- Apply version pinning by actions-lock by @moznion in https://github.com/moznion/gonstructor/pull/56

## [v0.6.0](https://github.com/moznion/gonstructor/compare/v0.5.1...v0.6.0) - 2025-09-19

- Support go 1.24+, drop go 1.23 or earlier by @moznion in https://github.com/moznion/gonstructor/pull/33
- chore: Configure Renovate by @renovate[bot] in https://github.com/moznion/gonstructor/pull/34
- Update module github.com/iancoleman/strcase to v0.3.0 by @renovate[bot] in https://github.com/moznion/gonstructor/pull/35
- Simplify workflows release.yml by @moznion in https://github.com/moznion/gonstructor/pull/38
- Update module github.com/moznion/gowrtr to v1.7.0 by @renovate[bot] in https://github.com/moznion/gonstructor/pull/36
- feat: Option to return value instead of pointer by @jj-style in https://github.com/moznion/gonstructor/pull/30
- Refactor: rename internal method: s/withPrefix/withConditionalPrefix/g by @moznion in https://github.com/moznion/gonstructor/pull/39
- Update module github.com/stretchr/testify to v1.11.1 by @renovate[bot] in https://github.com/moznion/gonstructor/pull/40
- Add support for specifying a prefix for generated setter methods by @moznion in https://github.com/moznion/gonstructor/pull/41
- :pencil: doc by @moznion in https://github.com/moznion/gonstructor/pull/42

## [v0.5.1](https://github.com/moznion/gonstructor/compare/v0.5.0...v0.5.1) - 2024-06-28

- Remove the typo `"` in the README file by @LintaoAmons in https://github.com/moznion/gonstructor/pull/24
- fix Generation error on Go 1.22 by @doilux in https://github.com/moznion/gonstructor/pull/27
- fix: panic when generating in package with generics by @jj-style in https://github.com/moznion/gonstructor/pull/26
- Drop the unsupported go runtime versions by @moznion in https://github.com/moznion/gonstructor/pull/28
- Upgrade GitHub Actions actions by @moznion in https://github.com/moznion/gonstructor/pull/29

## [v0.5.0](https://github.com/moznion/gonstructor/compare/v0.4.1...v0.5.0) - 2023-01-25

- Fix builder to use toLowerCamel for assignment parameter by @seitau in https://github.com/moznion/gonstructor/pull/20
- fix: multiple types import by @seitau in https://github.com/moznion/gonstructor/pull/21
- Tweak for #21 by @moznion in https://github.com/moznion/gonstructor/pull/22

## [v0.4.1](https://github.com/moznion/gonstructor/compare/v0.4.0...v0.4.1) - 2022-08-23

- Truncate and open a file when it open that file at the first time by @moznion in https://github.com/moznion/gonstructor/pull/18
- Use golanglint-ci instead of golint by @moznion in https://github.com/moznion/gonstructor/pull/19

## [v0.4.0](https://github.com/moznion/gonstructor/compare/v0.3.1...v0.4.0) - 2022-08-22

- add note about goimports by @kovetskiy in https://github.com/moznion/gonstructor/pull/7
- update go to 1.19, update deps, fix #9 by @kovetskiy in https://github.com/moznion/gonstructor/pull/11
- Upgrade the GitHub Actions configurations to support go 1.19.x by @moznion in https://github.com/moznion/gonstructor/pull/12
- Accept the multiple `--type` options in order to output the generated code of the multiple types into a single file by @moznion in https://github.com/moznion/gonstructor/pull/14
- Support `-propagateInitFuncReturns` CLI option by @moznion in https://github.com/moznion/gonstructor/pull/16

## [v0.3.1](https://github.com/moznion/gonstructor/compare/v0.2.0...v0.3.1) - 2021-02-04

- Add support for embedding by @kovetskiy in https://github.com/moznion/gonstructor/pull/6

## [v0.2.0](https://github.com/moznion/gonstructor/compare/v0.1.0...v0.2.0) - 2020-06-01

- avoid generating long lines by @kovetskiy in https://github.com/moznion/gonstructor/pull/4

## [v0.1.0](https://github.com/moznion/gonstructor/compare/v0.0.1...v0.1.0) - 2020-02-25

- Add installation guide by @Mushus in https://github.com/moznion/gonstructor/pull/2
- Use own toLowerCamel for arguments since ID must be id not iD. by @mattn in https://github.com/moznion/gonstructor/pull/1
- add -init parameter by @kovetskiy in https://github.com/moznion/gonstructor/pull/3

## [v0.0.1](https://github.com/moznion/gonstructor/commits/v0.0.1) - 2019-12-08
