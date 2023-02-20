module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "3.14.0"
  name = "FargateVPC"
  cidr = var.vpc_cidr

  azs             = ["eu-west-1a", "eu-west-1b" ]
  private_subnets = [var.priv_subnet1,]
  public_subnets  = [var.pub_subnet1, var.pub_subnet2]

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = {
    Terraform = "true"
    Environment = "dev"
    Owner  = "Vettom"
  }
}

