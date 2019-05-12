#!/usr/bin/env bash

docker stop mysql && docker rm mysql && rm -fr mysql && scripts/docker_new_mysql.sh