FROM golang:1.15 AS builder
WORKDIR /go/src/better-console-backend
COPY . .
RUN go mod download
RUN go install -ldflags '-w -extldflags "-static"'

# make application docker image use alpine
FROM alpine:3.10
# using timezone
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul
RUN apk add -U tzdata

WORKDIR /go/bin/
# copy config files to image
COPY --from=builder /go/src/better-console-backend/config/*.json ./config/
# copy execute file to image
COPY --from=builder /go/bin/better-console-backend .
EXPOSE 2021
CMD ["./better-console-backend"]
