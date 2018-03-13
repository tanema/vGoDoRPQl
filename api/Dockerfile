FROM golang

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/tanema/vGoDoRPQl/api
WORKDIR /go/src/github.com/tanema/vGoDoRPQl/api
RUN echo "machine api.github.com login tanema password [github-oauth-token]" > ~/.netrc
RUN go get -u github.com/golang/vgo
RUN CGO_ENABLED=0 GOOS=linux vgo build -a -installsuffix nocgo -o /app .

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
  then \
  api; \
  else \
  go get github.com/pilu/fresh && \
  fresh; \
  fi

EXPOSE 8080
