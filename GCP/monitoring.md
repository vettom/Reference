# Resource-monitoring

# Workspace
Is root entity that holds monitoring and configuration infromation
Can monitor GCP and AWS in single workspace
To access AWS account, must create a project with credential

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

