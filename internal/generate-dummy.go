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
				log.Fatalln("Fail to generate data")
				return
			}
			task.IsComplete = false
			if _, err := todoCollection.InsertOne(context.TODO(), task); err != nil {
				log.Fatalln("Fail to insert data")
			}
			fmt.Printf("Number.%d completed\n", no)
		}(i + 1)
	}

	wg.Wait()
}
