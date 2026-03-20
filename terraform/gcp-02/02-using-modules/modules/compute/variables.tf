variable "instance_name" {
  type = string
  description = "VM name"
}

variable "zone" {
  type = string
  default = "asia-southeast2-a"
}

variable "machine_type" {
  type = string
  default = "e2-micro"
}

variable "network_id" {
  type = string
  description = "VPC network's id"
}

variable "subnetwork_id" {
  type = string
  description = "subnet's id"
}

variable "ssh_pub_key_path" {
  type = string
  description = "Local path to public key SSH"
}