package db

type Config struct {
	Driver      string
	Dsn         string
	MaxOpen     int
	MaxIdle     int
	MaxLifetime int
	QueryLevel  string
}
