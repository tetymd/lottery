FROM golang:1.12

RUN git clone https://github.com/tetymd/lottery.git /lottery && \
    cd /lottery && \
    go build -o lottery-serve main.go

CMD ./lottery-serve
