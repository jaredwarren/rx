// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// recipe HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/jaredwarren/rx/design

package server

import (
	"context"
	"net/http"

	recipeviews "github.com/jaredwarren/rx/gen/recipe/views"
	goahttp "goa.design/goa/http"
)

// EncodeListResponse returns an encoder for responses returned by the recipe
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(recipeviews.RxRecipeCollection)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalIngredientListViewToIngredientListResponseBody builds a value of type
// *IngredientListResponseBody from a value of type
// *recipeviews.IngredientListView.
func marshalIngredientListViewToIngredientListResponseBody(v *recipeviews.IngredientListView) *IngredientListResponseBody {
	if v == nil {
		return nil
	}
	res := &IngredientListResponseBody{}
	if v.Ingredients != nil {
		res.Ingredients = make([]*IngredientResponseBody, len(v.Ingredients))
		for j, val := range v.Ingredients {
			res.Ingredients[j] = &IngredientResponseBody{
				Quantity: val.Quantity,
			}
			if val.Recipe != nil {
				res.Ingredients[j].Recipe = marshalRxRecipeViewToRxRecipeResponseBody(val.Recipe)
			}
		}
	}

	return res
}

// marshalRxRecipeViewToRxRecipeResponseBody builds a value of type
// *RxRecipeResponseBody from a value of type *recipeviews.RxRecipeView.
func marshalRxRecipeViewToRxRecipeResponseBody(v *recipeviews.RxRecipeView) *RxRecipeResponseBody {
	if v == nil {
		return nil
	}
	res := &RxRecipeResponseBody{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Version:     v.Version,
		Favorite:    v.Favorite,
		Rating:      v.Rating,
		Difficulty:  v.Difficulty,
		State:       v.State,
		Complete:    v.Complete,
	}
	if v.Ingredients != nil {
		res.Ingredients = marshalIngredientListViewToIngredientListResponseBody(v.Ingredients)
	}

	return res
}
