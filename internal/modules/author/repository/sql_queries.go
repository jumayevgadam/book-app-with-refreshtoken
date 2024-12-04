package repository

// sql queries for authors.
const (
	createAuthorQuery = `
		INSERT INTO authors (
			username, 
			email, 
			phone_number, 
			password, 
			bio, 
			avatar
		) VALUES (
			$1, 
			$2,
			$3,
			$4,
			COALESCE(NULLIF($5, ''), bio),
			COALESCE(NULLIF($6, ''), avatar)
		) RETURNING
		 	id;`
)
