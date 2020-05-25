init:
	terraform init

build:
	(cd packer && packer build ts.json)

apply:
	terraform apply

destroy:
	terraform destroy

deploy: build apply

ssh:
	ssh root@ts.accidental-bravery.com

openapi-container:
	(cd openapi && docker build --tag ts .)
	docker run -itp 10080:10080 -e TS3SERVER_LICENSE=accept -e TS3SERVER_QUERY_PROTOCOLS=http ts:latest

openapi-generate-go:
	(cd openapi && openapi-generator generate -i ts.yml -g go -o golib)

openapi-test:
	(cd openapi/golib && go test)