package defiscan

import (
  	"github.com/sirupsen/logrus"
	"context"
	"github.com/go-redis/redis/v8"
)

var log = logrus.New()
var ctx = context.Background()

type DbClient struct {
	client *redis.Client
}

func dbClient() (*DbClient, error) {
	// Инициализация клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес и порт Redis сервера
		Password: "",               // Пароль Redis сервера (если требуется)
		DB:       0,                // Номер базы данных Redis
	})

	// Проверка соединения с Redis
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Error('Connect to db failed')
	}

	return &DbClient{
		client: client,
	}, nil
}

func (db *DbClient) GetValue(key string) (string, error) {
	value, err := db.client.Get(key).Result()
	if err == redis.Nil {
		log.Error('Data not found in db')
		return "", nil
	} else if err != nil {
		log.Error('Fetch info from db failed')
		// Ошибка при получении значения
		return "", err
	}
	return value, nil
}

func (db *DbClient) SetValue(key string, value string) error {
	err := db.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		// Ошибка при записи значения
		return err
	}
	return nil
}
