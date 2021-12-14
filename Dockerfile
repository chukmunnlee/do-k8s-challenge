ARG TAG=1.17-alpine

FROM golang:${TAG} AS builder

ARG APP_DIR=/app

ENV CGO_ENABLED=0

WORKDIR ${APP_DIR}

COPY src src

WORKDIR ${APP_DIR}/src
RUN go build -o dns-edit *.go

FROM golang:${TAG}

ARG APP_DIR=/app
ENV PORT=8443

WORKDIR ${APP_DIR}

COPY --from=builder ${APP_DIR}/src/dns-edit dns-edit

ENTRYPOINT [ "./dns-edit" ]

CMD [ "-port", ${PORT} ]
