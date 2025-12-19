#! /bin/bash


export TS_EXPORTER_LATEST=$(git ls-remote git@github.com:moodyRahman/ts-exporter.git HEAD | cut -f1 | cut -c1-7)

echo $TS_EXPORTER_LATEST

docker compose down
docker compose up -d