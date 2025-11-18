# microservicios_parcial3

## Requisitos
- Docker y Docker Compose
- Go 1.22+
- MongoDB (se levanta con Compose)

## Pasos de ejecución

1. Copia `.env.sample` a `.env` y edita tus credenciales.
2. Construye y levanta los servicios:
   ```
   copy .env.sample .env
   docker compose build --no-cache
   docker compose up -d --build
   docker compose logs -f
   docker compose ps
   ```
3. Ejecuta pruebas unitarias e integración:
   ```
   go test ./services/... -v -cover
   ```
4. Realiza backup de la base de datos:
   ```
   ./backup/backup.sh
   ```
   (Usa Git Bash o WSL en Windows)

## Arquitectura

![Diagrama](diagrams/Screenshot%202025-11-13%20180227.png)

- 4 microservicios (CRUD) en Go
- MongoDB como base de datos
- Orquestación con Docker Compose
- Pruebas y CI/CD automatizadas

## Validación

- Colección Postman en `postman/`
- Reporte de cobertura generado por CI/CD
- Evidencias: logs en `docker compose logs -f`
