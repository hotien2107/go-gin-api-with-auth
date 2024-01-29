package db

func createEventsTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
		)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Create table fail: " + err.Error())
	}
}
