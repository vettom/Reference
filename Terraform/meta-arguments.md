# Meta Argumets
 Meta argumentes defined in any module have special meaning.
 
## count
 Creates multiple instances of module from single block. count.index provide indexed value. Count value must be known before apply.
 ```terraform
    resource "aws_instance" "Web" {
  ami = "am-xxx"
  count = 4
  tags = {
    Name = Webserver-${count.index}
  }
}
output "servers" {
  value = [ for X in aws_instance.Web[*].tags["Name"]: X]
}
```
## for_each loop
 > Can process a list or map objects. If List item, it must me converted to set. Value of Key must be know, and not dynmic. Cannot work with sensitive data.
### Process list
```terraform
        variable "CIDR" {
         type = list(string)
         default = ["10.1.0.0/16", "20.10.0.0/16" ]
       }
       resource "aws_vpc" "aws_vpc_01" {
         for_each = toset(var.CIDR)
         cidr_block = each.value
       }
```
### Processign map file. 
```terraform
    variable "CIDR" {
    type = map(string)
    default = {
      One = "10.1.0.0/16", 
      Two = "20.10.0.0/16"
    }
  }
  resource "aws_vpc" "aws_vpc_01" {
    for_each = var.CIDR 
    cidr_block = each.value
    tags = { 
      Name = each.key
       }
  }
```
## Providers

