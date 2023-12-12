FROM golang

ENV APP_HOME /go/src/gabrie30/hello-world
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
COPY . .
RUN go build .
EXPOSE 8080
CMD ["./hello-world"]
