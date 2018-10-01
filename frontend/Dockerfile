FROM node:latest
ENV NPM_CONFIG_LOGLEVEL warn
ARG app_env
ENV NODE_ENV $app_env
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 3000
RUN yarn install
CMD ["/app/run.sh"]
