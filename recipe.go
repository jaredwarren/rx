package recipe

import (
	"context"
	"log"

	recipesvc "github.com/jaredwarren/rx/gen/recipe"
)

// recipe service example implementation.
// The example methods log the requests and return zero values.
type recipeSvc struct {
	logger *log.Logger
}

// NewRecipe returns the recipe service implementation.
func NewRecipe(logger *log.Logger) recipesvc.Service {
	return &recipeSvc{logger}
}

// List recipes
func (s *recipeSvc) List(ctx context.Context) (res recipesvc.RecipeRecipeCollection, err error) {
	s.logger.Print("recipe.list")
	r := &recipesvc.RecipeRecipe{
		ID:    "123",
		Title: "recipe",
	}
	res = recipesvc.RecipeRecipeCollection{r}

	return
}
