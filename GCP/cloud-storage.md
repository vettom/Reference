# Storage
- Cloud Storage (Bucket)
	- standard
	- Nearline like once a month access
	- Coldline (access once in 90 days)
	- Archive storage (1s year)
- Cloud SQL
- Cloud Spanner (SQL)
- Firestore (NAS, NFS)
	- Data stored as documents and collections.
- Cloud Bigtable (Big Data)
	- No SQL
	- Semi-structured or structured data
	- Machine learning data
	- Analytical structured data

# Access Control
- IAM
- ACL : Define who has what access and level of access. Contains Scop (who)  and permissions
- Signed URL : Grant limited time access token that can be used by any user. Anyone with URL can invoke permitted operations. Time limited
- signed Policy Document 

# Data import service
- Transfer Appliance : Rack,capture and ship to google
= Storage transfer service : Online data transfer
- Offline Media Import: 

# Managed DB services
## Cloud SQL
- Fully managed MySql, Porstgres erc

## Cloud Spanner
- Combines the benefits of relational DB with non-relational horizontal scale
- scale to petabytes
- High HA and consistenct
- Used for financial and inventory applications

## Firestore
- No-SQL document db
- used for Mobile, web, and Iot
- ACID transaction,  ensure transaction is success or fail.

## Cloud Bigtable
- No SQL 
- Operational and analytical 
- Integrates with big data tools
- High volume and fast

## Memory store
- Managed redis
- High availability and failover
- sub-millisecond latency



