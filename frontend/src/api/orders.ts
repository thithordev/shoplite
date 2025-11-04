import client from './client'
import type { Order, CreateOrderPayload, ID } from '@utils/types'

export async function getOrders(): Promise<Order[]> {
  const { data } = await client.get('/orders')
  return data?.data ?? []
}

export async function getOrder(id: ID): Promise<Order> {
  const { data } = await client.get(`/orders/${id}`)
  return data?.data
}

export async function createOrder(payload: CreateOrderPayload): Promise<Order> {
  const { data } = await client.post('/orders', payload)
  return data?.data
}
