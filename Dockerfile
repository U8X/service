FROM docker.youplus.cc/go-grpc:1.8.3

# Deps
RUN go get -v github.com/go-redis/redis
RUN go get -v gopkg.in/mgo.v2
RUN go get -v github.com/Sirupsen/logrus

# Build app
ENV APP_SRC $GOPATH/src/github.com/U8X/service
COPY . $APP_SRC
RUN cd $APP_SRC && go build -v -o $GOPATH/bin/short_url_d
RUN rm -rf $APP_SRC

ENTRYPOINT ["short_url_d"]
