package author

const (
	queryInsertNewAuthor = `
		INSERT INTO authors
		(
			name,
			bio,
			birth_date
		) VALUES (?, ?, ?)
	`

	queryFindAuthorByID = `
		SELECT
			id,
			name,
			bio,
			birth_date,
			death_date
		FROM authors
		WHERE id = ?
	`

	queryFindAllAuthor = `
		SELECT
			id,
			name,
			bio,
			birth_date,
			death_date
		FROM authors
		ORDER BY updated_at DESC
		LIMIT ?
		OFFSET ?
	`

	queryUpdateNewAuthor = `
		UPDATE authors
		SET
			name = ?,
			bio = ?,
			birth_date = ?,
			death_date = ?,
			updated_at = NOW()
		WHERE id = ?
	`

	queryDeleteAuthorByID = `
		DELETE FROM authors WHERE id = ?
	`
)
