FROM golang:1.14.3-alpine
RUN apk add git
RUN mkdir /code
RUN mkdir /code/vendor
RUN go get github.com/korylprince/go-ad-auth
RUN go get github.com/tkanos/gonfig
WORKDIR /code
COPY . /code/
RUN go build  -o serv
CMD ./serv
#if [[ `sudo netstat -lunp | grep 9000  |wc -l` == "1" ]]; then echo "True";fi
