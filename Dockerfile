FROM golang:1.12 as builder

RUN git clone https://github.com/tetymd/lottery.git /lottery && \
    cd /lottery && \
    go build -o lottery-server main.go

FROM alpine:latest

WORKDIR /lottery
COPY --from=builder /lottery .

CMD ./lottery-server
