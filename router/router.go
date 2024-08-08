package router

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Data struct{
	Shimpments []Shipment `json:"shipments"`
}
type Shipment struct {
	Sender	`json:"sender"`
	Recipient	`json:"recipient"`
	Origin	`json:"origin"`
	Destination	`json:"destination"`
	Items []Items	`json:"items"`
	Payments	`json:"payments"`
	Options	`json:"options"`
	DeliveryType string `json:"deliveryType"`
}

type Sender struct{
	FirstName string 	`json:"firstName"`
	LastName string	`json:"lastName"`
	Email string	`json:"email"`
	Phone string	`json:"phone"`
}
type Recipient struct{
	FirstName string 	`json:"firstName"`
	LastName string	`json:"lastName"`
	Email string	`json:"email"`
	Phone string	`json:"phone"`
}

type Origin struct{
	Address string	`json:"address"`
	Reference string	`json:"reference"`
	Country string	`json:"country"`
	City string	`json:"city"`
	ZipCode string	`json:"zipCode"`
	ZipCode2 string	`json:"zipcode"`
  }

type Destination struct{
	Address string	`json:"address"`
	Reference string	`json:"reference"`
	Country string	`json:"country"`
	City string	`json:"city"`
	ZipCode string	`json:"zipCode"`
	ZipCode2 string	`json:"zipcode"`
}  

type Items struct{
	Size string	`json:"size"`
	Weight int	`json:"weight"`
}

type Payments struct{
	Insured bool	`json:"insured"`
  }  
type Options struct{
  	RequiresSignature bool	`json:"requiresSignature"`
	TwoFactorAuth bool	`json:"twoFactorAuth"`
	RequiresIdentification bool	`json:"requiresIdentification"`
	Notes string `json:"notes"`
  }

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { 

		var body Data
		c.BodyParser(&body)

		org_zipcode := body.Shimpments[0].Origin.ZipCode
		if org_zipcode != ""{
			body.Shimpments[0].Origin.ZipCode2 =org_zipcode
		}

		dst_zipcode := body.Shimpments[0].Destination.ZipCode
		if org_zipcode != ""{
			body.Shimpments[0].Origin.ZipCode2 =dst_zipcode
		}		
		
		autorizacion := c.Get("Authorization")
		if autorizacion == ""{
			c.Status(fiber.StatusBadGateway).JSON("Acceso no autorizado")
		}
		
		bodyBytes,err :=  json.Marshal(body)
		if err != nil{
			panic(err.Error())
		}		

		r, err := http.NewRequest("POST", "https://sandbox.99minutos.com/api/v3/orders", bytes.NewBuffer(bodyBytes))
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", autorizacion)
		if err != nil {
			panic(err)
		}
        
		client := &http.Client{}
			res, err := client.Do(r)
			if err != nil {
				panic(err)
		}


		defer res.Body.Close()
		var Response map[string]interface{}
		derr := json.NewDecoder(res.Body).Decode(&Response)
		if derr != nil {
			panic(derr)
		}

		return  c.Status(res.StatusCode).JSON(Response)
	})

	apiRoutes := app.Group("/api")
	v1APIRoutes := apiRoutes.Group("/v1")
	{
		// IMPLEMENT ROUTE
		v1APIRoutes.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })
	}
}
