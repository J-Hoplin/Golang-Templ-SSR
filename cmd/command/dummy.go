package command

import (
	"cicd/internal/database"
	"cicd/internal/database/model"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var GenerateDummy = &cobra.Command{
	Use: "dummy",
	Short: "Generate dummy Data",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load Dot env
		if err := godotenv.Load(); err != nil {
			log.Fatalln("Fail to load .env file:", err.Error())
		}
		database.Connect()
		defer database.Disconnect()
		
		var wg sync.WaitGroup

		for i := 0; i < 100;i++{
			wg.Add(1)
			go func(no int){
				defer wg.Done()
				
				todoCollection := database.GetCollection(database.TODO)
				var task model.TODOModel
				if err := faker.FakeData(&task); err != nil{
					log.Fatalln("Fail to generate data")
					return
				}
				task.IsComplete = false
				if _,err := todoCollection.InsertOne(context.TODO(),task); err != nil{
					log.Fatalln("Fail to insert data")
				}
				fmt.Printf("Number.%d completed\n",no)
			}(i + 1)
		}

		wg.Wait()
		return nil
	},
}

func init(){}