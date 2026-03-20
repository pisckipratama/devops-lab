# 1. Membuat Custom VPC
resource "google_compute_network" "vpc" {
  name                    = var.vpc_name
  auto_create_subnetworks = false
}

# 2. Membuat Subnetwork
resource "google_compute_subnetwork" "subnet" {
  name          = var.subnet_name
  ip_cidr_range = var.subnet_cidr
  region        = var.subnet_region
  network       = google_compute_network.vpc.id
}

# 3. Membuat Firewall (Allow SSH & HTTP)
resource "google_compute_firewall" "allow_http_ssh" {
  name    = "${var.vpc_name}-allow-http-ssh"
  network = google_compute_network.vpc.id

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["web-server"] # Ini akan cocok dengan tag di VM Anda
}