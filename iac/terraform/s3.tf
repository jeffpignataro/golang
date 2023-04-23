resource "aws_s3_bucket" "example" {
  bucket = "${local.account_name}-terraform-test-bucket"

  tags = {
    Name        = "Dummy test bucket"
    Environment = "Dev"
    Automation  = "Terraform"
  }
}
