FROM golang:1.19.2-bullseye
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /godocker
EXPOSE 8080
CMD [ "/godocker" ]
