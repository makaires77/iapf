#!/bin/bash

# Define o caminho base
BASE_PATH="./backend/GraphScraperPasteur"

# Cria a estrutura de diretórios para o módulo
mkdir -p "$BASE_PATH/cmd"
mkdir -p "$BASE_PATH/pkg"
mkdir -p "$BASE_PATH/internal/graphdb"
mkdir -p "$BASE_PATH/internal/models"
mkdir -p "$BASE_PATH/test"
mkdir -p "./docker"
mkdir -p "./kubernetes"

# Cria arquivos dentro da estrutura do módulo
touch "$BASE_PATH/cmd/main.go"
touch "$BASE_PATH/pkg/scraper.go"
touch "$BASE_PATH/internal/graphdb/neo4j.go"
touch "$BASE_PATH/internal/models/researcher.go"
touch "$BASE_PATH/internal/models/team.go"
touch "$BASE_PATH/internal/models/department.go"
touch "$BASE_PATH/internal/models/plataform_facility.go"
touch "$BASE_PATH/internal/models/transversal_project.go"
touch "$BASE_PATH/internal/models/reference_center.go"
touch "$BASE_PATH/test/scraper_test.go"
touch "$BASE_PATH/test/neo4j_test.go"
touch "./docker/Dockerfile"
touch "./kubernetes/deployment.yaml"

# chmod +x 01_scrap_pasteur.sh
# ./01_scrap_pasteur.sh
