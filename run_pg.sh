#! /bin/bash

docker run --name pg --restart=always -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres