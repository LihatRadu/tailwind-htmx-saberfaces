package main

import (
	//"saberfaces/sabers"
	"saberfaces/sabers"
  //"encoding/json"
//	"tailwind-htmx-saberfaces/sabers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main()  {
  println("Hello Saber!")

  engine := html.New("./templates/", ".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Get("/", func (c *fiber.Ctx) error {
    //return c.SendString("Hello Sabers!")
    sabersData := map[string][]sabers.Saber{
     "sabers": sabers.GetSabers(),
    }
    return c.Render("index", sabersData)
  })
  
  app.Get("/sabers", func (c *fiber.Ctx) error {
    //return c.SendString("Hello Sabers!")
    sabersData := map[string][]sabers.Saber{
     "sabers": sabers.GetSabers(),
    }
   // return c.Render("sabers", fiber.Map{"message":"Saber List"})
    return c.Render("sabers", sabersData)
  })
  app.Post("/sabers", func (c *fiber.Ctx) error {
     
    return c.Render("sabers", fiber.Map{"message":"Saber List"})
  })
  app.Get("/sabersfaces", func (c *fiber.Ctx) error {
    //return c.SendString("Hello Sabers!")
    return c.Render("sabers", fiber.Map{"":""})
  })
  app.Static("/assets", "./assets/")
  sabers.Saberfaces()
  
  app.Listen(":3000")

}
