FROM golang

WORKDIR /rest_api_project

COPY . .

RUN go mod init restapi && go mod tidy

RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal

RUN go build -o main .

EXPOSE 80:80

CMD ["./main"]
