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


func createMovie (db *sql.DB) {
	newMovie := Movie{0, "Inception", 2006}
	_, err := db.Exec("INSERT INTO movies VALUES(null,?,?);", newMovie.Title, newMovie.Year)
	if err != nil {
		panic(err)
	}
}


func readMovies (db *sql.DB) []Movie {
	movies := []Movie {}
	rows, err := db.Query("SELECT * FROM movies ORDER BY id DESC LIMIT 100")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		movie := Movie{}
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Year)
		if err != nil {
			panic(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func readMovie (db *sql.DB, id int) Movie {
	foundMovie := Movie{}
	row:= db.QueryRow("SELECT * FROM movies WHERE id=?", id)
	row.Scan(&foundMovie.ID, &foundMovie.Title, &foundMovie.Year)
	return foundMovie
}


func updateMovie (db *sql.DB, id int, newTitle string) {
	_, upErr := db.Exec("UPDATE movies SET title=? WHERE id=?", newTitle, id)
	if upErr != nil {
		panic(upErr)
	}
}

func deleteMovie  (db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM movies WHERE id=?;", id)
	if err != nil {
		panic(err)
	}
}

func main () {
	const file string = "movies.db"
	db, err := sql.Open("sqlite3", file)
	
	if err != nil {
		panic(err)
	}

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

	// Printing a single movie
	var movie Movie
	movie = readMovie(db, movies[0].ID)
	fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))


	updateMovie(db, movies[0].ID, "Interstellar")

	// Checking update
	movie = readMovie(db, movies[0].ID)
	fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))


	deleteMovie(db, movies[0].ID)

	// Checking state after deletion
	movies = readMovies(db)
	for _, movie := range movies {
		fmt.Printf("[%s] %s (%s)\n", strconv.Itoa(movie.ID), movie.Title, strconv.Itoa(movie.Year))
	}

}