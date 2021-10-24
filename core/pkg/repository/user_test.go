package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/vo"
)

var getUsersOptionsMock = GetUsersOptions{
	Page: vo.PaginateOptions{
		CurrentPage: 1,
		PageSize:    1,
	},
}

func TestUsersRepository(t *testing.T) {

	core.UseCoreSmp(func(ctx context.Context) {
		userRepository := NewUserRepositoryMongoDB()

		users, err := userRepository.GetUsers(ctx,
			getUsersOptionsMock)

		if err != nil {
			log.Fatalln("Error while get users from repository", err)
		}

		log.Println(*users)
	})
}
