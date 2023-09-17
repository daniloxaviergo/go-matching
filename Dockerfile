FROM golang

WORKDIR /app
COPY ./ /app
RUN go mod download
