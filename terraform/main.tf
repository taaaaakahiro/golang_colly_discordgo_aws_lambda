module "lambda" {
  source = "./modules/lambda"
  hook_real_estate = var.hook_real_estate
  target_url = var.target_url

}
