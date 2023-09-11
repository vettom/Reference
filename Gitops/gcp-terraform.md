# GCP, Kubernetes with Terraform

# Terraform
- [ ] Local TF install
    - Formt tf code with terraform fmt
    - Validate code with Terraform validate
    - Install [tfswitch](https://github.com/warrensbox/terraform-switcher)
    - Set terraform plugins directory outside of repo
    - Lint tf code with [tflint](https://github.com/terraform-linters/tflint/releases)
    - run tfsec on tf code
- [ ] State file
    - Store state file in S3 with Dynamodb lock
    - Import a resource to tfstate (Eg: instance)

## VPC tasks
- [ ] Using Tf create a vpc and subnet 
    - single subnet
    - secondary ip for K8s
    - Firewall rule allowing incoming conneciton on port 9000
- [ ] Create VPC module 
    - VPC name and subnets are variables
    - Create Cloud NAT and attach
    - Route outbound traffic via CloudNAT

## GKE
- [ ] Build standard GKE cluster
    - Use secondary ip for pod and services
    - Single node, pre-emptible node
    - Enable private node
    - Enable public endpoint
- [ ] Create GKE Cluster module 
    - All inputs should be variables
    - Delete defult nodepool and create new
    - Output: Clustername, endpoint
- [ ] Upgrade GKE cluster to latest version