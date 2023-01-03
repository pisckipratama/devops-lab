provider "google" {
  project = "pisckilab"
  region  = "asia-southeast2"
}

variable "subnet_secondary_ip_cidr_range" {
  description = "subnet for secondary dev environment"
}

variable "network_name" {
  description = "nama network kita"
}

variable "subnet_list" {
  description = "variable for all of subnets value"
  type = list(object({
    range  = string
    name   = string
    region = string
  }))
}

resource "google_compute_network" "development_network" {
  name                    = var.network_name
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "dev_subnet_01" {
  name          = var.subnet_list[0].name
  ip_cidr_range = var.subnet_list[0].range
  region        = var.subnet_list[0].region
  network       = google_compute_network.development_network.id
  secondary_ip_range {
    range_name    = "secondary-range-01"
    ip_cidr_range = var.subnet_secondary_ip_cidr_range
  }
}

resource "google_compute_subnetwork" "dev_subnet_02" {
  name          = var.subnet_list[1].name
  ip_cidr_range = var.subnet_list[1].range
  region        = var.subnet_list[1].region
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
