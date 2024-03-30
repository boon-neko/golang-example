FROM ubuntu:latest
LABEL authors="amoru"

ENTRYPOINT ["top", "-b"]