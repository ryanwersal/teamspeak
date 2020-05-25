module github.com/ryanwersal/terraform-provider-teamspeak

go 1.14

require (
	github.com/antihax/optional v1.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.13.0
	github.com/ryanwersal/teamspeak v0.0.0-00010101000000-000000000000
)

replace github.com/ryanwersal/teamspeak => ../openapi/golib
