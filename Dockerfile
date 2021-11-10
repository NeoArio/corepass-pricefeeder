


# We specify the base image we need for our
# go application
from  golang:1.15.6-buster
# We create an /app directory within our
# image that will hold our application source
# files

RUN mkdir /app
RUN mkdir /etc/mediagateway


# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app

ENV XCB_NODE_URL http://127.0.0.1:8545 \
    NETWORK_ID 3
    PRIVATE_KEY_BYTES test_key
    PRICE_FEED_CONTRACT_ADDRESS ab47780c6900023b87dc6bfb66214c11ddcaae54205d
    CTN_PRICE_API_ADDRESS https://stg.pingextest.eu/marketdata/instruments/btc_usd/history

# we run go build to compile the binary
# executable of our Go program
RUN make build
# Our start command which kicks off
# our newly created binary executable

ENTRYPOINT ["/app/pricefeederd"]
