# Custom VPC
resource "google_compute_network" "custom_vpc" {
  name                    = "piscki-portofolio-vpc"
  auto_create_subnetworks = false
}

# Create Subnet
resource "google_compute_subnetwork" "custom_subnet" {
  name          = "piscki-portofolio-subnet"
  ip_cidr_range = "10.0.29.0/24"
  region        = "asia-southeast2"
  network       = google_compute_network.custom_vpc.id
}

# Firewall rules open port 80 & 22
resource "google_compute_firewall" "allow_web_ssh" {
  name    = "allow-web-ssh"
  network = google_compute_network.custom_vpc.name

  allow {
    protocol = "tcp"
    ports    = ["22", "80"]
  }

  source_ranges = ["0.0.0.0/0"]
}