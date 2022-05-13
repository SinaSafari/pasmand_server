package config

import (
	"pasmand/internal/config/database"
	"pasmand/internal/config/redis"
)

func SetupDependencies() {
	database.SetupDatabase()
	redis.SetupRedis()
}
