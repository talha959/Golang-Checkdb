package tableexists

import "database/sql"

func Table(db *sql.DB, tableName string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_name = $1
        )
    `

	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
