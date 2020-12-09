# Build the project
FROM golang:1.14 

WORKDIR /usr/app
ADD . .

RUN apt-get -y update && apt-get -y install imagemagick

ENV PATH="$PATH:/usr/app/scripts"

RUN make build
RUN make test
RUN make

EXPOSE 8080

CMD ["./main"]
