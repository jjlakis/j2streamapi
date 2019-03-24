FROM golang
ADD . /go/src/github.com/jjlakis/j2streamapi
RUN go get github.com/go-chi/chi
RUN go build -o /go/bin/j2streamapi /go/src/github.com/jjlakis/j2streamapi/main.go 

CMD /go/bin/j2streamapi
EXPOSE 8080
