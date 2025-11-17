#!/bin/bash
set -e

# Intentar cargar variables desde .env en la raíz (si existe)
if [ -f ../.env ]; then
  export $(grep -v '^#' ../.env | xargs)
fi

TIMESTAMP=$(date +"%Y%m%d-%H%M")

# Construir URI si no está explícita
MONGO_URI=${MONGO_URI:-"mongodb://${MONGO_USER}:${MONGO_PASSWORD}@localhost:27017/${MONGO_DB}?authSource=admin"}

mkdir -p ./backups
mongodump --uri="$MONGO_URI" --out=./backups/backup-$TIMESTAMP
echo "Backup generado en ./backups/backup-$TIMESTAMP"
