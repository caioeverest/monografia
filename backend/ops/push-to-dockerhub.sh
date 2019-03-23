#!/usr/bin/env bash

echo "####################################################################"
echo "Fazendo push da imagem"
docker images ${IMAGEM}
echo "####################################################################"

docker push ${1}