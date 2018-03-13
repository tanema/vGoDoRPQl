# vGoDoRPQl
Example stack of:
- [vGo](https://github.com/golang/vgo)
  - [Gorm](https://github.com/jinzhu/gorm)
  - [gqlgen](github.com/vektah/gqlgen)
- Docker
- React
- PostgresQL
- GraphQL

Fork of this project: https://github.com/McMenemy/GoDoRP

Disclaimer: This project is not actively supported and not recommended for production apps. Hope it serves as a learning resource.

## Benefits
* Start a vGoDoRPQl project with one command on any computer with docker-compose installed
* Dev mode features hot reloading on code changes for both the GoLang backend and React frontend (no need to rebuild containers while coding)
* Anyone can contribute to your project locally without having to setup/install GOPATH, Postgres, node etc
* Easily modifiable graphQL schema. (See makefile `gen-schema` rule for command)
* Dev environment is the same as production environment

## Getting started:
* download [docker-compose](https://docs.docker.com/compose/install/) if not already installed
Then run the following commands:

```bash
$ git clone https://github.com/tanema/vGoDoRPQl.git
$ cd vGoDoRPQl
$ make
$ #open localhost:3000 to see it in action
```
