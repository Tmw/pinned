FROM golang:1.11.1 as builder-backend
RUN mkdir /app
ADD ./backend /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -o pinned-backend .

FROM node:10.8.0 as builder-frontend
RUN mkdir /app
ADD ./frontend /app/
WORKDIR /app
RUN npm install
RUN npm run build

FROM alpine:latest
RUN apk add ca-certificates
RUN mkdir /app
COPY --from=builder-frontend /app/build/ ./app/static
COPY --from=builder-backend /app/pinned-backend ./app/pinned-backend
WORKDIR app
EXPOSE 4000
CMD ["./pinned-backend"]
