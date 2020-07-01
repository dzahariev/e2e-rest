FROM golang AS localenv
RUN apt-get update && apt-get install -y wait-for-it make && apt-get clean
WORKDIR /go/src/github.com/dzahariev/e2e-rest/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /main main.go

FROM scratch AS release
WORKDIR /app
COPY --from=localenv /main /app/
ENTRYPOINT [ "./main" ]
EXPOSE 8080