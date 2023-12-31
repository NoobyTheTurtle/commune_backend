#!/bin/bash

case $1 in
  start)
    docker-compose --env-file .env -p commune_backend up;;
  start_db)
    docker-compose --env-file .env -p commune_backend up postgres;;
  stop)
    docker-compose -p commune_backend down ;;
  rmi)
    docker-compose -p commune_backend down --rmi all ;;
  swagger)
    swag init --parseInternal -g cmd/app/main.go ;;
  *)
    echo "Using: ./commune_backend.sh [start, start_db, stop, rmi, swagger]" ;;
esac;