provider "google" {
  project = "pisckilab"
  region  = "asia-southeast2"
}

variable "subnet_ip_cidr_range" {
  description = "subnet for dev environment"
}

variable "subnet_secondary_ip_cidr_range" {
  description = "subnet for secondary dev environment"
}

resource "google_compute_network" "development_network" {
  name                    = "development-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "dev_subnet_01" {
  name          = "dev-subnet-01"
  ip_cidr_range = var.subnet_ip_cidr_range
  region        = "asia-southeast2"
  network       = google_compute_network.development_network.id
  secondary_ip_range {
    range_name    = "secondary-range-01"
    ip_cidr_range = var.subnet_secondary_ip_cidr_range
  }
}

data "google_compute_network" "default_network_existing" {
  name = "default"
}

resource "google_compute_subnetwork" "dev_subnet_02" {
  name          = "dev-subnet-02"
  ip_cidr_range = "10.110.0.0/16"
  network       = data.google_compute_network.default_network_existing.id
  region        = "asia-southeast2"
}

output "development_network_id" {
  value = google_compute_network.development_network.id
}

output "dev_subnet_01_gateway" {
  value = google_compute_subnetwork.dev_subnet_01.gateway_address
}
