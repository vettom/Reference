# Functions or modules are called in format Name(Arg1, arg2)
Function works with existing files before start only. Functions does not participate in dependency graph, so dynamically created files cannot be used in operation
> TF does not support user defined functions.


## templatefile function
 Reads file at a given path and renders its content as template using a supplied set of variables 
```terraform
  templatefile(path, vars)
user_data = templatefile("user_data.tftpl", { department = var.user_department, name = var.user_name })
Including template file and passing values for variables in the file. Variable name and department is used in file, value is from variable defined in TF. 
```
 vars argument must be an argument. With in template file each keys in map is available as variable
 .tftpl is recomended extention. 
## Loop 
```terraform
 variable "users" {
  type = list(string)
  default = ["Freya", "Maria", "Jenn", "Maria" ]
}

output "names" {
  value = [for Name in var.users : lower(Name) ]
}
# use toset(var.users) to get unique list
```

## Use Count
 Count has special meaning in TF and it keeps an index. Inde is used to access element in list.
```terraform
  resource "aws_iam_user" "aws_iam_user_01" {
    count = length(var.users)
    name = var.users[count.index]
  }
```
## Lookup function 
 Retrieves value of a single element given key.
```terraform
 variable "aws_amis" {
  type = map
  default = {
    "us-east-1" = "ami-0739f8cdb239fe9ae"
    "us-west-2" = "ami-008b09448b998a562"
    "us-east-2" = "ami-0ebc8f6f580a04647"
  }
}

AMI = lookup(var.aws_amis, var.aws_region) Retuns value of aws_region in the variable.
```
## Try 
  Evaluates all of its argument expresstions and returns the result of first one that does not produce error.
```terraform
    variable "name" {
    type = map(string)
    default = "Denny"
}
try(var.name, "NoName")
```
## file
 file("path") file("/tmp/file1") Read contents of file 
## Eelement 
 element(list, index ) return value of element at that index

## values 
values(map) : returns value of all items in the map as a list.
## flatten
flatten(list1, list2) Combines multiple lists in to one.

## Numeric functions
 - max(3,2,62,2)  return 62
 - min(3,2,62,2)  Return 2
 - abs(-3.2) = 3.2  Returns absolute +ve value 
 - ceil(3.2) = 4  Returns closest whole number that is greater than or equal
 - floor(3.2) = 3
 - 
## String functions
 - chomp("Hello\r \n") = Hello Remove carriage return
 - format("Today is %s", Tue) = Today is Tue like print command string substitution

## Filesystem functions
- abspath(path.root) current absolute path
- 

```terraform
 NET = $(contains(var.list1, "prod") ? var.prodnet : var.devnet )
 #If var.list contains element prod. 
```