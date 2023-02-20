# Validation options

https://learn.hashicorp.com/tutorials/terraform/variables?in=terraform/certification-associate-tutorials#reach-skip-nav 
 Understand how validation is uased and handled. Here ensuring legth and type of input.
> error_message must be sentence start with Capital letter and must end with . or ?
```terraform
 validation {
    condition     = length(var.resource_tags["project"]) <= 16 && length(regexall("[^a-zA-Z0-9-]", var.resource_tags["project"])) == 0
    error_message = "The project tag must be no more than 16 characters, and only contain letters, numbers, and hyphens."
  }

  validation {
    condition     = length(var.resource_tags["environment"]) <= 8 && length(regexall("[^a-zA-Z0-9-]", var.resource_tags["environment"])) == 0
    error_message = "The environment tag must be no more than 8 characters, and only contain letters, numbers, and hyphens."
  }
```
### Validate variable starts with
 Variable validation can be introduce to verify variable follow certain conventions. Here in this example CIDR must start with 10.0
```terraform
    variable "CIDR" {
    validation {
     condition = can(regex("^10.0", var.CIDR))
     error_message = "Cidr must start with 10.0."
    }
}
```



## slice function
  Used to extract field from the list. Accepts 2 arguments start and end element. Start element is inclusive and end not.
```terraform
  slice(var.private_subnet_cidr_blocks, 0, 2)
```
length() : to check length of values

