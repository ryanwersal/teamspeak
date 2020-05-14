output "droplet_ip" {
  description = "IP of the Teamspeak droplet"
  value       = digitalocean_droplet.ts.ipv4_address
}
