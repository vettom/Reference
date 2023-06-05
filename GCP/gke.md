# Containers
- Is Isolated user spaces
- No kernel
## Uses multiple 
- Processes. Starts nd stops processes
- Linux namespace. Limits visibility
- cgroups: managed resources consumption
- Union file systems

# Cloudbuild
Used for building container
- Can retrieve source code for your build from various locations

# Migrate for Aunthos
- moves and convert workloads  in to containers
- Workloads can start as physical or VMS
- Data can be migrated all at once or streamed

# Migration
- create processing cluster
- migration source
- generate and preview plan
- generate artifacts
- test
- deploy

# Building small image
- Start with small base image and add necessary

## Builder pattern
- Compile application in a container with compiler
- Start with small image
- Copy compiled binary file from build container 
- Install Cacerts  
- Remove apk cache in /var/cache/apk

```bash
FROM golang:alpine as build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp

FROM alpine
# Install CA cert and remove cache
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk
# Copy just binary from build container
COPY --from=build-env /app/goapp /goapp
EXPOSE 8080
ENTRYPOINT ./goapp

```

# Google container registry
- Uses global caching of common images to speed up upload
- Automated vulnerability scanner

# Google container builder
- Container  build tool