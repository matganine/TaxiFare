FROM ubuntu:18.04 AS build
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y wget
RUN cd /tmp && \
    wget https://dl.google.com/go/go1.15.7.linux-amd64.tar.gz && \
    tar -xf go1.15.7.linux-amd64.tar.gz && \
    mv go /usr/local
RUN mkdir -p /app/taxi-fare/
WORKDIR /app/taxi-fare/
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
ENV GOPROXY="https://proxy.golang.org"
RUN apt-get update && \
    apt-get install -y git && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
ENV CGO_ENABLED 0
COPY ./ ./
RUN go mod vendor
RUN go mod tidy
RUN go build -mod=vendor .

FROM ubuntu:18.04 AS release
LABEL maintainer="ralph.romanos@gmail.com"
WORKDIR /usr/local/bin/
COPY --from=build /app/taxi-fare/TaxiFare .
COPY data data/
RUN apt-get update && \
    apt-get install -y ca-certificates mtr && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
ENV TF_DATA_PATH=/usr/local/bin/data/
ENV TF_RIDES_FILE=rides.json
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/TaxiFare"]
CMD ["-v", "1", "-logtostderr"]