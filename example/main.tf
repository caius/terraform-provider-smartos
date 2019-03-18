data "smartos_image" "base_64_2013Q2" {
  uuid = "0084dad6-05c1-11e3-9476-8f8320925eea"
}

/*data "smartos_zone" "test1" {
  uuid = "3fced63c-ce05-404f-f635-ecff8144ce55"
}*/


/*resource "smartos_zone" "test1" {
  name       = "test1"
  image_uuid = "${data.smartos_image.base_64_2013Q2.id}"
  type       = "os"
}*/

