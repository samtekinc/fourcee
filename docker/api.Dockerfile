FROM --platform=${BUILDPLATFORM} golang:1.19-alpine as build

COPY ./ /build

ARG TARGETOS
ARG TARGETARCH
RUN cd /build/; CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/api ./cmd/api;

FROM alpine:latest
RUN apk add --update --no-cache git
COPY --from=build /bin/api /bin/api

CMD ["/bin/api"]