# Hierarchy

Org ID
|-Folders
  |- Projects
  	| Resouce

- Projects are associated with single billing account
- Permissions are applied top to bottom. Less restrictive access at Org level can override restrictive permission on a resource


# Resource Hierarchy
- Global : Images, snapshots, network
- Zonal : External IP
- Regional : Instances, disks

# Quotas
Default quotas are configured to stop run away resources 
- Rate quota
- Allocation quota 

# Tags and Labels
Tags are usually assigned to instances only and used for netowrking/firewall
Labels are filtered thrugh to billing as well

# Projects
