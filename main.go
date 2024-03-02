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
	for _, movie := range movies {
		fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}

	newMovie := createMovie(db)
	movieId :=  newMovie.ID

	// Printing list of movies after creation
	movies = readMovies(db)
	for _, movie := range movies {
		fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}

	// Printing a single movie
	var movie Movie
	movie = readMovie(db, movieId)
	fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))


	updateMovie(db, movieId, "Interstellar")

	// Checking update
	movie = readMovie(db, movieId)
	fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))


	deleteMovie(db, movieId)

	// Checking state after deletion
	movies = readMovies(db)
	for _, movie := range movies {
		fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}

	db.Close()

}