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


# Deployments
- Blue-green
	Complete new environment, and switch with service selector label
- Canary deploy
	Portion of traffic diverted by using same label and managing number of pods. Not best without tool like Istio
- A/B testing
  Portion of traffic based on routing rule like Http header, geo location, etc. Ideal to measure effectiveness of functionality. For eg: new features, and select target customer
- Shadow testing
	Deploy and run running along side but hidden. Incoming requests are mirrored and replayed in mirrored env. Must ensure shadow test do not trigger side effect. Allows testing of new features with full traffic without impacting production.

# Rollout
	kubectl rollout manages deployment of Deployments, daemonsets or Stateful sets
	kubectl rollout undo to rollback to prev version. 
	By default previos 10 version information retained

	kubectl scal --replicas to scale deployment

> Use session affinity in LB if switching between different version can cause issue.
> Volume: DownwardAPI enables container to  identify pod it is running. Useful if container app should be uniq in cluster.
