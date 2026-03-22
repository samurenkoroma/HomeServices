# Home Services — Умная ферма / Тепличное хозяйство

Бэкенд для управления полями, культурами, планами посадок, циклами выращивания и задачами ухода.

## Стек
- Go 1.25
- Fiber + JWT
- Postgres + GORM (миграции)
- DDD + CQRS + Domain Events

## Запуск
```bash
cp .env.example .env
docker compose up -d db
go run ./cmd/server