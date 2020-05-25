output "droplet_ip" {
  description = "IP of the Teamspeak droplet"
  value       = digitalocean_floating_ip.ip.ip_address
}
