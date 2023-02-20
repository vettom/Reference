# Query Data source
 Each workspce can query output of another to polulate values. For eg, app build job can query information from
VPC work space to decide ELB and other details
```terraform
    data "terraform_remote_state" "vpc" {
  backend = "local"

  config = {
    path = "../learn-terraform-data-sources-vpc/terraform.tfstate"
  }
}

```

- Create VPC and network in a work space
- create app in another workspace and use data from VPC state.