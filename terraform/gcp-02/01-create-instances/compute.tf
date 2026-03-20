# Extra block storage
resource "google_compute_disk" "extra_disk" {
  name = "data-disk"
  type = "pd-standard"
  zone = "asia-southeast2-a"
  size = 10
}

# VM Instance
resource "google_compute_instance" "app_server" {
  name         = "docker-web-server"
  machine_type = "e2-micro"
  zone         = "asia-southeast2-a"
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
    network    = google_compute_network.custom_vpc.id
    subnetwork = google_compute_subnetwork.custom_subnet.id
    access_config {
      # give IP public
    }
  }

  metadata = {
    ssh-keys = "ansible-user:${file("~/.ssh/id_rsa.pub")}"
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