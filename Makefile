.PHONY: dev
dev:
	@make build
	@make example-plan

.PHONY: build
build:
	go build

.PHONY: example-init
example-init:
	cd example && terraform init

.PHONY: example-plan
example-plan: example-init
	cd example && terraform plan
