package petshop

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pet struct {
	Name    string
	Species string
}

var pets = []Pet{
	{"miaomiao", "cat"},
}

func filterPetList(ctx context.Context, pets []Pet, name string, species string) (ps []Pet) {
	if name == "" && species == "" {
		return pets
	}
	for i := range pets {
		if !strings.Contains(pets[i].Name, name) {
			continue
		}
		if species != "" && pets[i].Species != species {
			continue
		}
		ps = append(ps, pets[i])
	}
	return
}

func GetPetList(ctx context.Context, name string, species string) []Pet {
	return filterPetList(ctx, pets, name, species)
}

func StartGin() error {
	r := gin.Default()
	r.GET("/pets", func(c *gin.Context) {
		name := c.Query("name")
		species := c.Query("species")
		c.JSON(http.StatusOK, GetPetList(c.Request.Context(), name, species))
	})
	return r.Run()
}
