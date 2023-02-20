# Modules
 I is a folder containing terraform files. Recommended to have 3 files minumum. main.tf, variables.tf, outputs.tf
> Modules must be imported with get -upgrade=true option for to fetch updates. 
>  tf init had -upgrade. Version available for TF cloud and registry only.
## Importing module
```shell
    terraform get ./module1   #Updates once, does not check for changes or upgrade
    terraform get -update ./module1  #Will downloade latest version always
    terraform init -upgrade # Notice it is upgrade not update for init.
```

## Including a module.
 ```bash
     module "module1" {
        source = "./folder1"
        version = "1.2.2"
        var1 = var.vpcid
        var2 = "something"
     }
 ```

# Refactoring
 Instance a is now instance b
```terraform
    moved {
     from = aws_instance.a
     to   = aws_instance.b
    }
```