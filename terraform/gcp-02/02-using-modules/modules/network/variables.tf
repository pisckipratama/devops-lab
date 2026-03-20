variable "vpc_name" {
  type        = string
  description = "Nama VPC Network"
}

variable "subnet_name" {
  type        = string
  description = "Nama Subnetwork"
}

variable "subnet_cidr" {
  type        = string
  description = "Blok IP untuk Subnetwork (misal: 10.0.1.0/24)"
}

variable "subnet_region" {
  type        = string
  default     = "asia-southeast2"
}