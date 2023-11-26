FROM golang:1.19 AS build-stage

WORKDIR /app

COPY . .

RUN make

FROM gcr.io/distroless/base-debian11 AS release-stage

COPY --from=build-stage /app/bin/scheduler /scheduler
COPY .env .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/scheduler" ]