# teamspeak

## Overview

Repository storing a configuration for deploying teamspeak to DigitalOcean. This largely exists as a means of tinkering with [`packer`](https://www.packer.io/) and [`terraform`](https://www.terraform.io/). In addition, the end goal is to have a Terraform "provider" for the Teamspeak WebQuery API. This is being worked on via authoring an [OpenAPI definition](https://swagger.io/docs/specification/) and using [openapi-generator](https://github.com/OpenAPITools/openapi-generator) to generate a Go library to wrap in a provider. The provider consumes the library to mutate resources via TS' management API.

## Packer

The basic Teamspeak DigitalOcean droplet image is built using `packer` using the definition in `packer/ts.json` (the various sibling files are
configuration files and scripts for provisioning the image).

## Terraform

The `*.tf` files define the DigitalOcean resources for the Teamspeak server (its hosting droplet, routes, and so on). It consumes the `packer` generated image. Once the Terraform provider is minimally complete and usable it will also handle provisioning and managing the Teamspeak server.

## OpenAPI

If you want to tinker (or read) the OpenAPI definition it can be found at `openapi/ts.yml`. Once modifications are made, the root `Makefile` can be used to generate the go library and execute the tests:

```bash
make openapi-generate-go openapi-test
```

To make it easier to test out WebQuery (using [`httpie`](https://github.com/jakubroztocil/httpie) or curl) there is a `Dockerfile` which pulls in the latest `teamspeak` docker image and exposes the new WebQuery port. To run this, use the `Makefile`:

```bash
make openapi-container
```

The API token will be in the output of the command and must be passed with all requests in a header named `x-api-key`. [See the TS Server 3.12.x announcement post](https://community.teamspeak.com/t/teamspeak-server-3-12-x/3916) for further details.

## Terraform Provider

A minimal prototype of a Terraform provider is implemented. It can manage a very select (read: minimal) set of properties of a channel but appears to work correctly for create/deletes. To compile the provider you can run:

```bash
make tf-provider-build
```

Further development will heavily depend on the continued development of the OpenAPI specification.
