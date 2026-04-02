variable "lb_name" {
  type        = string
  description = "Nama dasar untuk komponen Load Balancer"
}

variable "network_id" {
  type = string
}

variable "zone" {
  type    = string
  default = "asia-southeast2-a"
}

variable "backend_instances" {
  type        = list(string)
  description = "Daftar self_link dari VM yang akan dimasukkan ke backend"
}