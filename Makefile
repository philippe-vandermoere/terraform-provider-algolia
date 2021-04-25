ifndef VERBOSE
	MAKEFLAGS += --no-print-directory
endif

default: install

.PHONY: install lint unit apply destroy

OS_ARCH=linux_amd64
HOSTNAME=registry.terraform.io
NAMESPACE=philippe-vandermoere
NAME=algolia
VERSION=99.99.99
TERRAFORM_PLUGINS_DIRECTORY=~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

install:
	mkdir -p ${TERRAFORM_PLUGINS_DIRECTORY}
	go build -o ${TERRAFORM_PLUGINS_DIRECTORY}/terraform-provider-${NAME}
	cd examples && rm -rf .terraform && rm -f .terraform.lock.hcl
	cd examples && terraform init

lint:
	golangci-lint run

unit:
	TF_ACC=1 go test ./internal/provider -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

apply:
	cd examples && terraform apply

destroy:
	cd examples && terraform destroy
