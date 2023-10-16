#!/bin/bash

export UID=${UID}
export GID=${GID}

rm -r /tmp/postgres/
mkdir -p /tmp/postgres
docker-compose -f docker-compose-postgress-dev.yml down
docker-compose -f docker-compose-postgress-dev.yml up
