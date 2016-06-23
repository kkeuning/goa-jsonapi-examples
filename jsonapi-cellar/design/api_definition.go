package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("jsonapi-cellar", func() {
	Title("A json:api goa cellar example")
	Description("API based on the goa-cellar and examples in the jsonapi.org spec")
	Contact(func() {
		Name("goa team")
		Email("admin@goa.design")
		URL("http://goa.design")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("goa guide")
		URL("http://goa.design/getting-started.html")
	})
	Host("cellar.goa.design")
	Scheme("http")
	BasePath("/cellar")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})

	BasicAuthSecurity("admin_pass", func() {
		Description(`Basic authentication method, for global admin authentication.

Here are very secret credentials:
* username: wine
* password: lover
`)
	})
	Trait("jsonapi", func() {
		ContentType("application/vnd.api+json")
	})
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
