package code

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchItems(t *testing.T) {
	category := "phones"
	sqlDB, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	rows := sqlmock.
		NewRows([]string{"item_name"}).
		AddRow("Samsung S9").AddRow("iPhone X")
	sqlMock.
		ExpectQuery(`SELECT item_name FROM items WHERE category = \?`).
		WithArgs(category).WillReturnRows(rows)

	items, err := FetchItems(sqlDB, category)

	require.NoError(t, err)
	assert.Equal(t, []string{"A", "B"}, items)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}