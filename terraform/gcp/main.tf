provider "google" {
  project = "pisckilab"
  region  = "asia-southeast2"
}

variable "subnet_ip_cidr_range" {
  description = "ip range untuk semua subnet"
  type        = list(string)
}

variable "subnet_secondary_ip_cidr_range" {
  description = "subnet for secondary dev environment"
}

variable "network_name" {
  description = "nama network kita"
}

variable "subnet_name_list" {
  description = "nama subnet 01"
  type        = list(string)
}

resource "google_compute_network" "development_network" {
  name                    = var.network_name
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "dev_subnet_01" {
  name          = var.subnet_name_list[0]
  ip_cidr_range = var.subnet_ip_cidr_range[0]
  region        = "asia-southeast2"
  network       = google_compute_network.development_network.id
  secondary_ip_range {
    range_name    = "secondary-range-01"
    ip_cidr_range = var.subnet_secondary_ip_cidr_range
  }
}

resource "google_compute_subnetwork" "dev_subnet_02" {
  name          = var.subnet_name_list[1]
  ip_cidr_range = var.subnet_ip_cidr_range[1]
  region        = "asia-southeast1"
  network       = google_compute_network.development_network.id
}

data "google_compute_network" "default_network_existing" {
  name = "default"
}

resource "google_compute_subnetwork" "testing-subnet" {
  name          = "testing-subnet"
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
