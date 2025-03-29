#!/bin/bash

container_tool=podman

docker_image_name=eosorio/personalf-services-currencies
version_today=$(date "+%Y%m%d")

# Produced image should be dev and after testing promoted as stable (removing the dev tag)
$container_tool build -t $docker_image_name:${version_today}-dev -f Containerfile .

# Tagging image as latest
$container_tool tag $docker_image_name:${version_today}-dev $docker_image_name:latest
