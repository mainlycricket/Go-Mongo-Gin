### build

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \    
    go mod download

COPY . .

### dev build

FROM builder AS dev-build

COPY . .

RUN go install github.com/cespare/reflex@latest

CMD ["reflex", "-r", "\\.go$", "-s", "--", "sh", "-c", "go build -o bin/ ./cmd/server && bin/server"]

### prod build

FROM builder AS prod-build

COPY . .

RUN go build -o /app/bin/server /app/cmd/server/

### deploy

FROM scratch AS deploy

COPY --from=prod-build  /app/bin/server /app/server

ENV GIN_MODE=release
EXPOSE 8080

ENTRYPOINT [ "/app/server" ]
