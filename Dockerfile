FROM golang:1.11.1 as builder-backend

RUN mkdir /app
ADD ./backend /app/
WORKDIR /app

RUN CGO_ENABLED=0 go build -o pinned-backend .

CMD ["/app/pinned"]

FROM node:10.8.0 as builder-frontend
run mkdir /app
add ./frontend /app/
workdir /app

run npm install
run npm run build

FROM alpine:latest
run apk add ca-certificates
run mkdir /app
copy --from=builder-frontend /app/build/ ./app/static
copy --from=builder-backend /app/pinned-backend ./app/pinned-backend

workdir app
cmd ["./pinned-backend"]

