# teamspeak

Repository storing a configuration for deploying teamspeak to DigitalOcean. This largely
exists as a means of tinkering with [`packer`](https://www.packer.io/) and [`terraform`](https://www.terraform.io/).
In addition, the end goal is to have a Terraform "provider" for the Teamspeak WebQuery API. This is being worked on
via authoring an [OpenAPI definition](https://swagger.io/docs/specification/) and using [openapi-generator](https://github.com/OpenAPITools/openapi-generator)
to generate a Go library to wrap in a provider.
