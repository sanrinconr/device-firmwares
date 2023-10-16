#!/bin/bash
docker-compose down
docker-compose up
docker logs device-firmware-container -f
