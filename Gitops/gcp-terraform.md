# GCP, Kubernetes with Terraform

# Terraform
- [ ] Local TF install
    - Formt tf code with terraform fmt
    - Validate code with Terraform validate
    - Install [tfswitch](https://github.com/warrensbox/terraform-switcher)
    - Set terraform plugins directory outside of repo
    - Lint tf code with [tflint](https://github.com/terraform-linters/tflint/releases)
    - run tfsec on tf code
    - Store state file in S3 with Dynamodb lock

## VPC tasks
- [ ] Using Tf create a vpc and subnet
    - single subnet
    - secondary ip for K8s
    - Firewall rule allowing incoming conneciton on port 9000
- [ ] Create VPC module
    - VPC name and subnets are variables
