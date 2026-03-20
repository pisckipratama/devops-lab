output "network_id" {
  value       = google_compute_network.vpc.id
  description = "ID dari VPC yang terbuat"
}

output "subnetwork_id" {
  value       = google_compute_subnetwork.subnet.id
  description = "ID dari Subnetwork yang terbuat"
}