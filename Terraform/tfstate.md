# terraform.tfstate
 Stores state of infrastructure managed by TF. 
 > If state file is modified, must ensure configuration updated to reflect change.
### Migration to cloud
  Create necessary configuration and run Terraform init. 
  TF Backend can be migrated without making local copy.
  No option to migrate out of TF Cloud by default.
  Modifying state file does not have any impact until apply/plan executed
  
### Tfstate commands
```terraform
    tf show  : show saved plan
    tf state list  List resources in the state
          - pull  Pull current state and show on Stdout
          - push  Push state to remote
          - show  To show attributes of specific resource. eg aws_vpc.vpc01
          - replace-provider  Replace provider for a resource in state.
          - mv to move resource to another state file or rename.
```
### Recreate a resource
 -replace command allos individual resource to be reprovisioned. taint is old. 