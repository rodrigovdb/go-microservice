FROM golang

ENV APP_PATH /go/src/app
ENV BIN_PATH /go/bin/

ADD . $APP_PATH/
COPY .env $BIN_PATH/

RUN go get github.com/joho/godotenv

RUN go install -i app

ENTRYPOINT /go/bin/app

EXPOSE 8080
