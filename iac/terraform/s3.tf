resource "aws_s3_bucket" "example" {
  bucket = "my-tf-test-bucket"

  tags = {
    Name        = "Dummy test bucket"
    Environment = "Dev"
    Automation  = "Terraform"
  }
}
