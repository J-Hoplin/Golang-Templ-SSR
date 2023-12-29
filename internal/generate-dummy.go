package internal

import (
	"cicd/internal/database"
	"cicd/internal/database/model"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-faker/faker/v4"
)

func GenerateDummy() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()

			todoCollection := database.GetCollection(database.TODO)
			var task model.TODOModel
			if err := faker.FakeData(&task); err != nil {
				log.Printf("Number.%d - Fail to generate data: %v", no, err.Error())
				return
			}
			task.IsComplete = false
			if _, err := todoCollection.InsertOne(context.TODO(), task); err != nil {
				log.Printf("Number.%d - Fail to insert data: %v", no, err.Error())
			}
			fmt.Printf("Number.%d completed\n", no)
		}(i + 1)
	}

	wg.Wait()
}
