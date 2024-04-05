package postgres

func (psqlDB *PsqlDB) createUserTable() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);	
	`

	_, err := psqlDB.DB.Exec(createUserTable)

	if err != nil {
		panic("Create users table fail " + err.Error())
	}
}
