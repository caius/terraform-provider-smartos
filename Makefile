.PHONY: dev
dev:
	@make example-plan

.PHONY: build
build: test
	go build

.PHONY: test
test:
	cd smartos/ && go test

.PHONY: example-init
example-init: build
	cd example && terraform init

.PHONY: example-plan
example-plan: example-init
	cd example && terraform plan --input=false
