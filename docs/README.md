# Структура проекта `messago`

```plaintext
messago/
├── cmd/                   // Точки входа (CLI, HTTP API)
│   └── app/               // Основное приложение
│       └── main.go
├── internal/              // Внутренняя логика проекта
│   ├── api/               // http handlers + interfaces
│   ├── bootstrap/         // fx bs приложения
│   ├── config/            // app's configuration files
│   ├── entity/            // app's entities 
│   ├── pkg/               // packages wich shoud be reused
│   ├── valueobject/       // valueobjects are here
│   ├── usecases/
├── migrations/            // postgres' migrations
├── docs/                  // Документация (swagger, md-файлы)
├── go.mod                 // Файл модуля Go
└── go.sum                 // Контрольная сумма зависимостей
