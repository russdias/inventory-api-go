package main

import (
	"log"
	"math/rand"

	"github.com/jaswdr/faker/v2"
	"github.com/russelldias98/go-api/initializers"
	"github.com/russelldias98/go-api/models"
)

func main() {
	log.Println("Seeding database")
	initializers.SetupEnv()
	initializers.ConnectToDB()

	fake := faker.New()

	// Create 10 categories
	for i := 0; i < 10; i++ {
		initializers.DB.Create(&models.Category{
			Name:        fake.Person().Name(),
			Description: fake.Lorem().Sentence(10),
		})
	}

	// Create 10 suppliers
	for i := 0; i < 10; i++ {
		initializers.DB.Create(&models.Supplier{
			Name:        fake.Person().Name(),
			ContactInfo: fake.Person().Contact().Phone,
		})
	}

	// Create 50 products
	for i := 0; i < 50; i++ {
		initializers.DB.Create(&models.Product{
			Name:        fake.Company().Name(),
			Description: fake.Lorem().Sentence(10),
			Quantity:    rand.Intn(100),
			Price:       rand.Float64() * 100,
			SupplierID:  uint(rand.Intn(10)),
			CategoryID:  uint(rand.Intn(10)),
		})
	}
}
