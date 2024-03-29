# Resource-monitoring

# Workspace
Is root entity that holds monitoring and configuration infromation
Can monitor GCP and AWS in single workspace
To access AWS account, must create a project with credential
Workspace can monitor multiple projects.
metrics writer access is required for resources to write to monitoring host project workspace

## Dashboard
Grouping and filters allow to limit results
Outlier mode allow ading threshold, also like top 10

Monitoring metric writer role required for log writting
Log sync is used to aggregate logs at project, Folder or Org level. 


# Big Query and Data Studio
Exporting data to BigQ enables analysing data andvisualising in Data Studio
Data Studio provides visualisation with Data Studio.  Converts raw data to metrics.
Cloud Pubsub enable streaming logs to applications or endpoint

Debugger supports live snapshot of production app

# Logging 
- Collect
	Automatic logging on most. Logs are organised by project. Additional log parsing using fluentbit filters
- Analyze 
	Analyze data in realtime with logs exploere, pub/sub, dataflow,  BugQuery or archived logs from CloudStorage
- export 
	Export to cloudstorage or pub/sub or BigQ. Export logbased metrics to Monitoring
- Retain
	Retention from 1-3650 day (default 30days. Admin logs for 400 days. 

# Types of logs
- Cloud Audit logs
- Agent logs
- Network logs

# Application Performance management
- Debugger
	Can examin running workload. Debug session can be shared. Capture snapshots. Integrated with IDE's. Integrated with VCS
- Tracing
	Latency analysis. Track App engine, compute engine.  Automatic issue detection
- Profiler
   Low profile heap profiling. Supports Vm's, App engine, Compute, and GKE. CPU and Heap picture of application.


## Operations based tools
   - Monitoring 
   	- Dashboards
   	= Metrics Explorer
   	- Uptime checks
   	- alerting
   	- services health
   - Logging
   	- Logs explorer
   	- Logs Dashboard
   	- Log-based metrics
   	- Log Analytics
   - Error reporting
   	Reports crashes on system
   	Stacktrace also provided
   - Service Monitoring
   	Intra-service dependencies
   	Define and monitor SLO
   	Know when Error budget exceed

# Why monitor
- Continual improvements
	Capacity considerations, improve performance, improe security
- Dashboards
	Metrics in a summary view, automated detection, trend analysis
- Alerting
- Debugging
	Triggers, system outage, data loss, monitoring failure. Start with analysing signal data. Postmortom with documentation, root cause analysis

# Monitoring type
 - Blackbox : Monitoring externally like a customer would
 - Whitebox: Internal based on metrics
# Metrics must be SMART
- Specific
	like 95% response with in 10ms
- Measurable
- Achievable
- Relevant
- Time Bound
	Availability per year/month?


# Goldent signals
- Latency
	How much time to return the result. Page load latency, number of requests waiting, time till complete data return
-  Traffic
	Current system demand. Requests per seconds, cuncorrent sessions, Network IO, transactions, retrievals per sec etc
- Saturation
	How close to capacity systems are. %memory,CPu, threadpool, disk users etc
- Errors
	System failures and other issues. Failed requess, number of 400/500 errors

# SLI
	Measure of reliability. Number of good events/valid events
	## Ways to set SLI
	- Application level metrics
	- Logs processing
	- Synthetic/Client data
	- Client-side Instrumentation

# SLO
  reliability target. 

# SLA
  Agreement with customer on minimum level of service. Often have compensations attached to it. 

# Alerts
- Precision : relevant alerts/ Relevant alerts + Irrelevant alerts
- Recall : Relevant alerts/ Relevant alerts + missed alerts
- Detection time
- Reset time

# Service monitoring
- Summary of service status, SLI status etc

# Network Intelligence Center
Diagnose network connectivity issue, improve security and monitor real time network perfoamcne
- Connectivity test
- Performance Dash
- Network topology
- Fireall insights
- Network analyzer
