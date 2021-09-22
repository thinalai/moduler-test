FROM golang:1.17 as build
WORKDIR /go/src/task6
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -v -o ./bin/task6

FROM gcr.io/distroless/static:latest
COPY --from=build /go/src/task6/bin/task6 /app
ENTRYPOINT [ "/app" ]
