FROM golang:1.11 as source
WORKDIR /code
COPY . /code

FROM golang:1.11 as build
WORKDIR /build
COPY --from=source /code/pkg /build
RUN CGO_ENABLED=0 GOARCH=amd64 go build -mod vendor -o app

FROM scratch
COPY --from=build /build/app /
ENTRYPOINT ["/app"]
