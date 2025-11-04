package repositories

import (
	"testing"
	"time"

	"shoplite/internal/models"
	"shoplite/internal/testutil"

	"github.com/stretchr/testify/require"
)

func TestOrderRepository_Basic(t *testing.T) {
	db := testutil.NewTestDB(t)
	testutil.TruncateAll(t, db)
	custRepo := NewCustomerRepository(db)
	prodRepo := NewProductRepository(db)
	repo := NewOrderRepository(db)

	cust := &models.Customer{Name: "Bob", Email: "bob@example.com"}
	require.NoError(t, custRepo.Create(cust))
	prod := &models.Product{Name: "Gizmo", Price: 5.5, Stock: 5}
	require.NoError(t, prodRepo.Create(prod))

	ord := &models.Order{CustomerID: cust.ID, OrderDate: time.Now(), Status: "pending"}
	require.NoError(t, db.Create(ord).Error)
	item := &models.OrderItem{OrderID: ord.ID, ProductID: prod.ID, Quantity: 2, Price: 5.5}
	require.NoError(t, db.Create(item).Error)

	// FindAll
	orders, err := repo.FindAll()
	require.NoError(t, err)
	require.Len(t, orders, 1)
	require.Len(t, orders[0].Items, 1)

	// FindByID
	found, err := repo.FindByID(ord.ID)
	require.NoError(t, err)
	require.Equal(t, ord.ID, found.ID)
	require.Len(t, found.Items, 1)
}
