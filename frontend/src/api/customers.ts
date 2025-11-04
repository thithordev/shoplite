import client from './client'
import type { Customer, CreateCustomerPayload, ID } from '@utils/types'

export async function getCustomers(): Promise<Customer[]> {
  const { data } = await client.get('/customers')
  return data?.data ?? []
}

export async function getCustomer(id: ID): Promise<Customer> {
  const { data } = await client.get(`/customers/${id}`)
  return data?.data
}

export async function createCustomer(payload: CreateCustomerPayload): Promise<Customer> {
  const { data } = await client.post('/customers', payload)
  return data?.data
}
