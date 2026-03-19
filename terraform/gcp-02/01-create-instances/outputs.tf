output "web_server_ip" {
  value       = google_compute_instance.app_server.network_interface.0.access_config.0.nat_ip
  description = "Public IP of our Golang web server"
}