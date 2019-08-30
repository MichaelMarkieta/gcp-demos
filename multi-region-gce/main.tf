provider "google" {
  project = "michaelmarkieta-gcp-demos"
}

provider "google-beta" {
  project = "michaelmarkieta-gcp-demos"
}

#########################################################################################################
# GLOBAL LOAD BALANCER
#########################################################################################################

resource "google_compute_global_address" "default" {
  name    = "multi-region-gce"
  project = "michaelmarkieta-gcp-demos"
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "multi-region-gce-port-80"
  ip_address = "${google_compute_global_address.default.address}"
  port_range = "80"
  target     = "${google_compute_target_http_proxy.default.self_link}"
}

resource "google_compute_target_http_proxy" "default" {
  name    = "multi-region-gce"
  url_map = "${google_compute_url_map.default.self_link}"
}

resource "google_compute_url_map" "default" {
  name            = "multi-region-gce"
  default_service = "${google_compute_backend_service.default.self_link}"
}

resource "google_compute_backend_service" "default" {
  name             = "multi-region-gce-backend"
  protocol         = "HTTP"
  port_name        = "multi-region-gce"
  timeout_sec      = 10
  session_affinity = "NONE"

  backend {
    group = "${google_compute_region_instance_group_manager.instance_group_manager_canada.instance_group}"
  }

  backend {
    group = "${google_compute_region_instance_group_manager.instance_group_manager_singapore.instance_group}"
  }

  health_checks = ["${google_compute_health_check.default.self_link}"]
}

resource "google_compute_health_check" "default" {
  project             = "michaelmarkieta-gcp-demos"
  name                = "autohealing-health-check"
  check_interval_sec  = 5
  timeout_sec         = 5
  healthy_threshold   = 2
  unhealthy_threshold = 10

  http_health_check {
    request_path = "/"
    port         = "80"
  }
}

#########################################################################################################
# COMPUTE ENGINE
#########################################################################################################

resource "google_compute_instance_template" "default" {
  name         = "instance-template"
  machine_type = "n1-standard-1"

  disk {
    source_image = "debian-cloud/debian-9"
  }

  metadata_startup_script = "sudo apt-get update && sudo apt-get install apache2 -y && echo '<!doctype html><html><body><h1>Welcome to my GCP Demo!</h1></body></html>' | sudo tee /var/www/html/index.html"

  network_interface {
    network       = "default"
    access_config = {}
  }

  tags = ["http"]

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_compute_firewall" "default" {
  name    = "default-allow-http"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["http"]
}

#########################################################################################################
# COMPUTE ENGINE - CANADA
#########################################################################################################

resource "google_compute_region_instance_group_manager" "instance_group_manager_canada" {
  provider           = "google-beta"
  name               = "instance-group-manager"
  base_instance_name = "canada"

  version {
    name              = "latest"
    instance_template = "${google_compute_instance_template.default.self_link}"
  }

  region      = "northamerica-northeast1"
  target_size = "3"

  auto_healing_policies {
    health_check      = "${google_compute_health_check.default.self_link}"
    initial_delay_sec = 30
  }
}

#########################################################################################################
# COMPUTE ENGINE - SINGAPORE
#########################################################################################################

resource "google_compute_region_instance_group_manager" "instance_group_manager_singapore" {
  provider           = "google-beta"
  name               = "instance-group-manager"
  base_instance_name = "singapore"

  version {
    name              = "latest"
    instance_template = "${google_compute_instance_template.default.self_link}"
  }

  region      = "asia-southeast1"
  target_size = "3"

  auto_healing_policies {
    health_check      = "${google_compute_health_check.default.self_link}"
    initial_delay_sec = 30
  }
}

#########################################################################################################
# COMPUTE ENGINE - LONDON (europe-west2)
#########################################################################################################

#
#
#
#
#
#
#
