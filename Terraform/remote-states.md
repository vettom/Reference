# Using remote state file.
Importing state file to make use of variables. 

 - Source state files output variables only can be imported.
 - Data module terraform_remote_stat used to aceess state file local or remote
 - Data is available as data.terraform_remote_state.{import name}.outputs.{VAR}
 - Path must include state filename for local.

> Sharing state not recommended as it require access to full state file. Ideally post to external source.   

## Import local stat file
```terraform
      data terraform_remote_state "mastervpc" {
          backend = "local"
          config = {
            path = "../VPC/terraform.tfstate"
          }
        }
```
## Import remote state file.
 
```terraform
    data "terraform_remote_state" "vpc" {
      backend = "remote"
    
      config = {
        organization = "hashicorp"
        workspaces = {
          name = "vpc-prod"
        }
      }
    }
```