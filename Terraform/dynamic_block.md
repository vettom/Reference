# Dynamic Block
 Special kind of block to iterate through repeatable configurations.
> For list had to assign iterator, for map could possibly use  DMZ_Ingress.port["toport]
 
```terraform
  dynamic "ingress" {
    iterator = port
    for_each  = var.DMZ_Ingress
    content {
      from_port = port.value 
      to_port = port.value 
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }
```