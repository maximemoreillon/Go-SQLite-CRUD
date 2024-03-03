package main

import (
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)


type Movie struct {
	ID int
	Title string
	Year int
}


func main () {


	db := openDb()

	createTableIfNotExists(db)

	// Printing list of movies before insertion
	var movies []Movie
	movies = readMovies(db)
	fmt.Println("Current list of movies:")
	for _, movie := range movies {
		fmt.Printf("- [%s] %s (%s)\n\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}
	fmt.Println("")

	// Creating a movie
	newMovie := createMovie(db)
	movieId :=  newMovie.ID
	fmt.Printf("Created [%s] %s (%s)\n\n", strconv.Itoa(newMovie.ID), newMovie.Title, strconv.Itoa(newMovie.Year))

	// Printing list of movies after creation
	movies = readMovies(db)
	fmt.Println("Current list of movies:")
	for _, movie := range movies {
		fmt.Printf("- [%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}
	fmt.Println("")

	// Printing a single movie
	movie := readMovie(db, movieId)
	fmt.Printf("Queried movie: [%s] %s (%s)\n\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))

	// Updating a movie
	updatedMovie := updateMovie(db, movieId, "Interstellar")
	fmt.Printf("Updated movie [%s] %s (%s)\n\n", strconv.Itoa(updatedMovie.ID), updatedMovie.Title, strconv.Itoa(updatedMovie.Year))

	// Deleting a movie
	deleteMovie(db, movieId)

	// Checking state after deletion
	movies = readMovies(db)
	fmt.Println("Current list of movies:")
	for _, movie := range movies {
		fmt.Printf("- [%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}
	fmt.Println("")

	db.Close()

}