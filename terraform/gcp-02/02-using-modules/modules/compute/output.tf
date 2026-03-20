output "instance_public_ip" {
  value       = google_compute_instance.app_server.network_interface.0.access_config.0.nat_ip
  description = "IP Public dari VM yang baru dibuat"
}