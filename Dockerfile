FROM ubuntu:latest
LABEL authors="yawai"

ENTRYPOINT ["top", "-b"]