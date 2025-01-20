# Структура проекта `messago`

```plaintext
messago/
├── cmd/                   // Точки входа (CLI, HTTP API)
│   └── app/               // Основное приложение
│       └── main.go
├── internal/              // Внутренняя логика проекта
│   ├── app/               // Сценарии использования (application layer)
│   ├── domain/            // Бизнес-домен (domain layer)
│   │   ├── models/        // Сущности (entities) и value objects
│   │   ├── repositories/  // Интерфейсы репозиториев
│   │   └── services/      // Доменные сервисы
│   ├── infrastructure/    // Инфраструктурный слой
│   │   ├── database/      // Логика доступа к БД PostgreSQL, Redis, etc.
│   │   ├── events/        // События, message brokers Kafka, RabbitMQ, etc.
│   │   ├── http/          // HTTP API (роуты, middleware)
│   │   └── ws/            // WebSocket серверы (если нужно)
│   └── config/            // Конфигурации (файлы, переменные окружения)
├── pkg/                   // Внешние пакеты, которые можно переиспользовать
├── docs/                  // Документация (swagger, md-файлы)
├── tests/                 // Интеграционные тесты
├── go.mod                 // Файл модуля Go
└── go.sum                 // Контрольная сумма зависимостей
