package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Account is the account resource media type.
var Account = MediaType("application/vnd.account+json", func() {
	Description("A tenant account")
	UseTrait("jsonapi") // Will set content type to "application/vnd.api+json"
	Attributes(func() {
		Attribute("data", func() {
			Attribute("type", String, func() {
				Default("account")
			})
			Attribute("id", Integer, "ID of account")
			Attribute("attributes", func() {
				Attribute("name", String, "Name of account")
				Attribute("created_at", DateTime, "Date of creation")
				Attribute("created_by", String, "Email of account owner", func() {
					Format("email")
				})
				//Required("id", "href", "name")
			})
		})
	})
	View("default", func() {
		Attribute("data", func() {
			Attribute("id", Integer, "ID")
			Attribute("type", String, "type")
			Attribute("attributes", func() {
				Attribute("name")
				Attribute("created_at")
				Attribute("created_by")
			})
		})
	})
	View("tiny", func() {
		Attribute("data", func() {
			Attribute("id", Integer, "ID")
			Attribute("type", String, "type")
			Attribute("attributes", func() {
				Attribute("name")
			})
		})
	})
})

// Bottle is the bottle resource media type.
var Bottle = MediaType("application/vnd.bottle+json", func() {
	UseTrait("jsonapi") // Will set content type to "application/vnd.api+json"
	Description("A bottle of wine")
	Reference(BottlePayload)
	Attributes(func() {
		Attribute("data", func() {
			Attribute("id", Integer, "ID of bottle")
			Attribute("type", String, func() {
				Default("bottle")
			})
			Attribute("attributes", func() {
				Attribute("href", String, "API href of bottle")
				Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
					Minimum(1)
					Maximum(5)
				})
				Attribute("created_at", DateTime, "Date of creation")
				Attribute("updated_at", DateTime, "Date of last update")
				// Attributes below inherit from the base type
				Attribute("name")
				Attribute("vineyard")
				Attribute("varietal")
				Attribute("vintage")
				Attribute("color")
				Attribute("sweetness")
				Attribute("country")
				Attribute("region")
				Attribute("review")

				//Required("id", "href", "name", "vineyard", "varietal", "vintage", "color")
			})
			//Required("id", "href", "name", "vineyard", "varietal", "vintage", "color")
		})
		Attribute("relationships", func() {
			Attribute("account", Account, "Account that owns bottle")
		})
		Attribute("links", func() {
			Attribute("self", String, "API href of bottle")
			Attribute("related", String, "API href of account")
		})
	})

	View("default", func() {
		Attribute("data", func() {
			Attribute("id")
			Attribute("type")
			Attribute("attributes", func() {
				Attribute("href", String, "API href of bottle")
				Attribute("rating")
				Attribute("account", Account, "Account that owns bottle")
				Attribute("created_at", DateTime, "Date of creation")
				Attribute("updated_at", DateTime, "Date of last update")
				// Attributes below inherit from the base type
				Attribute("name")
				Attribute("vineyard")
				Attribute("varietal")
				Attribute("vintage")
				Attribute("color")
				Attribute("sweetness")
				Attribute("country")
				Attribute("region")
				Attribute("review")
				Attribute("account", func() {
					View("tiny")
				})

				//Required("href", "name", "vineyard", "varietal", "vintage", "color")
			})
			Required("id")
		})
		Attribute("relationships", func() {
			Attribute("account", Account, "Account that owns bottle")
		})
		Attribute("links", func() {
			Attribute("self", String, "API href of bottle")
			Attribute("related", String, "API href of account")
		})
	})

	View("tiny", func() {
		Attribute("data", func() {
			Attribute("id")
			Attribute("type")
			Attribute("attributes", func() {
				Attribute("name")
				Attribute("rating")
			})
		})
	})

	View("full", func() {
		Attribute("data", func() {
			Attribute("id")
			Attribute("type")
			Attribute("attributes", func() {
				Attribute("name")
				Attribute("account")
				Attribute("rating")
				Attribute("vineyard")
				Attribute("varietal")
				Attribute("vintage")
				Attribute("color")
				Attribute("sweetness")
				Attribute("country")
				Attribute("region")
				Attribute("review")
				Attribute("created_at")
				Attribute("updated_at")
			})
		})
	})
})
