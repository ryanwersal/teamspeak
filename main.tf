provider "digitalocean" {}

locals {
  region = "nyc3"
}

data "digitalocean_ssh_key" "default" {
  name = "Default"
}

data "digitalocean_droplet_snapshot" "snapshot" {
  name_regex  = "ts-packer-[0-9]+"
  region      = local.region
  most_recent = true
}

resource "digitalocean_droplet" "ts" {
  name     = "teamspeak"
  image    = data.digitalocean_droplet_snapshot.snapshot.id
  size     = "s-1vcpu-1gb"
  region   = local.region
  ssh_keys = [data.digitalocean_ssh_key.default.fingerprint]
}

resource "digitalocean_domain" "domain" {
  name = "accidental-bravery.com"
}

resource "digitalocean_record" "ts" {
  domain = digitalocean_domain.domain.name
  type   = "A"
  name   = "ts"
  value  = digitalocean_droplet.ts.ipv4_address
}

resource "digitalocean_project" "ts" {
  name        = "teamspeak"
  description = "Resources for the Teamspeak server"
  purpose     = "Voice Chat"
  environment = "Production"
  resources = [
    digitalocean_droplet.ts.urn,
    digitalocean_domain.domain.urn
  ]
}
