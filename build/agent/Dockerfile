# syntax=docker/dockerfile:1
FROM golang:1.17 as runner-builder
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0 
ENV GOOS=linux

RUN cd cmd/agent;go build -o /runner -mod mod -a .;

FROM golang:1.17 as k6-builder
WORKDIR /build
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69; \
echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list; \
 apt-get update; \
 apt-get install k6; \
 go install go.k6.io/xk6/cmd/xk6@latest; \
 xk6 build latest \
 --with github.com/mostafa/xk6-kafka@latest \
 --with github.com/dgzlopes/xk6-kv@latest \
 --with github.com/grafana/xk6-redis \
 --with github.com/dgzlopes/xk6-redis@latest \
 --with github.com/walterwanderley/xk6-stomp; 


FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=runner-builder /runner /bin/runner
COPY --from=k6-builder /build/k6 /bin/k6
ENTRYPOINT ["/bin/runner"]
