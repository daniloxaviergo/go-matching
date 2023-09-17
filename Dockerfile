FROM golang

# Set destination for COPY
WORKDIR /app

COPY ./ /app

RUN go mod download
