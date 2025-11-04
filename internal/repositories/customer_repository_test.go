package repositories

import (
	"testing"

	"shoplite/internal/models"
	"shoplite/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestCustomerRepository_CRUD(t *testing.T) {
	db := testutil.NewTestDB(t)
	testutil.TruncateAll(t, db)
	repo := NewCustomerRepository(db)

	// Create
	c := &models.Customer{Name: "Alice", Email: "alice@example.com"}
	require.NoError(t, repo.Create(c))
	require.NotZero(t, c.ID)

	// FindAll
	list, err := repo.FindAll()
	require.NoError(t, err)
	require.Len(t, list, 1)

	// FindByID
	found, err := repo.FindByID(c.ID)
	require.NoError(t, err)
	require.Equal(t, "Alice", found.Name)

	// Not found
	_, err = repo.FindByID(9999)
	require.Error(t, err)
}
