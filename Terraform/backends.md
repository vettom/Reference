# Backends
 By default local (file) backend is used and state file reated in current directory.
 When a new back end is configured you must run terraform init and will be prompted for migration
- A configuration can only have single back end
- Backend configuration does not accept named value like variables. No interpolation allowed.
```terraform
    backend "remote" {
    organization = "example_corp"

    workspaces {
      name = "my-app-prod"
    }
```
 
## Terraform Cloud
 If using TF Cloud, you cannot use bakend configuration
```terraform
    cloud {
    organization = "vettom"

    workspaces {
      name = "dvtest"
    }
  }
```

## Partial configuration
  You dont have to configure all required fileds for backend. Items like credentials can be provided as arguments.
  - File  : -backend-config=PATH option when running terraform init
  - Command line : To specify a single key/value pair, use the -backend-config="KEY=VALUE" option when running terraform init
  - Interactive :  TF prompts for values while running init.
