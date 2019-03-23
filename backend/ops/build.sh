#!/usr/bin/env bash

version="v$1"
echo "##########################################################################"
echo "#                                                                        #"
echo "# Buildando para repo caiobarcelos/monografia com tag backend-${version} #"
echo "#                                                                        #"
echo "##########################################################################"

IMAGEM="caiobarcelos/monografia:backend-$1"

docker build -t ${IMAGEM} .

echo "##########################################################################"
echo "#                                                                        #"
echo "#                        Imagem gerada ${IMAGEM}                         #"
echo "#                                                                        #"
echo "##########################################################################"

sh ops/push-to-dockerhub.sh ${IMAGEM}