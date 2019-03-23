#!/usr/bin/env bash

cd ../

version="v$1"

echo "####################################################################"
echo "#                                                                  #"
echo "#        Iniciando processo de releasing da versão $version        #"
echo "#                                                                  #"
echo "####################################################################"

echo "Tudo certo? (S/n)"
read resp
if [[ resp = "n" ]]; then
    exit
fi

echo "Fechando commits"
git add .
git commit -am "Fechando versão: $version"

echo "Gerando nova tag"
git tag -a ${version} -m "versao ${version}"
git push