data "smartos_image" "base_64_2013Q2" {
  name    = "base64"
  version = "13.2.0"
}

output "base_64_2013Q2_uuid" {
  value = "${data.smartos_image.base_64_2013Q2.id}"
}
