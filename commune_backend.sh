#!/bin/bash

case $1 in
  start)
    docker-compose --env-file .env -p commune_backend up;;
  stop)
    docker-compose -p commune_backend down ;;
  rmi)
    docker-compose -p commune_backend down --rmi all ;;
  *)
    echo "Using: ./commune_backend.sh [start, stop, rmi]" ;;
esac;