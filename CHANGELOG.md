# Changelog

## [0.4.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.3.0...v0.4.0) (January 23, 2021)

### Added

- Resource `api_key`:
  - new optional field `validity` to define the expiration date of key.
  - new optional field `max_queries_per_ip_per_hour` in replacement of `max_queries_per_ip_peer_hour`

### Deprecated

- Resource `api_key`:
  - field `max_queries_per_ip_peer_hour` is deprecated  and removed in 0.5.0

## [0.3.0](https://github.com/philippe-vandermoere/terraform-provider-algolia/compare/v0.2.0...v0.3.0) (August 30, 2020)

### Added

- New Data Source: index

### Changed

- Move the generate Algolia Api Key from attribute `id` into the Sensitive attribute `key`.

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
