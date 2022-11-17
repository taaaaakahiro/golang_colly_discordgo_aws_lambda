module "lambda" {
  source = "./modules/lambda"
  hook_real_estate = var.hook_esc_key
  target_url = var.target_url

}

#module "build" {
#  source = "./modules/build"
#}
