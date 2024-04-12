package postgres

func (psqlDB *PsqlDB) createRoomTable() {
	createRoomTable := `
	CREATE TABLE IF NOT EXISTS rooms (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		createdUser INT REFERENCES users(id),
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP
	);
	`

	_, err := psqlDB.DB.Exec(createRoomTable)
	if err != nil {
		panic("Create room table fail: " + err.Error())
	}
}
