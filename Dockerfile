FROM golang:buster as builder
COPY . /go/wimp/
WORKDIR /go/wimp
# we build a go binary but we cannto use CGO since we use C functions to drop the priviledges
RUN go get github.com/ipinfo/go-ipinfo/ipinfo && GOOS=linux go build -ldflags '-extldflags "-D_FORTIFY_SOURCE=2,-static,-Wl,-z,noexecstack,relro"' -o wimp wimp.go
FROM busybox
# No CGO  means copy glibc library, ugly but works
RUN mkdir -p /lib/x86_64-linux-gnu /lib64
COPY libpthread.so.0 /lib/x86_64-linux-gnu/
COPY libc.so.6 /lib/x86_64-linux-gnu/
COPY ld-linux-x86-64.so.2 /lib64/
COPY --from=builder /go/wimp/wimp /wimp
EXPOSE 81
ENTRYPOINT ["/wimp"]
