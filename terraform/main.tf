module "lambda" {
  source = "./modules/lambda"
  hook_real_estate = var.hook_esc_key
  target_url = var.target_url

}

#module "s3" {
#  source = "./modules/s3"
#}
#
#module "build" {
#  source = "./modules/build"
#}
