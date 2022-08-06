package database

type Config struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     int    `env:"DATABASE_PORT,default=5432"`
	Database string `env:"DATABASE_NAME,default=go-product-showcase"`
	User     string `env:"DATABASE_USER,default=postgres"`
	Password string `env:"DATABASE_PASSWORD,default=secret"`
}
