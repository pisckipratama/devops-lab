terraform {
  backend "gcs" {
    bucket = "tf-state-piscki-devops-029"
    prefix = "terraform/state"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = "nifty-matrix-483111-m9"
  region  = "asia-southeast2"
  zone    = "asia-southeast2-a"
}