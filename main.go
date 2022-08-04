package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// recipe represents data about a record recipe.
type recipe struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Url         string `json:"url"`
	UserPicture string `json:"userPicture"`
	ImageUrl    string `json:"imageUrl"`
	Date        string `json:"date"`
	Votes       int    `json:"votes"`
}

// recipes slice to seed record recipe data.
var recipes = []recipe{
	{ID: "1", Title: "Vegan BBQ Burger with caramelized onions", Author: "Sumit", Votes: 23, Url: "https://www.hellofresh.de/recipes/undefined-61a61f91d7eb5d760943d835", UserPicture: "", Date: "January", ImageUrl: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/vegan-bbq-burger-mit-karamellisierten-zwiebeln-eaeffed7.jpg"},
	{ID: "2", Title: "Roast Veggie & Garlic Crouton Salad with Creamy Pesto, Parmesan & Almonds ", Author: "Renan", Votes: 20, Url: "https://www.hellofresh.com.au/recipes/roast-veggie-garlic-crouton-salad-61b96df10db0c84a862bacc1", UserPicture: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/roast-veggie-garlic-crouton-salad-7fd44915.jpg", Date: "February", ImageUrl: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/roast-veggie-garlic-crouton-salad-7fd44915.jpg"},
	{ID: "3", Title: "Teriyaki sesame chicken with sugar snap peas and fried rice", Author: "Fiona", Votes: 25, Url: "https://www.hellofresh.com/recipes/teriyaki-chicken-tenders-5b16cb16ae08b533b21a1862", UserPicture: "", Date: "March", ImageUrl: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/teriyaki-chicken-tenders-f2ea8fe8.jpg"},
	{ID: "4", Title: "Apple and blueberry tart with whipped cream and toasted almonds", Author: "Ouessanne", Votes: 17, Url: "#", UserPicture: "", Date: "April"},
	{ID: "5", Title: "Lemongrass Chicken Banh Mi Style Salad", Author: "Rebecca", Votes: 21, Url: "https://www.hellofresh.co.uk/recipes/lemongrass-chicken-banh-mi-style-salad-621cbe0dc366f8583b20d296", UserPicture: "", Date: "May", ImageUrl: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/lemongrass-chicken-banh-mi-style-salad-a5695221.jpg"},
	{ID: "6", Title: "Spicy halloumi tacos with caramelized onions and aioli", Author: "Alex", Votes: 23, Url: "https://www.hellofresh.de/recipes/wurzig-feurige-halloumi-tacos-6215f4639e130d1257780d3f", UserPicture: "", Date: "June", ImageUrl: "https://img.hellofresh.com/f_auto,fl_lossy,q_auto,w_1200/hellofresh_s3/image/wurzig-feurige-halloumi-tacos-fff0af46.jpg"},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/recipes", getRecipes)
	router.GET("/recipe/:id", getRecipeByID)
	router.POST("/recipes", postRecipes)

	router.Run("localhost:8080")
}

// getRecipes responds with the list of all recipes as JSON.
func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

// postRecipes adds a recipe from JSON received in the request body.
func postRecipes(c *gin.Context) {
	var newRecipe recipe

	// Call BindJSON to bind the received JSON to
	// newRecipe.
	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	// Add the new recipe to the slice.
	recipes = append(recipes, newRecipe)
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// getRecipeByID locates the recipe whose ID value matches the id
// parameter sent by the client, then returns that recipe as a response.
func getRecipeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of recipes, looking for
	// a recipe whose ID value matches the parameter.
	for _, a := range recipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
}
