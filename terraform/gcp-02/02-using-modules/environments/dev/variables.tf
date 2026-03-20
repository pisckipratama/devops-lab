variable "project_id" {
  type = string
}

variable "region" {
    type = string
    default = "asia-southeast2"
}

variable "zone" {
  type = string
  default = "asia-southeast2-a"
}

variable "bucket" {
  type = string
}

variable "environment" {
  type = string
}

variable "ssh_key_file" {
  type = string
}