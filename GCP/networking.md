# Netowrking

# VPC network types
- default : Every project. Once subnet per region
- Auto mode : 
	Regional IP allocation
	Fixed /20 subnetwork per region
	expandable up to 16
- Custom mode


- networks do not have IP ranges
- Auto ode subnet can be converted to custom mode
- Subnets span across zones in the region

# IP address
- Internal DNS name [hostname]-[zone].c.[]roject].internal

# Routing and Firewalls
- Routes trffic between subnets with in an vetwork
- lso default route for external
- Can create spcial route

- Firewall rules must allow connection between network
- Firewall rules are applied across network
- Connections can be managed at instance level

- Private Google access is configured per subnet.
- VM's can  live migrate and is default option for VM during maintenance

# Cloud interconnect (Direct connect equiv)
Cloud Interconnect provides low latency, high availability connections that enable you to reliably transfer data between your on-premises and Google Cloud Virtual Private Cloud (VPC) networks.
- Dedicated Interconnect provides a direct physical connection between your on-premises network and Google's network.
- Partner Interconnect provides connectivity between your on-premises and VPC networks through a supported service provider.

# Cloud VPN
Cloud VPN securely extends your peer network to Google's network through an IPsec VPN tunnel. Traffic is encrypted and travels between the two networks over the public internet.
- For low volume data connecitons
- Requires VPN gateway onpremise
- Routed over internet
- Must establish 2 tunnels
- MTU cannot be greater than 1460
- Supports
	- Site-to-site VPN
	- Static Routes
	- Dynamic Routes (Cloud Router)
	- IKEv1 and V2 Cihers
# HA VPN
- 99.99 SLA
- 2-4 tunnels required
- GCP choses 2 IP addresses
- Must use dynamic routing
- Supports active-active or active-passive
- Can configure with AWS VPG 
- Can be used to connect 2 Google cloud as well

## Cloud router
Google Cloud Router enables dynamic route updates between your Compute Engine VPN and your non-Google network. Cloud Router eliminates the need to configure static routes and automatically discovers network topology changes.

# Configuring VPN HA
- Create Cloud VPN Gateway at each end (or on prem gateway)
- Create a cloud Router in each network
- Create Tunnel one per interface. 2 required for HA 99.99 SLA. Onpremised can be 2 device or 2 network interfaces.
- Create Router interface for each router interface
- Create BGP peer for each interface
- Firewall rule to allow connection between netowrks
- Enable global routing if multi-region connection required.

> HA VPN is a regional resource and cloud router that by default only sees the routes in the region in which it is deployed. To reach instances in a different region than the cloud router, you need to enable global routing mode for the VPC. This allows the cloud router to see and advertise routes from other regions.

# Interconnect
## Layer 3
- Direct peering. Dedicated line. Create connection to supported co-location router and BGP enabled.
- Careeier Peering (via carrier network)
## Layer 2
- Dedicated Interconnect
- Partner Interconnect

# Peering
For access to Google services
 - No SLA
 - Access via google public IP address
 - Dicect peering require access to edge point

# Shared VPC
- Allows multiple project to access resources in common netowrk


