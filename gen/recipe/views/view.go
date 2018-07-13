// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// recipe views
//
// Command:
// $ goa gen github.com/jaredwarren/rx/design

package views

import (
	goa "goa.design/goa"
)

// RxRecipeCollection is the viewed result type that is projected based on a
// view.
type RxRecipeCollection struct {
	// Type to project
	Projected RxRecipeCollectionView
	// View to render
	View string
}

// RxRecipeCollectionView is a type that runs validations on a projected type.
type RxRecipeCollectionView []*RxRecipeView

// RxRecipeView is a type that runs validations on a projected type.
type RxRecipeView struct {
	// Unique recipe ID
	ID *string
	// Recipe Title
	Title *string
	// Long description of recipe
	Description *string
	// Version Number e.g. 1.0.1
	Version *string
	// List of ingredients
	Ingredients *IngredientListView
	// Is a favorite, basically a tag
	Favorite *bool
	// rating between 0-1
	Rating *float32
	// rating between 0-1
	Difficulty *float32
	// e.g. chopped, sliced, etc.. might need to be array.
	State *string
	// If it's been added/included
	Complete *bool
}

// IngredientListView is a type that runs validations on a projected type.
type IngredientListView struct {
	// List of ingredients
	Ingredients []*IngredientView
}

// IngredientView is a type that runs validations on a projected type.
type IngredientView struct {
	Recipe *RxRecipeView
	// TODO: make UnitOfMeasure
	Quantity *string
}

// Validate runs the validations defined on the viewed result type
// RxRecipeCollection.
func (result RxRecipeCollection) Validate() (err error) {
	switch result.View {
	case "default", "":
		err = result.Projected.Validate()
	case "tiny":
		err = result.Projected.ValidateTiny()
	case "ingredient":
		err = result.Projected.ValidateIngredient()
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny", "ingredient"})
	}
	return
}

// Validate runs the validations defined on RxRecipeCollectionView using the
// "default" view.
func (result RxRecipeCollectionView) Validate() (err error) {
	for _, item := range result {
		if err2 := item.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTiny runs the validations defined on RxRecipeCollectionView using
// the "tiny" view.
func (result RxRecipeCollectionView) ValidateTiny() (err error) {
	for _, item := range result {
		if err2 := item.ValidateTiny(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateIngredient runs the validations defined on RxRecipeCollectionView
// using the "ingredient" view.
func (result RxRecipeCollectionView) ValidateIngredient() (err error) {
	for _, item := range result {
		if err2 := item.ValidateIngredient(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Validate runs the validations defined on RxRecipeView using the "default"
// view.
func (result *RxRecipeView) Validate() (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	return
}

// ValidateTiny runs the validations defined on RxRecipeView using the "tiny"
// view.
func (result *RxRecipeView) ValidateTiny() (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	return
}

// ValidateIngredient runs the validations defined on RxRecipeView using the
// "ingredient" view.
func (result *RxRecipeView) ValidateIngredient() (err error) {
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	return
}

// Validate runs the validations defined on IngredientListView.
func (result *IngredientListView) Validate() (err error) {
	for _, e := range result.Ingredients {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Validate runs the validations defined on IngredientView.
func (result *IngredientView) Validate() (err error) {
	if result.Recipe != nil {
		if err2 := result.Recipe.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}