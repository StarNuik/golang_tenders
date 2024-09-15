#! /bin/bash

docker-compose \
    --env-file ./test/test.env \
    -f ./compose.test.yaml \
    $@