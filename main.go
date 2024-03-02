package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)


type Movie struct {
	ID int
	Title string
	Year int
}

func openDb () *sql.DB {
	const file string = "movies.db"
	db, err := sql.Open("sqlite3", file)
	
	if err != nil {
		panic(err)
	}
	return db
}

func createTableIfNotExists (db *sql.DB) {
	const create string = `
		CREATE TABLE IF NOT EXISTS "movies" (
		"id"	INTEGER,
		"title"	TEXT NOT NULL,
		"year"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`
	_, err := db.Exec(create);
	if err != nil {
		panic(err)
	}
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

	createMovie(db)

	// Printing list of movies after creation
	movies = readMovies(db)
	for _, movie := range movies {
		fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}

	movieId :=  movies[0].ID

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