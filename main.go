package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type error struct {
	Error string `json:"error"`
}
type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	for _, item := range movies {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,error{"movie not found"})
}

func createMovie(c *gin.Context) {
	var movie Movie

	if err := c.BindJSON(&movie); err != nil {
		return
	}

	movies = append(movies, movie)
	c.IndentedJSON(http.StatusCreated, movie)
}

func updateMovie(c *gin.Context) {	
	var movie Movie
	err:= c.ShouldBindJSON(&movie)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,error{"bad_request"})
		return
	} else {
		params := c.Param("id")
		for index, item := range movies {
			if item.ID == params {
				cur_movie := &movies[index]
				cur_movie.Title = movie.Title
				cur_movie.Director = movie.Director

				c.IndentedJSON(http.StatusOK, movie)
				return
			}
		}
	}
}

func deleteMovie (c *gin.Context) {
	params := c.Param("id")

	for index, item := range movies {
		if item.ID == params {
			movies = append(movies[:index], movies[index+1:]...)
			c.IndentedJSON(http.StatusNoContent, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, error{"not_found"})
}

func main() {
	movies = append(movies, Movie{ID: "1", Title: "What We Do in the Shadows", Director: &Director{Firstname: "Taika", Lastname: "Waititi"}})
	movies = append(movies, Movie{ID: "2", Title: "Idioterne", Director: &Director{Firstname: "Lars", Lastname: "Trier"}})
	movies = append(movies, Movie{ID: "3", Title: " Chaising Amy", Director: &Director{Firstname: "Kevin", Lastname: "Smith"}})
	
	router := gin.Default()

	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", createMovie)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)
	
	http.ListenAndServe(":8080", router)

}
