FROM "golang:1.11-alpine"

WORKDIR /go/src/github.com/faithanalog/poka-pi-ilo-pi-sitelen-pona
COPY . .
ENV GOBIN=/go/bin
RUN go install ./cmd/sitelen-pona-api-server/sitelen-pona-api-server.go
CMD /go/bin/sitelen-pona-api-server
