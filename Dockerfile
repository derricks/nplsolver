# Creates a docker image for the nplsolver binary in this repo
#
FROM ubuntu:14.04
MAINTAINER derrick.schneider@gmail.com

# Build from source so that the binary is correct for the platform
RUN apt-get update && apt-get install -y golang
RUN mkdir -p /tmp/src/nplsolver
ADD "." "/tmp/src/nplsolver"
RUN export GOPATH=/tmp && cd /tmp/src/nplsolver && go build nplsolver.go

RUN mkdir -p /var/nplsolver
RUN mkdir -p /etc/nplsolver
RUN cp /tmp/src/nplsolver/nplsolver /var/nplsolver

# Add files
ADD "solver.properties" "/var/nplsolver"
ADD "data/enable1.txt" "/etc/nplsolver"

# Configure
EXPOSE 15432

RUN cd /var/nplsolver
