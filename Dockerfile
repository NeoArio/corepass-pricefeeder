


# We specify the base image we need for our
# go application
from  golang:1.15.6-buster
# We create an /app directory within our
# image that will hold our application source
# files

RUN apt update -y && apt upgrade -y

# set if needed!
ARG HTTP_PROXY=""
ARG SWAGGER_VERSION=v0.27.0
RUN set -ex \
 && export https_proxy=${HTTP_PROXY} http_proxy=${HTTP_PROXY} \
 && curl -o /usr/local/bin/swagger -L https://github.com/go-swagger/go-swagger/releases/download/${SWAGGER_VERSION}/swagger_linux_amd64 \
 && chmod +x /usr/local/bin/swagger \
 && unset https_proxy http_proxy


RUN mkdir /app
RUN mkdir /etc/mediagateway


# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app


ENV SWAGGER_PORT 7000

RUN swagger generate spec -o ./swagger.json

# we run go build to compile the binary
# executable of our Go program
RUN make build
# Our start command which kicks off
# our newly created binary executable

ENTRYPOINT ["/app/pricefeederd"]
