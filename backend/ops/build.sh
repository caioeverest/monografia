#!/usr/bin/env bash

echo " _______             __  __        __ "
echo "/       \           /  |/  |      /  |"
echo "$$$$$$$  | __    __ $$/ $$ |  ____$$ |"
echo "$$ |__$$ |/  |  /  |/  |$$ | /    $$ |"
echo "$$    $$< $$ |  $$ |$$ |$$ |/$$$$$$$ |"
echo "$$$$$$$  |$$ |  $$ |$$ |$$ |$$ |  $$ |"
echo "$$ |__$$ |$$ \__$$ |$$ |$$ |$$ \__$$ |"
echo "$$    $$/ $$    $$/ $$ |$$ |$$    $$ |"
echo "$$$$$$$/   $$$$$$/  $$/ $$/  $$$$$$$/ "
echo "                                      "
echo "                                      "
echo "                                      "

echo "Buildando para repo caiobarcelos/monografia with tag backend-$0"

IMAGEM="caiobarcelos/monografia:backend-$0"

docker build -t ${IMAGEM} .

echo "Imagem gerada ${IMAGEM}"
echo ""
echo "Detalhes:"
docker images ${IMAGEM}

sh push-to-dockerhub.sh ${IMAGEM}