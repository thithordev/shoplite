package repositories

import (
	"testing"

	"shoplite/internal/models"
	"shoplite/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestProductRepository_CRUD(t *testing.T) {
	db := testutil.NewTestDB(t)
	testutil.TruncateAll(t, db)
	repo := NewProductRepository(db)

	// Create
	p := &models.Product{Name: "Widget", Price: 9.99, Stock: 10}
	require.NoError(t, repo.Create(p))
	require.NotZero(t, p.ID)

	// FindAll
	list, err := repo.FindAll()
	require.NoError(t, err)
	require.Len(t, list, 1)

	// FindByID
	found, err := repo.FindByID(p.ID)
	require.NoError(t, err)
	require.Equal(t, "Widget", found.Name)

	// Not found
	_, err = repo.FindByID(9999)
	require.Error(t, err)
}
