FROM ubuntu:latest
LABEL authors="shoma"

ENTRYPOINT ["top", "-b"]