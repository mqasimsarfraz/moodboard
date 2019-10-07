FROM golang:1.13.1-alpine3.10 as go-builder

ENV PACKAGE github.com/MQasimSarfraz/moodboard
ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/$PACKAGE

# create directories for binary and install dependencies
RUN mkdir -p /out && apk --no-cache add git

# copy sources, test and build the application
COPY . ./
RUN go vet ./...
RUN go test --parallel=1 ./...
RUN go build -v -ldflags="-s -w" -o /out/moodboard ./cmd/moodboard


# build the final container image
FROM alpine:3.10

EXPOSE 3080

COPY --from=go-builder /out/moodboard /

ENTRYPOINT ["/moodboard"]
