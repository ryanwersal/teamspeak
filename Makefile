init:
	terraform init

build:
	(cd packer && packer build ts.json)

apply:
	terraform apply

destroy:
	terraform destroy

upgrade: build apply

ssh:
	ssh root@ts.accidental-bravery.com