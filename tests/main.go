package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var count int
func main() {
  go sendReq("http://localhost/apiV1/defi", 1000)
  time.Sleep(5*time.Second)
  go sendReq("http://localhost/apiV1/defi1", 1000)
  time.Sleep(5*time.Second)
  go sendReq("http://localhost/apiV1/defi2", 1000)
  time.Sleep(5*time.Second)
  go sendReq("http://localhost/apiV1/defi4", 1000)
  for {
    go sendReq("http://localhost/apiV1/defi3s", 1)
    time.Sleep(time.Second/ 1000)
    count++
  }
  time.Sleep(5*time.Minute)

}

func sendReq(apiURL string, numRequests int) {
  // Создаем HTTP клиент
	client := &http.Client{
		Timeout: time.Second * 10, // Установите таймаут для запроса
	}

	// Переменные для хранения общего времени и количества успешных запросов
	totalTime := time.Duration(0)
	successfulRequests := 0

	for i := 1; i <= numRequests; i++ {
		// Формируем запрос к API
		request, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			log.Fatalf("Ошибка при создании запроса: %v", err)
		}

		// Измеряем время выполнения запроса
		startTime := time.Now()
		response, err := client.Do(request)
		elapsedTime := time.Since(startTime)

		if err != nil {
			log.Printf("Ошибка при выполнении запроса: %v", err)
		} else {
			// Увеличиваем счетчик успешных запросов
			successfulRequests++

			// Читаем тело ответа
			_, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("Ошибка при чтении тела ответа: %v", err)
			}

			// Закрываем тело ответа
			err = response.Body.Close()
			if err != nil {
				log.Printf("Ошибка при закрытии тела ответа: %v", err)
			}

			// Выводим информацию о запросе
			fmt.Printf("Эндпоинт - %s, Запрос #%d: Статус - %s, Время выполнения - %s\n", apiURL, count, response.Status, elapsedTime)
			totalTime += elapsedTime
		}
	}

	// Выводим статистику
	// fmt.Printf("Всего запросов: %d\n", numRequests)
	// fmt.Printf("Успешных запросов: %d\n", successfulRequests)
	// fmt.Printf("Среднее время выполнения: %s\n", totalTime/time.Duration(successfulRequests))
}
