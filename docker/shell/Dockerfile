FROM golang:1.19 as dev
ENV LINE_APP_ROOT=/go/src/app
ENV app_env=test
RUN apt update
RUN apt install -y default-mysql-client-core/stable
WORKDIR /go/src/app
CMD ["bash"]
