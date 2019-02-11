package main

import (
	"html/template"
	//"log"
	//"net/http"
	"os"
)

type Recipe struct {
	Name         string
	Ingredients  []string
	Instructions []string
}

var test Recipe = Recipe{
	Name: "Almond Flour Biscuit",
	Ingredients: []string{
		"2 cup Blanched almond flour",
		"2 tsp Gluten-free baking powder",
		"1/2 tsp Sea salt",
		"2 large Egg (beaten)",
		"1/3 cup Butter",
	},
	Instructions: []string{
		"Preheat the oven to 350 degrees F (177 degrees C).",
		"Line a baking sheet with parchment paper.",
		"Mix dry ingredients together in a large bowl.",
		"Stir in wet ingredients.",
		"Scoop tablespoonfuls of the dough onto the lined baking sheet (a cookie scoop is the fastest way).",
		"Form into rounded biscuit shapes (flatten slightly with your fingers).",
		"Bake for about 15 minutes, until firm and golden.",
		"Cool on the baking sheet.",
	},
}

func main() {
	tmp, err := template.ParseFiles("./static/template.html")
	if err != nil {
		panic(err)
	}
	tmp.ExecuteTemplate(os.Stdout, "recipe", test)
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
