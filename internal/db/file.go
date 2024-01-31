package db

func createImageTable() {
	createImagesTable := `
		CREATE TABLE IF NOT EXISTS images (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			url TEXT NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
		);
	`

	_, err := DB.Exec(createImagesTable)
	if err != nil {
		panic("Create images table fail: " + err.Error())
	}

	createTagsTable := `
		CREATE TABLE IF NOT EXISTS tags (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
		);
	`

	_, err = DB.Exec(createTagsTable)
	if err != nil {
		panic("Create tags table fail: " + err.Error())
	}

	createImageTagTable := `
		CREATE TABLE IF NOT EXISTS image_tag (
			id SERIAL PRIMARY KEY,
			imageId INTEGER NOT NULL,
			tagId INTEGER NOT NULL,
			FOREIGN KEY (imageId) REFERENCES images(id),
			FOREIGN KEY (tagId) REFERENCES tags(id)
		);
	`

	_, err = DB.Exec(createImageTagTable)
	if err != nil {
		panic("Create image_tag table fail: " + err.Error())
	}
}
