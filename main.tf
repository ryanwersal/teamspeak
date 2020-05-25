provider "digitalocean" {}

locals {
  region        = "nyc3"
  all_addresses = ["0.0.0.0/0", "::/0"]
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

resource "digitalocean_floating_ip" "ip" {
  region     = local.region
  droplet_id = digitalocean_droplet.ts.id
}

data "http" "ip" {
  url = "https://api.ipify.org/"
}

locals {
  self_ip_cidr = ["${chomp(data.http.ip.body)}/32"]
}

resource "digitalocean_firewall" "fw" {
  name = "teamspeak-firewall"

  droplet_ids = [digitalocean_droplet.ts.id]

  # SSH
  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = local.self_ip_cidr
  }

  # Teamspeak voice
  inbound_rule {
    protocol         = "udp"
    port_range       = "9987"
    source_addresses = local.all_addresses
  }

  # Teamspeak files
  inbound_rule {
    protocol         = "tcp"
    port_range       = "30033"
    source_addresses = local.all_addresses
  }

  # Teamspeak WebQuery
  inbound_rule {
    protocol         = "tcp"
    port_range       = "10080"
    source_addresses = local.self_ip_cidr
  }

  # HTTPS (teamspeak server list queries)
  outbound_rule {
    protocol              = "tcp"
    port_range            = "443"
    destination_addresses = local.all_addresses
  }

  # DNS
  outbound_rule {
    protocol              = "udp"
    port_range            = "53"
    destination_addresses = local.all_addresses
  }
}

resource "digitalocean_domain" "domain" {
  name = "accidental-bravery.com"
}

resource "digitalocean_record" "ts" {
  domain = digitalocean_domain.domain.name
  type   = "A"
  name   = "ts"
  value  = digitalocean_floating_ip.ip.ip_address
}

resource "digitalocean_project" "ts" {
  name        = "teamspeak"
  description = "Resources for the Teamspeak server"
  purpose     = "Voice Chat"
  environment = "Production"
  resources = [
    digitalocean_droplet.ts.urn,
    digitalocean_domain.domain.urn,
    digitalocean_floating_ip.ip.urn,
  ]
}

provider "teamspeak" {
  endpoint  = "http://${digitalocean_record.ts.fqdn}:10080/1"
  api_token = var.teamspeak_token
}

resource "teamspeak_channel" "channels" {
  depends_on = [digitalocean_droplet.ts, digitalocean_domain.domain]
  for_each   = toset(["League of Legends", "Apex Legends", "Raft", "Vermintide"])

  channel_name        = each.value
  channel_topic       = "${each.value} Topic"
  channel_description = "${each.value} Description"
  permanent           = true
}
