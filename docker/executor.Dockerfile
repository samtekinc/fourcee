FROM --platform=${BUILDPLATFORM} golang:1.19-alpine as build

COPY ./ /build

ARG TARGETOS
ARG TARGETARCH
RUN cd /build/; CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/tf-execution-cli ./cmd/tf-execution-cli;

FROM alpine:latest
RUN apk add --update --no-cache git
COPY --from=build /bin/tf-execution-cli /bin/tf-execution-cli
RUN mkdir -p /efs

CMD ["/bin/tf-execution-cli"]