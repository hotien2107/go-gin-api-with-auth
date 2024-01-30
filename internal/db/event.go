package db

func createEventsTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			location TEXT NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
		);
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Create events table fail: " + err.Error())
	}
}
