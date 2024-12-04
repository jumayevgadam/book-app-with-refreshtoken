FROM ubuntu:latest
LABEL authors="gadam"

ENTRYPOINT ["top", "-b"]