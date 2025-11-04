import axios from 'axios'
import { notification } from 'antd'

const client = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  headers: { 'Content-Type': 'application/json' },
})

client.interceptors.response.use(
  (res) => res,
  (error) => {
    const message = error?.response?.data?.message || error.message || 'Request error'
    notification.error({ message: 'Error', description: message })
    return Promise.reject(error)
  }
)

export default client
