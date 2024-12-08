package repository

// sql queries for authors.
const (
	createAuthorQuery = `
		INSERT INTO authors (
			username, 
			email, 
			password, 
			bio, 
			avatar
		) VALUES (																																																				
			$1, 
			$2,
			$3,
			NULLIF($4, ''),
			NULLIF($5, '')
		) RETURNING
		 	id;`

	// getAuthorQuery is.
	getAuthorQuery = `
		SELECT 
			id, 
			username, 
			email, 
			password, 
			COALESCE(bio, '') AS bio, 
			COALESCE(avatar, '') AS avatar
		FROM 
			authors
		WHERE
			id = $1;`

	// countAuthorsQuery is.
	countAuthorsQuery = `
		SELECT COUNT(id) FROM authors;`

	// listAuthorQuery is.
	listAuthorQuery = `
		SELECT
			id, username, email, password, COALESCE(bio, '') AS bio, COALESCE(avatar, '') AS avatar
		FROM	
			authors
		ORDER BY id DESC OFFSET $1 LIMIT $3;`

	// updateAuthorQuery is.
	updateAuthorQuery = `
		UPDATE authors
		SET 
			username = COALESCE(NULLIF($1, ''), username),
			email = COALESCE(NULLIF($2, ''), email),
			password = COALESCE(NULLIF($3, ''), password),
			bio = COALESCE(NULLIF($4, ''), bio),
			avatar = COALESCE(NULLIF($5, ''), avatar)
		WHERE 
			id = $6;`

	// deleteAuthorQuery is.
	deleteAuthorQuery = `
		DELETE 
			FROM authors
		WHERE 
			id = $1;`
)
