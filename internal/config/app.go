package config

// App - app config.
// All related environment variables must be prefixed with APP_.
// Except for some generic environment variables like TZ, etc.
type App struct {
	TZ string `envconfig:"TZ" default:"Europe/Moscow"`

	Name           string `envconfig:"APP_NAME" default:"messago"`
	Environment    string `envconfig:"APP_ENV" default:"prod"`
	Host           string `envconfig:"APP_HOST" default:"localhost"`
	Port           int    `envconfig:"APP_PORT" default:"8081"`
	LogLevel       string `envconfig:"APP_LOG_LEVEL" default:"info"`
	PrettyLogs     bool   `envconfig:"APP_LOG_PRETTY" default:"false"`
	SwaggerPath    string `envconfig:"APP_SWAGGER_PATH" default:"swagger"`
	AuthSecret     string `envconfig:"APP_AUTH_SECRET" default:"secret"`
	RateLimitRPS   int    `envconfig:"APP_RATE_LIMIT_RPS" default:"10"`
	RateLimitBurst int    `envconfig:"APP_RATE_LIMIT_BURST" default:"200"`
}
