FROM golang:1.22 AS builder
WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM hashicorp/terraform:1.5 AS runtime
COPY --from=builder /app/main /bin/terraform-cascade
ENTRYPOINT ["/bin/terraform-cascade"]
