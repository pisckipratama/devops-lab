output "instance_public_ip" {
  value       = google_compute_instance.app_server.network_interface.0.access_config.0.nat_ip
  description = "IP Public dari VM yang baru dibuat"
}

output "instance_self_link" {
  value = google_compute_instance.app_server.self_link
  description = "uniq identity link from VM for insert to LB"
}