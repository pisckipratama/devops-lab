# instance group
resource "google_compute_instance_group" "web_group" {
  name = "${var.lb_name}-ig"
  description = "unmanaged instance group for web servers"
  zone = var.zone
  network = var.network_id
  instances = var.backend_instances

  named_port {
    name = "http"
    port = 80
  }
}

# Health check
resource "google_compute_health_check" "http_health_check" {
  name = "${var.lb_name}-hc"
  check_interval_sec = 5
  timeout_sec = 5

  tcp_health_check {
    port = 80
  }
}