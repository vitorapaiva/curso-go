para builda
docker container run --rm -i -v ${PWD}:/usr/src/app -w /usr/src/app golang:1.18 go build -v

para executar
docker container run --rm -i -v ${PWD}:/usr/src/app -w /usr/src/app golang:1.18 go run main

para fazer as duas coisas
docker container run --rm -i -v ${PWD}:/usr/src/app -w /usr/src/app golang:1.18 make buildrun