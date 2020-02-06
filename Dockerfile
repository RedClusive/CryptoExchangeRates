FROM golang:latest

WORKDIR $GOPATH/src/github.com/RedClusive/ccspectator
COPY . .

RUN go get -d -v .
RUN go build -o /app .

ENTRYPOINT ["./entrypoint.sh"]
CMD ["/app"]


