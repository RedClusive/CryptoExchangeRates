#!/bin/bash

set -e

host="db"
port="5432"
cmd="$@"

>&2 echo "!!!!!!!! Check db for available !!!!!!!!"

until curl http://"$host":"$port"; do
  >&2 echo "DB is unavailable - sleeping"
  sleep 10
done

>&2 echo "DB is up - executing command"

sleep 120

exec $cmd