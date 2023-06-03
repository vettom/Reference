# Cloud Loadbalancer
- Software defined decentralised
- Global
	 - https
	 - SSL Proxy
	 - TCP Proxy
- Regional
	- Internal
	- Network LB
	- Internal Https

# Managed instance group
- Identical instances from templates (First step to create template)
- Can be resized
- ensure all are running
- Zone or regional
- Autoscaling

## Autoscaling policy
- Based on CPU
- LB capacity
- Monitorng metrics
- Queue based workload

## Https load balancing
- Global
- Anycast IP
- Http on port 80 or 8080
- Auto scaling, no prewarking
- URL mapping
- SSLtermination (up to 15 SSL certifciates)
- QUIC transport protocol

## Netwrok Endpoint Group
- Configuration object that specified a group of back end endpoints or service
- Ideal for containers   

# Cloud CDN
- Caching is decided by Cache modes.\
## Cache modes
- Use_origin_headers
- Cache_all_static
- force_cache_call

# TCP Proxy
- Global for unencrypted non-http traffic

# Choosing LB
- Only proxy LB supports IP v6
- LB can act as reverse proxy to server traffic
