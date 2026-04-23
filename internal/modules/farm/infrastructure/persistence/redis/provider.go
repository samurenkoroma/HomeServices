package redis

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/infrastructure/configs"
	"samurenkoroma/services/internal/modules/auth/domain"
	redis2 "samurenkoroma/services/pkg/redis"

	"github.com/redis/go-redis/v9"
)

type RedisFarmProvider struct {
	client *redis.Client

	users domain.Repository
}

func (p *RedisFarmProvider) ProviderName() string {
	return "auth"
}

func NewRedisFarmProvider(tx *sql.Tx) repository.RepositoryProvider {
	conf := configs.LoadConfig()
	client, err := redis2.NewClient(conf.Redis)
	if err != nil {
		panic(err)
	}
	return &RedisFarmProvider{
		client: client,
	}
}

func (p *RedisFarmProvider) Users() domain.Repository {
	if p.users == nil {
		p.users = NewRedisUserRepository(p.client)
	}
	return p.users
}
