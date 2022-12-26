FROM golang:1.18-alpine
WORKDIR /app
COPY . .
COPY ./go.* .
RUN go mod download
COPY . .

EXPOSE 8000

ENTRYPOINT [ "make", "run" ]