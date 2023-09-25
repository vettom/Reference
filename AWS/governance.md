# Governance

## AWS Organisations
    - Logging account
    - Create and destroy AWS accounts
    - Reserved instance discount across accounts
    - Single billing account
    - Service control policies to limit user permission

## Sharing resourcer with RAM
 Resource Access Manager allows sharing resource with other accounts.
 Share following
 - Transit gateways
 - VPC Subnets
 - Licenses
 - Route 53 resolver
 - Dedicated hosts 
 - etc 

## Cross account role access
  Single credential can assume role in another account
  In Primary account
  - Create an IAM role (Another AWS account)
  - Define Permissions. (s3, c2 etc)
  - Define groups/user allowed to assume role.

## Inventory - AWS Config
    - Aloows to query resource
    - Enforce rule and automatically fix (eg: Public s3 access deny)
    - Track changes, when it changed, who changed etc

## Directory Service
 Fully managed AD. 
 - Microsoft AD
 - AD Connector. Tunnel between AWS and on-prom AD
 - Simple AD powered by Samba

## Cost and Budget
    - Cost Explorer. Visualise costs besed on various factors including tags
    - AWS Budget
        - Cost Budget. How much spendign
        - Usage budget. How much are we using
        - Reservation budget.
        - Savings plan budget

## Optimising Costs
    - AWS CUR (Cost and usage report) 