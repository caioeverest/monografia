#!/usr/bin/env bash


echo " _______                       __                __                            __                      __                            __                  __       ";
echo "/       \                     /  |              /  |                          /  |                    /  |                          /  |                /  |      ";
echo "$$$$$$$  | __    __   _______ $$ |____         _$$ |_     ______          ____$$ |  ______    _______ $$ |   __   ______    ______  $$ |____   __    __ $$ |____  ";
echo "$$ |__$$ |/  |  /  | /       |$$      \       / $$   |   /      \        /    $$ | /      \  /       |$$ |  /  | /      \  /      \ $$      \ /  |  /  |$$      \ ";
echo "$$    $$/ $$ |  $$ |/$$$$$$$/ $$$$$$$  |      $$$$$$/   /$$$$$$  |      /$$$$$$$ |/$$$$$$  |/$$$$$$$/ $$ |_/$$/ /$$$$$$  |/$$$$$$  |$$$$$$$  |$$ |  $$ |$$$$$$$  |";
echo "$$$$$$$/  $$ |  $$ |$$      \ $$ |  $$ |        $$ | __ $$ |  $$ |      $$ |  $$ |$$ |  $$ |$$ |      $$   $$<  $$    $$ |$$ |  $$/ $$ |  $$ |$$ |  $$ |$$ |  $$ |";
echo "$$ |      $$ \__$$ | $$$$$$  |$$ |  $$ |        $$ |/  |$$ \__$$ |      $$ \__$$ |$$ \__$$ |$$ \_____ $$$$$$  \ $$$$$$$$/ $$ |      $$ |  $$ |$$ \__$$ |$$ |__$$ |";
echo "$$ |      $$    $$/ /     $$/ $$ |  $$ |        $$  $$/ $$    $$/       $$    $$ |$$    $$/ $$       |$$ | $$  |$$       |$$ |      $$ |  $$ |$$    $$/ $$    $$/ ";
echo "$$/        $$$$$$/  $$$$$$$/  $$/   $$/          $$$$/   $$$$$$/         $$$$$$$/  $$$$$$/   $$$$$$$/ $$/   $$/  $$$$$$$/ $$/       $$/   $$/  $$$$$$/  $$$$$$$/  ";
echo "                                                                                                                                                                  ";
echo "                                                                                                                                                                  ";
echo "                                                                                                                                                                  ";

echo "Fazendo push da imagem"

docker push ${1}