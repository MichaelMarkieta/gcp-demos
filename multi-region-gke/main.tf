provider "google" {
  project = "michaelmarkieta-gcp-demos"
}

provider "google-beta" {
  project = "michaelmarkieta-gcp-demos"
}

/*

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

  backend {
    group = "${google_compute_region_instance_group_manager.instance_group_manager_london.instance_group}"
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

*/

#########################################################################################################
# KUBERNETES ENGINE
#########################################################################################################

resource "google_container_cluster" "cluster_us_central1" {
  name     = "cluster-us-central1"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count = 1
  master_auth {
    username = ""
    password = ""
    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

resource "google_container_node_pool" "cluster_us_central1_preemptible_nodes" {
  name       = "nodepool-us-central1-preemptible"
  location   = "us-central1"
  cluster    = "cluster-us-central1"
  node_count = 1
  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"
    metadata = {
      disable-legacy-endpoints = "true"
    }
    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}
