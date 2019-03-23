#!/usr/bin/env bash

version="v$1"


echo " _______             __                                         ";
echo "/       \           /  |                                        ";
echo "$$$$$$$  |  ______  $$ |  ______    ______    _______   ______  ";
echo "$$ |__$$ | /      \ $$ | /      \  /      \  /       | /      \ ";
echo "$$    $$< /$$$$$$  |$$ |/$$$$$$  | $$$$$$  |/$$$$$$$/ /$$$$$$  |";
echo "$$$$$$$  |$$    $$ |$$ |$$    $$ | /    $$ |$$      \ $$    $$ |";
echo "$$ |  $$ |$$$$$$$$/ $$ |$$$$$$$$/ /$$$$$$$ | $$$$$$  |$$$$$$$$/ ";
echo "$$ |  $$ |$$       |$$ |$$       |$$    $$ |/     $$/ $$       |";
echo "$$/   $$/  $$$$$$$/ $$/  $$$$$$$/  $$$$$$$/ $$$$$$$/   $$$$$$$/ ";
echo "                                                                ";
echo "                                                                ";
echo "                                                                ";


echo "Iniciando processo de releasing da versão $version"

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