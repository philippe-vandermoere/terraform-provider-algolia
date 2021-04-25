# Terraform Provider Algolia

[![licence](https://img.shields.io/github/license/philippe-vandermoere/terraform-provider-algolia)](./LICENSE)
![test](https://github.com/philippe-vandermoere/terraform-provider-algolia/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/philippe-vandermoere/terraform-provider-algolia)](https://goreportcard.com/report/github.com/philippe-vandermoere/terraform-provider-algolia)
[![codecov](https://codecov.io/gh/philippe-vandermoere/terraform-provider-algolia/branch/main/graph/badge.svg?token=3O9O1LKS38)](https://codecov.io/gh/philippe-vandermoere/terraform-provider-algolia)

This repository is a Algolia provider for [Terraform](https://www.terraform.io).

See [usage](https://registry.terraform.io/providers/philippe-vandermoere/algolia/latest/docs).

## Developer

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.16

### Installation

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command:

````bash
git clone https://github.com/philippe-vandermoere/terraform-provider-algolia
cd terraform-provider-algolia
make install
````
