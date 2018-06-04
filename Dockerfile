FROM golang:1.10 as tester
LABEL tester=true
WORKDIR /go/src/github.com/operator-framework/operator-manifests
COPY . .
RUN go test ./cmd/catalogbuilder
RUN go get github.com/operator-framework/operator-lifecycle-manager/cmd/validator
RUN validator ./manifests

FROM golang:1.10 as builder
LABEL builder=true
WORKDIR /go/src/github.com/operator-framework/operator-manifests
COPY cmd cmd
RUN go build -o bin/catalogbuilder ./cmd/catalogbuilder
RUN chmod +x bin/catalogbuilder

FROM scratch
LABEL main=true
WORKDIR /
COPY --from=builder /go/src/github.com/operator-framework/operator-manifests/bin/catalogbuilder /catalogbuilder
COPY manifests /manifests
COPY operator-manifests.catalogsource.yaml operator-manifests.catalogsource.yaml
COPY operator-manifests.configmap.yaml operator-manifests.configmap.yaml

CMD ["/catalogbuilder"]
