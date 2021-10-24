package core

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/monitprod/core/pkg/loaders/database"
	"github.com/monitprod/core/pkg/util"
)

// UseCoreSmp is simple form of UseCore method
func UseCoreSmp(execution func(ctx context.Context)) {
	ctx := context.Background()

	UseCore(ctx, func() error {
		execution(ctx)
		return nil
	})
}

func UseCore(ctx context.Context, execution func() error) error {
	start(ctx)

	err := execution()

	defer close(ctx)

	if err != nil {
		log.Fatalln("Error at execution on UseCore\n", err)

		return err
	}

	return nil
}

func start(ctx context.Context) {
	err := godotenv.Load(util.GetRootPath() + "/.env")

	if err != nil {
		log.Println("INFO: Core dot env not initialized:", err)
	}

	database.ConnectClient(ctx)

	log.Println("Core Started!")
}

func close(ctx context.Context) {
	err := database.DisconnectClient(ctx)

	if err != nil {
		log.Printf("Is not possible disconnect client,\n" +
			"if you running locally, don't worry :)")
	}

}
