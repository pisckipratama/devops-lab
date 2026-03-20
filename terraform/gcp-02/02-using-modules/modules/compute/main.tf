# Extra block storage
resource "google_compute_disk" "extra_disk" {
  name = "${var.instance_name}-data-disk"
  type = "pd-standard"
  zone = var.zone
  size = 10
}

# VM Instance
resource "google_compute_instance" "app_server" {
  name         = var.instance_name
  machine_type = var.machine_type
  zone         = var.zone
  tags         = ["web-server"]

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2404-lts-amd64"
      size  = 20
    }
  }

  attached_disk {
    source      = google_compute_disk.extra_disk.id
    device_name = "extra-data"
  }

  network_interface {
    network    = var.network_id
    subnetwork = var.subnetwork_id
    access_config {
      # give IP public
    }
  }

  metadata = {
    ssh-keys = "ansible-user:${file("${var.ssh_pub_key_path}")}"
  }

  metadata_startup_script = <<-EOF
    #!/bin/bash

    while fuser /var/lib/dpkg/lock-frontend > /dev/null 2>&1 || fuser /var/lib/dpkg/lock >/dev/null 2>&1; do
      sleep 3;
    done

    # Matikan interaksi manual
    export DEBIAN_FRONTEND=noninteractive
    export NEEDRESTART_MODE=a

    # 2. Instalasi Docker Engine
    apt-get update
    apt-get install -y docker.io

    # 3. Pastikan Docker berjalan dan otomatis nyala saat VM reboot
    systemctl start docker
    systemctl enable docker

    # 4. Deploy Aplikasi menggunakan Docker
    # -d : Menjalankan container di background (detached)
    # -p 80:80 : Memetakan port 80 di VM ke port 80 di dalam Container
    # --restart always : Container akan otomatis menyala lagi jika VM di-restart
    docker run -d -p 80:80 --name whoami-app --restart always traefik/whoami
  EOF
}