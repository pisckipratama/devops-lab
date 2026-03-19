resource "google_storage_bucket" "portofolio_bucket" {
  name          = "piscki-portofolio-bucket-2429"
  location      = "ASIA-SOUTHEAST2"
  force_destroy = true
}