package design

// import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = API("rx", func() {
	Title("Recipe Service")
	Description("HTTP service for managing your recipes")
	// Server("http://localhost:8080")
})
