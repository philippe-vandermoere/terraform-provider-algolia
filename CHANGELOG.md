# Changelog

## [0.7.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.6.0...v0.7.0) (April 25, 2021)

### Added

- Resource `index`:
  - new optional field `attribute_for_distinct` to define the name of the de-duplication attribute to be used with the distinct feature.
  - new optional field `distinct` to enables de-duplication or grouping of results

- Datasource `index`:
  - new field `distinct`

- CI:
  - Implement build test
  - Implement Terraform acceptance test
  - Implement Codecov report
  - Implement Github dependabot for go modules dependencies

### Changed

- Update algolia SDK to v3.18.1
- Update terraform SDK to v2.6.1
- Implement terraform best practice for provider:
  - Rename provider golang namespace `algolia` to `internal/provider`
  - Use apiClient struct for pass algoliaClient

## [0.6.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.5.0...v0.6.0) (March 21, 2021)

### Added

- New Resource: `index`

### Fixed

- Datasource `index`:
  - rename `attributesfor_faceting` to `attributes_for_faceting`

## [0.5.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.4.0...v0.5.0) (February 28, 2021)

### Changed

- Update go version to 1.16
- Update algolia SDK to v3.16.0
- Update terraform SDK to v2.4.4
- Update golangci-lint to v1.37.1

### Removed

- Resource `api_key`:
  - remove deprecated field `max_queries_per_ip_peer_hour`

## [0.4.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.3.0...v0.4.0) (January 23, 2021)

### Added

- Resource `api_key`:
  - new optional field `validity` to define the expiration date of key
  - new optional field `max_queries_per_ip_per_hour` in replacement of `max_queries_per_ip_peer_hour`

### Deprecated

- Resource `api_key`:
  - field `max_queries_per_ip_peer_hour` is deprecated and removed in 0.5.0

## [0.3.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.2.0...v0.3.0) (August 30, 2020)

### Added

- New Data Source: index

### Changed

- Move the generate Algolia Api Key from attribute `id` into the Sensitive attribute `key`

## [0.2.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.1.0...v0.2.0) (August 27, 2020)

### Added

- Terraform Registry [Documentation](https://www.terraform.io/docs/registry/providers/docs.html)

### Changed

- Using the Changelog format based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)

### Fixed

- Fix import GPG key in CD

## [0.1.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/releases/tag/v0.1.0) (August 26, 2020)

### Added

- New Provider: algolia
- New Resource: api_key
