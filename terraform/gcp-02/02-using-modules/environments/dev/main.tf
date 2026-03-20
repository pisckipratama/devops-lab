module "dev_network" {
  source = "../../modules/network"

  vpc_name = "${var.environment}-vpc"
  subnet_name = "${var.environment}-subnet-jkt"
  subnet_cidr = "10.0.1.0/24"
  subnet_region = "asia-southeast2"
}

module "web_server_frontend" {
  source = "../../modules/compute"

  instance_name = "${var.environment}-frontend-server"
  machine_type = "e2-micro"
  zone = "asia-southeast2-a"

  network_id = module.dev_network.network_id
  subnetwork_id = module.dev_network.subnetwork_id

  ssh_pub_key_path = var.ssh_key_file
}