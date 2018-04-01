package code

import "database/sql"

func FetchItems(db *sql.DB, category string) ([]string, error) {
	rows, err := db.Query(
		"SELECT item_name FROM items WHERE category = ?", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]string, 0)
	for rows.Next() {
		var item string
        if err := rows.Scan(&item); err != nil {
            return items, err
        }
        items = append(items, item)
	}
	return items, nil
}