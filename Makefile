init:
	terraform init

build:
	(cd packer && packer build ts.json)

apply:
	terraform apply

log-apply:
	TF_LOG_PATH=logs.txt TF_LOG=trace terraform apply

destroy:
	terraform destroy

deploy: build apply

ssh:
	ssh root@ts.accidental-bravery.com

openapi-container:
	(cd openapi && docker build --tag ts .)
	docker run -itp 10080:10080 -e TS3SERVER_LICENSE=accept -e TS3SERVER_QUERY_PROTOCOLS=http ts:latest

openapi-generate-go:
	(cd openapi && openapi-generator generate --input-spec ts.yml --generator-name go --output golib \
		--git-repo-id teamspeak --git-user-id ryanwersal --minimal-update \
		--package-name teamspeak --model-package teamspeak)

openapi-test:
	(cd openapi/golib && go test)

tf-provider-build:
	(cd terraform && go build -o ../terraform-provider-teamspeak)