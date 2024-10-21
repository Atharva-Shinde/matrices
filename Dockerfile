FROM quay.io/fedora/fedora-minimal AS builder

RUN microdnf install -y golang

WORKDIR /app

COPY . .

# CGO_ENABLED=0 gives statically linked binary	
RUN CGO_ENABLED=0 go build -o main .

FROM quay.io/fedora/fedora-minimal

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]