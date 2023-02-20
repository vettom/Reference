# Variables

### Ways of providing input variables
 - As input argument -var "key=value"
 - In .tfvars ot auto.tfvars  these files are processed automatically
 - Provide var file as argument -var-file="$PATH/variables.tf"
 - Environment variables starting with TF_VAR_ 
 - Via UI, if not specified TF will prompt. 
 > Vriables in root module can be set via CLI or env variable. Variables set in module MUST be passed as argument in module block.
## Workspace variable
```terraform
  Workspace = terraform.workspace  #Assign variable with workspace name.
  Workspace = "${ terraform.workspace == "default" ? "${terraform.workspace}-Prod" : "Dev" }"
```

# variable is defined with variable keyword
```terraform
    variable "var1" {}
    # String variable
    variable "vpcname" {
        type = string
        default = "myvpc"
}
# Intiger variable  
    variable "sshport" {
        type = number
        default = 22
}
# Boolian variable
    variable "enabled" {
        default = true

}
# List vriable
    variable "mylist" {
        type = LIST(string)
        default = [ 1,2,3,4]

}
# MAP variable where key value store.
    variable "mymap" {
        type = MAP(string)
        default = {
          key1 = value1
          "key2" = "value2"
        }

    }
# Variable with expression
    variable NET = ${var.env == "prod" ? var.prod_net : var.dev_net}
```
## Interpolation of variable can be skipped with $$ eg  $$var.id becomes ${var.id}

# Local Variables
  These variables are local to the module itself. Can have any number of variables.
  To use local variable, prefix it with local
  ```python
        locals{
            var1 = "Valid for this module only."
        }

        Var2 = ${local.var1}-instance
  ```
 #Data types
## Tuple 
  > Allows you to specify multiple data type in single variable unlike list
```terraform
variable "mytuple" {
  type = tuple([string, number, string])
  default = ["cat", 3 , "dog"]
}
```
## Object
 Complex map that allows multiple data types.
```terraform
  variable "myobject" {
  type = object({name = string, port = lisr(number)})
  default = {
    name = "dv"
    port = [22,80]
  }
}
```