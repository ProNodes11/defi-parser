package main

import (
	"context"
	"encoding/json"
	dfscan "github.com/ProNodes11/defi-parser/defiscan"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Log message")
	log.Warn("Log message")
	log.Error("Log message")
	log.WithField("value", 42).Info("Log message with value")
	dfscan.db.GetValue('mydata')
	// log.WithFields("event", "event", "topic", "topic", "key", "key").Fatal("Failed to send event")
	// upload()
	// api()

}

func api() {
	// Создание экземпляра маршрутизатора Gin
	router := gin.Default()

	// Главная страница
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Добро пожаловать на главную страницу!",
		})
	})

	// Маршрут с динамическим параметром
	router.GET("/users/:id", func(c *gin.Context) {
		// Получение значения параметра из URL
		userID := c.Param("id")

		// Возвращаем ответ с информацией о пользователе
		c.JSON(http.StatusOK, gin.H{
			"message": "Вы запросили информацию о пользователе с ID " + userID,
		})
	})

	// Вложенный маршрут
	api := router.Group("/apiV1")
	{
		api.GET("/defi", func(c *gin.Context) {
			jsonData := download()
			c.JSON(200, jsonData)
		})
		api.GET("/defi1", func(c *gin.Context) {
			jsonData := download()
			c.JSON(200, jsonData)
		})
		api.GET("/defi2", func(c *gin.Context) {
			jsonData := download()
			c.JSON(200, jsonData)
		})
		api.GET("/defi4", func(c *gin.Context) {
			jsonData := download()
			c.JSON(200, jsonData)
		})
	}

	// Запуск сервера на порту 8080
	router.Run(":8080")
}

func download() interface{} {
	ctx := context.Background()

	// Создание клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Адрес и порт Redis сервера
		Password: "password",   // Пароль (если требуется)
		DB:       0,            // Индекс базы данных
	})

	// Получение данных из Redis
	result, err := client.Get(ctx, "mydata").Result()
	if err != nil {
		log.Fatal(err)
	}

	// Декодирование данных из формата Redis в JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}
