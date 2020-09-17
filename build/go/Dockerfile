FROM golang:1.14-alpine

WORKDIR /app/ip2loc
COPY . .

RUN go build -o ./bin/ip2loc .

ARG USER=application
ARG GROUP=application
RUN addgroup -g 1000 -S ${GROUP} && \
        adduser -u 1000 -S ${USER} -G ${GROUP}

RUN chown -hR application:application /app/ip2loc

USER $USER

CMD ["./bin/ip2loc"]