import client from './client'
import type { Product, CreateProductPayload, ID } from '@utils/types'

export async function getProducts(): Promise<Product[]> {
  const { data } = await client.get('/products')
  return data?.data ?? []
}

export async function getProduct(id: ID): Promise<Product> {
  const { data } = await client.get(`/products/${id}`)
  return data?.data
}

export async function createProduct(payload: CreateProductPayload): Promise<Product> {
  const { data } = await client.post('/products', payload)
  return data?.data
}
