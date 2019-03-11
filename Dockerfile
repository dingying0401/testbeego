FROM golang:1.12
WORKDIR /go/bin/testbeego
COPY . /go/bin/testbeegio
EXPOSE 33333
ENTRYPOINT ["./testbeego"]
