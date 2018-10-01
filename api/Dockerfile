FROM golang:latest
ARG app_env
ENV APP_ENV $app_env
RUN mkdir -p /app
ADD . /app/
WORKDIR /app
EXPOSE 8080
RUN go build -o main .
CMD ["/app/run.sh"]
