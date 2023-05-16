FROM golang:1.19 as dev
ENV LINE_APP_ROOT=/go/src/app
COPY docker/conf/air.toml $LINE_APP_ROOT/air.toml
COPY src/server/go.mod src/server/go.sum $LINE_APP_ROOT/src/server/
COPY lib/go.mod lib/go.sum $LINE_APP_ROOT/lib/
WORKDIR $LINE_APP_ROOT/src/server
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "../../air.toml"]
# CMD ["sleep", "7200"]
