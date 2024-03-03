package main

import "database/sql"


func createMovie (db *sql.DB) Movie {
	newMovie := Movie{0, "Inception", 2006}
	res, err := db.Exec("INSERT INTO movies VALUES(null,?,?);", newMovie.Title, newMovie.Year)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return readMovie(db, int(id))
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


func updateMovie (db *sql.DB, id int, newTitle string) Movie {
	_, err := db.Exec("UPDATE movies SET title=? WHERE id=?", newTitle, id)
	if err != nil {
		panic(err)
	}


	return readMovie(db, id)
}

func deleteMovie  (db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM movies WHERE id=?;", id)
	if err != nil {
		panic(err)
	}
}