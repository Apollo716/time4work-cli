package config

import "os"

func DBPass() string {
	return os.Getenv("DB_PASS")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}
