package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = Service("recipe", func() {
	Description("The recipe service retrieves recipes.")
	HTTP(func() {
		Path("/recipe")
	})
	Method("list", func() {
		Description("List recipes")
		Result(CollectionOf(Recipe), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
})


recipe 
 - result of steps

step
  - ingredients
  - action(s)

ingredient
 - quantity of recipe


// maybe I can organize things like git tree?


// quantity may need to scale sub recipe/step quantities. 

// RecipePayload ...
var RecipePayload = Type("Recipe", func() {
	Description("Recipe ...")
	// minimum data required to create or update a recipe
	Attribute("id", UInt, "ID")                // Attribute value of type integer
	Attribute("title", String, "Recipe Title") // Attribute value of type integer
	Attribute("description", String, "Long description of recipe")
	Attribute("image", String, "Image of recipe")

	Attribute("started", Boolean)
	Attribute("duration", "Time")
	Attribute("progress", "Time", "percent complete")
	Attribute("finished", Boolean)

	// dependance and parent are dependant on parent recipe???? how to make recipe reusable for other recipes
	Attribute("dependance", String, "")
	Attribute("parent", UInt, "parent ID")

	Attribute("ingredients", "IngredientList", "List of ingredients")

	// TODO: might need specific version for each component
	Attribute("version", String, "Version Number e.g. 1.0.1")

	// TODO: Fix Times
	// "times" maybe should be calculated from steps, if possible.
	// see if I can calculate time left, based on steps left?
	// Attribute("prep_time", DateTime, "Amount of time to prepare")
	// Attribute("cook_time", DateTime, "Amount of time to cook")
	// Attribute("wait_time", DateTime, "Amount of time to wait for things such as mairnading")

	// equipment is a compilation of all sub steps
	// Equipment should belong to step???
	Attribute("equipment", "EquipmentList", "List of tools needed")

	//Attribute("directions", CollectionOf("application/recipe.step+json"), "List of steps")
	// Attribute("steps", "StepList", "List of steps")
	// Directions might need to be an array of grouped steps i.e. prep_steps cook_steps, plate_steps..
	// Directions might need to be a "workflow", multiple steps, some can be completed async

	Attribute("notes", ArrayOf(Note), "List of dated notes")

	// meta data
	Attribute("favorite", Boolean, "Is a favorite, basically a tag")
	//Attribute("tags", CollectionOf("application/recipe.tag+json"), "List of tags")
	Attribute("tags", TagList)
	Attribute("rating", Float32, "rating between 0-1", func() {
		Minimum(0)
		Maximum(1)
	})
	Attribute("difficulty", Float32, "rating between 0-1", func() {
		Minimum(0)
		Maximum(1)
	})
	Attribute("source", String, "URL Source of recipe")

	Required("title")
})

// Ingredient recipe or individual component and a quantity.
var Ingredient = Type("Ingredient", func() {
	Attribute("recipe", Recipe, "")
	// might need level and/or parent?
	Attribute("quantity", String, "TODO: make UnitOfMeasure")
})

// IngredientList is an array of ingredients. Can't use "ArrayOf" directly otherwise we get an initialization error.
var IngredientList = Type("IngredientList", func() {
	Attribute("ingredients", ArrayOf(Ingredient), "List of ingredients")
})

// Recipe ...
var Recipe = ResultType("application/vnd.rx.recipe", func() {
	Description("A recipe")
	ContentType("application/json")
	Reference(RecipePayload)
	Attributes(func() {
		Attribute("id", String, "Unique recipe ID")
		Attribute("title")
		Attribute("description")
		// Attribute("images")
		// Attribute("quantity", UnitOfMeasure, "quantity, measure, servings, yield e.g. 4 cups.")
		// Attribute("prep_time", DateTime, "Amount of time to prepare")
		// Attribute("cook_time", DateTime, "Amount of time to cook")
		// Attribute("wait_time", DateTime, "Amount of time to wait for things such as mairnading")
		// Attribute("cookware", CollectionOf("application/recipe.cookware+json"), "List of tools needed")
		Attribute("version", String, "Version Number e.g. 1.0.1")
		// Attribute("ingredients", CollectionOf("application/recipe.recipe+json"), "List of ingredients")
		Attribute("ingredients")
		// Attribute("ingredients", "IngredientList", "List of ingredients")
		// Attribute("directions", CollectionOf("application/recipe.step+json"), "List of steps") // ??? might need to be an array of grouped steps i.e. prep_steps cook_steps, plate_steps..
		// Attribute("categories", CollectionOf("application/recipe.category+json"), "List of categories, basically same as tag")
		Attribute("favorite", Boolean, "Is a favorite, basically a tag")
		// Attribute("tags", CollectionOf("application/recipe.tag+json"), "List of tags")
		Attribute("rating", Float32, "rating between 0-1", func() {
			Minimum(0)
			Maximum(1)
		})
		Attribute("difficulty", Float32, "rating between 0-1", func() {
			Minimum(0)
			Maximum(1)
		})
		// Attribute("source", Source, "Source of recipe")
		// Attribute("notes", CollectionOf("application/recipe.note+json"), "List of dated notes")
		Attribute("state", String, "e.g. chopped, sliced, etc.. might need to be array.")
		Attribute("complete", Boolean, "If it's been added/included")
		// Attribute("created", DateTime, "First created")
		// Attribute("updated", DateTime, "Last Updated")

		Required("id", "title")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		// Attribute("images")
		// Attribute("quantity")
		// Attribute("prep_time")
		// Attribute("cook_time")
		// Attribute("wait_time")
		// Attribute("cookware")
		// Attribute("version")
		// Attribute("ingredients")
		// Attribute("directions")
		// Attribute("categories")
		// Attribute("favorite")
		// Attribute("tags")
		// Attribute("rating")
		// Attribute("difficulty")
		// Attribute("source")
		// Attribute("notes")
		// Attribute("created")
		// Attribute("updated")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		// Attribute("quantity")
		// Attribute("state")
		// Attribute("complete")
	})

	View("ingredient", func() {
		Attribute("title")
		// Attribute("quantity")
		// Attribute("state")
		// Attribute("complete")
	})
})

// TimeList is an array of times. Can't use "ArrayOf" directly otherwise we get an initialization error.
var TimeList = Type("TimeList", func() {
	Attribute("time", ArrayOf(Time), "List of ingredients")
})

// Time is an array of times. Can't use "ArrayOf" directly otherwise we get an initialization error.
var Time = Type("Time", func() {
	Attribute("time", "DateTime", "List of ingredients")
	Attribute("name", "String", "List of ingredients")
})

// Equipment ...
var Equipment = Type("Equipment", func() {
	Description("A specific tool")
	Attributes(func() {
		Attribute("id", String, "Unique ID")
		Attribute("description", String, "long description")
		Attribute("parts", "EquipmentList", "list of parts or attachments")
		Attribute("settings", ArrayOf(String), "settings, e.g. temprature")
		// Attribute("setup", CollectionOf("application/recipe.step+json"), "Steps to setting up")
		Attribute("complete", Boolean, "If it's been checked")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("name")
		Attribute("description")
		Attribute("parts")
		Attribute("settings")
		Attribute("setup")
	})
})

// EquipmentList is an array of Equipment. Can't use "ArrayOf" directly otherwise we get an initialization error.
var EquipmentList = Type("EquipmentList", func() {
	Attribute("equipment", ArrayOf(Equipment), "List of ingredients")
})

// TagList is an array of tags. Can't use "ArrayOf" directly otherwise we get an initialization error.
var TagList = Type("TagList", func() {
	Attribute("tag", ArrayOf(Tag), "List of tags")
})

// TTag ...
var TTag = Type("Tag", func() {
	Description("A Tag")
	Attributes(func() {
		Attribute("id", String, "Unique ID")
		Attribute("name", String, "tag name")
		Attribute("created", String, "TODO: DateTime")
		Attribute("updated", String, "TODO: DateTime")

		Required("id", "note")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("created")
		Attribute("updated")
	})
})

// Note ...
var Note = Type("Note", func() {
	Description("A Note")
	Attributes(func() {
		Attribute("id", String, "Unique ID")
		Attribute("note", String, "Note text")
		Attribute("created", String, "TODO: DateTime")
		Attribute("updated", String, "TODO: DateTime")

		Required("id", "note")
	})
	View("default", func() {
		Attribute("id")
		Attribute("note")
		Attribute("created")
		Attribute("updated")
	})
})

// recipe ...
// var _ = Resource("recipe", func() {
// 	BasePath("/recipe")
// 	DefaultMedia(RecipeMedia)

// 	Action("list", func() {
// 		Description("List recipes")
// 		Scheme("http")
// 		Routing(GET("/"))
// 		Response(OK, "text/html")
// 		Response(InternalServerError, ErrorMedia)
// 	})

// 	Action("show", func() {
// 		Description("Display an recipe by id")
// 		Routing(GET("/:id"))
// 		Params(func() {
// 			Param("id", String, "Recipe ID")
// 		})
// 		Response(OK, "text/html")
// 		Response(Created, "text/html")
// 		Response(InternalServerError, ErrorMedia)
// 		Response(NotFound)
// 	})

// 	Action("update", func() {
// 		Description("")
// 		Routing(PATCH("/:id"))
// 		Params(func() {
// 			Param("id", String, "Recipe ID")
// 		})

// 		Payload("RecipePayload")

// 		Response(NoContent)
// 		Response(OK)
// 		Response(NotFound)
// 		Response(BadRequest, ErrorMedia)
// 		Response(InternalServerError, ErrorMedia)
// 	})

// 	Action("create", func() {
// 		Routing(POST("/"))
// 		Payload("RecipePayload")
// 		Description("")

// 		Response(OK)
// 		Response(InternalServerError, ErrorMedia)
// 		Response(Created)
// 	})

// 	Action("delete", func() {
// 		Description("")
// 		Routing(DELETE("/:id"))
// 		Params(func() {
// 			Param("id", String, "Recipe ID")
// 		})
// 		Response(NoContent)
// 		Response(NotFound)
// 		Response(InternalServerError, ErrorMedia)
// 	})
// })

// var _ = Resource("image", func() {
// 	BasePath("/images")

// 	Action("upload", func() {
// 		Routing(POST("/"))
// 		Description("Upload an image")
// 		Response(OK, ImageMedia)
// 	})

// 	Action("show", func() {
// 		Routing(GET("/:id"))
// 		Description("Show an image metadata")
// 		Params(func() {
// 			Param("id", String, "Image ID")
// 		})
// 		Response(OK, ImageMedia)
// 		Response(NotFound)
// 	})

// 	Files("/download/*filename", "images/") // Serve files from the "images" directory
// })
