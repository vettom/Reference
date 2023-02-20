# Provisioner
Supports many provisioners, local as well as remote
```terraform
    provisioner "local-exec" {
  command = "echo ${aws_vpc.id}"
}
```

### Chef 
 Installs, configures and run chef provisioner on remote host. supports ssh as well as winrm
 
### file 
  Used to copy file from TF server to instance vi ssh or winrm
```terraform
provisioner "file" {
  source = "/path"
  destination = "path"
}
```

### remote-exec
  execute a script or command on remote host. Generally used to trigger config manager tasks. To execute any script, use file to copy first
> Inline executed in order, result of last command only evaluated to on_failure=error
- inline  : Tye commands to execute here
- script  : Local script to be copied and executed
- scripts : path to scripts to be copied and executed.
```terraform
        provisioner "remote_exec" {
          inline = [
            "/tmp/script.sh"
            "yum install -y httpd"
          ]
        }
```
### local-exec
  execute something on local TF server
  
### salt-masterless
  Provisions host built by TF using salt states without connecting to Salt server
  
### null_resource 
  Allows configuration of provisioners that are not associated with single existing instance.
  

# Provisioners type
 - Are scripts used as part of creation or distruction
 Use cases
 - Bootstrap an instance
 - Clean up before destroy
 - Run configuration management etc
 - can be used on local or remote
  By default provisioners run during provisioning. It will not execute on update or refresh.
 Destroy time provisioners must specify when == "destroy" with in provisioning block
If removed from configuration, destroy provisioner will fail. And at next apply Terraform will try to run it again.
For this reason destroy provisioners must be safe to run multiple times. During destroy, destroy-provisioners are
executed before removing resource.

## Provisioner Fail
  Failed provisioners can leave host in semi configured state. Failed resources will be recreated at next apply cycle.
  Provisioning is not retried.
  ```terraform
    on_failure="fail or continue"
```