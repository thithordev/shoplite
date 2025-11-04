// Shared TypeScript interfaces for API entities

export type ID = number

export interface Customer {
  id: ID
  name: string
  email: string
  created_at?: string
  updated_at?: string
}

export interface Product {
  id: ID
  name: string
  price: number
  stock: number
  created_at?: string
  updated_at?: string
}

export interface OrderItem {
  id: ID
  order_id: ID
  product_id: ID
  quantity: number
  price: number
  product?: Product
  created_at?: string
  updated_at?: string
}

export interface Order {
  id: ID
  customer_id: ID
  customer?: Customer
  order_date: string
  status: 'pending' | 'paid' | 'shipped' | 'cancelled'
  items: OrderItem[]
  created_at?: string
  updated_at?: string
}

// Requests
export interface CreateCustomerPayload { name: string; email: string }
export interface CreateProductPayload { name: string; price: number; stock: number }
export interface CreateOrderItemPayload { product_id: ID; quantity: number; price: number }
export interface CreateOrderPayload {
  customer_id: ID
  order_date: string
  status?: Order['status']
  items: CreateOrderItemPayload[]
}
