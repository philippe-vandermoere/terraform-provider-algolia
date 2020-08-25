ifndef VERBOSE
	MAKEFLAGS += --no-print-directory
endif

default: install

.PHONY: install lint unit

OS_ARCH=linux_amd64
HOSTNAME=registry.terraform.io
NAMESPACE=philippe-vandermoere
NAME=algolia
VERSION=99.99.99
TERRAFORM_PLUGINS_DIRECTORY=~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

install: lint unit
	mkdir -p ${TERRAFORM_PLUGINS_DIRECTORY}
	go build -o ${TERRAFORM_PLUGINS_DIRECTORY}/terraform-provider-${NAME}
	cd examples && terraform init

lint:
	 golangci-lint run

unit:
	go test ./algolia
