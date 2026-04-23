package redis

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/infrastructure/configs"
	"samurenkoroma/services/internal/modules/auth/domain"
	redis2 "samurenkoroma/services/pkg/redis"

	"github.com/redis/go-redis/v9"
)

type RedisAuthProvider struct {
	client *redis.Client

	users domain.Repository
}

func (p *RedisAuthProvider) ProviderName() string {
	return "auth"
}

func NewRedisAuthProvider(tx *sql.Tx) repository.RepositoryProvider {
	conf := configs.LoadConfig()
	client, err := redis2.NewClient(conf.Redis)
	if err != nil {
		panic(err)
	}
	return &RedisAuthProvider{
		client: client,
	}
}

func (p *RedisAuthProvider) Users() domain.Repository {
	if p.users == nil {
		p.users = NewRedisUserRepository(p.client)
	}
	return p.users
}
