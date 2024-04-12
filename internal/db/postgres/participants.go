package postgres

func (psqlDB *PsqlDB) createParticipantTable() {
	createParticipantTable := `
	CREATE TABLE IF NOT EXISTS participants (
		id SERIAL PRIMARY KEY,
		userId INT REFERENCES users(id),
		roomId INT REFERENCES rooms(id),
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP
	);
	`

	_, err := psqlDB.DB.Exec(createParticipantTable)
	if err != nil {
		panic("Create room table fail: " + err.Error())
	}
}
