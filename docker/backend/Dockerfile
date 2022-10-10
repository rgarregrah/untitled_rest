FROM golang:1.19-bullseye

WORKDIR /app

COPY ./backend .
RUN go mod tidy

# CMD [ "go", "run", "main.go" ]