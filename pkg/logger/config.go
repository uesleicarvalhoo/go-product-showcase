package logger

type Config struct {
	LogLevel    string `json:"log_level"`
	ConsoleJSON bool   `json:"console_json"`
}
