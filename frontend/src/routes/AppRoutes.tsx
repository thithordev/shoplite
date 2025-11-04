import { Route, Routes, Navigate } from 'react-router-dom'
import CustomersPage from '@pages/Customers/CustomersPage'
import CustomerDetail from '@pages/Customers/CustomerDetail'
import ProductsPage from '@pages/Products/ProductsPage'
import ProductDetail from '@pages/Products/ProductDetail'
import OrdersPage from '@pages/Orders/OrdersPage'
import OrderDetail from '@pages/Orders/OrderDetail'

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/customers" replace />} />
      <Route path="/customers" element={<CustomersPage />} />
      <Route path="/customers/:id" element={<CustomerDetail />} />
      <Route path="/products" element={<ProductsPage />} />
      <Route path="/products/:id" element={<ProductDetail />} />
      <Route path="/orders" element={<OrdersPage />} />
      <Route path="/orders/:id" element={<OrderDetail />} />
      <Route path="*" element={<Navigate to="/customers" replace />} />
    </Routes>
  )
}
