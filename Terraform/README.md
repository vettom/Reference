# Learning Terraform

## Files and extensions
```bash
    .tf  : Terroaform files
    .tf.json : Json format files
    terraform.tfstate : State file
    .terraform.lock.hcl  : File containing version of provider to use. 
                          If lock.hcl not present, letest version or as per .tf file downloaded
                          If version in tf above lock file, init wi ll fail with message requiring upgrade
                          Always good idea to include lock.hcl to GIT
    terraform.tfvars or *.auto.tfwars  : Variables store. Will load automatically. option to add custom by -var-file 
    .terraform : Folder for configuration
    _override.tf.json : Override files by default. Override is merged other appended.
    $HOME/.terraform.d/credentials.tfrc.json  : API token stored in plain text after auth with TF Cloud.
    .terraform/providers/registry.terraform.io/hashicorp/  : Location of providers. Each version downloaded stored here.
``` 
## Providers
  - Provider information/block can be globally defined
  - multiple providers can be configured
  - providers must match prefix eg aws_instance match aws provider.
  - provider must provide list of supported resources to terraform
  - each provider must be initialised before it can be used.
  - multiple provider instances can be configured to support different regions.
  - Aliases can be used to differentiate providers.
  Provider plugins can be shared with by setting plugin sharing.
#### Provider Lock
  Provides a mechanism to update .terraform.lock.hcl file with required versions, hash keys, checksum etc. Checksum verification is *trust on first use*
When a new provider is added and initialised, plugin is downloaded and trusted. This can be challenging if not installing the provider from TF registry or from private registry.
terraform provider lock lets you pre calculate the checksum and add to lock file.
```terraform
    terraform providers lock  # Updates lock file with checksum from official registry,
    terraform providers lock -fs-mirror=dir  # Refresh from local mirror.
                             -net-mirror=url # Update from mirror URL
```

#### Upgrading provider
  ```bash
      terraform init -upgrade   : Allows upgrade of provider overriding lock.hcl
  ```

## Securing keys
- Variables
- AWS CLI
- External vault provider.

### CLI configuration file and auth
  .terraformrc file of TF_CLI_CONFIG defined configuration file. It follwos .tf syntax
- credentials : Configure credentials to use with TF cloud enterprise etc
- credential_helper : External helper programe for storage and retrieval of credentials
- plugin_cache_dir : Enables plugins caching and sets location.
```terraform
    {
  "credentials": {
    "app.terraform.io": {
      "token": "2cxvSkOBS05Nfdfsdgwrrwgc3dzF5g3BgxxQmy3U"
    }
  }
}
```
## Sharing Plugins 
 .terraformrc file provides option for sharing plugins and storing in separate location. Set location in file or TF_PLUGIN_CACHE_DIR
```terraform
    plugin_cache_dir = "$HOME/.terraform.d/plugin-cache"
```
## Variables
  Veriables are defined with key var.x. A Map is key value store of variable
  var.x    : gets value of variable X
  var.MAP["KEY"]  : Value of Key
  ${var.LIST}     : List variable


### Interpolation syntax
   Interpolcation syntax can accept values from other variables, perform evaluation, match expression etc
   Conditionals can be call to functions.

```bash
    access_key = ${var.aws_access_key}
    ${Count.index + 1 }
    subnet = ${ var.env == "prod" ? var.prod.subnet : var.dev.subnet }
```

## Configuration file loading
  - Files are loaded in alphabetical order. 
  - Must have .tf or .tf.json extention. 
  - can have both json and tf syntax file in same folder

Overrid files are loaded after non-override files again in alphabetic order
- loaded configurations are appended, not merged. 
- Because of not merge, 2  resource with same name will result in validation error
> Overrid files are MERGED

* Override can be added to file that ends with override.tf or override.tf.json

# Output
  Provides ability query information after apply or refresh operations
  - Only single output variable allowed in single output block
```bash
    output "NAME" { value = ${VALUE} }
    output "aws_instance_dns_name" { 
      value = ${aws_instance.webserver.public_dns} 
    }
```
- Returning and storing sensitive information
Will output as "sensitive" however data is stored in the state file. Can be retrieved using query
```bash
    output "aws_instance_dns_name" { 
          value = ${aws_instance.webserver.public_dns} 
          sensitive = true  
        }
    terraform output #Will show all the outputs including sensitive.
```

## Resources configuration
 Resources are defined as per the requirements of resource definition.
 ```bash
    resource <resource type> <resource name {
      <resource properties>
    }

 ```
### Meta parameters or Arguments
  every resource has some meta parametes. Some changes in newer versions
  - depends_on  handles hidden resource or module dependancy
  - count  Number of resources to create. (in new version available for modules as well)
  - for_each Loop through number 
  - provider defined which provider to use for the resource
  - lifecycle  define resource behaviour, like create_before_destroy. ignore_changes, prevent_destroy

## Datasource
  Is configuration fetched from external Data source or even Terraform instance. Each data source is mapped to provider based on *longest prefix matching*. 
  ```bash
      data <data source name> <name of the data type>{}
      data "aws_availability.zones" "available" {}
  ```

## Dependencies
 - Implicite : default by AWS order like IP after instance creation
 - Explicit  : User defined like S3 before instance creation
 - Nondependant : Can be created in parallel

## Common Tags
    AWS provides a data block called default_tags which can be used to apply tags to all instances created
```terraform
    default_tags {
    tags = {
      Managed_by = "Terraform"
    }
  }
```

## Merge Tags 
    Variables can be defined in Maps and use merge to combine it all.
```terraform
    tags = merge({
    "Name" = "vpc01"
    },
    var.vpc_tags,
    var.owner_tags)
```

 


 
