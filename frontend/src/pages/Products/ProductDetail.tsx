import { Alert, Button, Form, Input, InputNumber, Space, notification } from 'antd'
import { useNavigate, useParams } from 'react-router-dom'
import { useEffect, useState } from 'react'
import PageHeader from '@components/PageHeader'
import { createProduct, getProduct } from '@api/products'
import * as yup from 'yup'

const schema = yup.object({
  name: yup.string().required('Name is required'),
  price: yup.number().moreThan(0, 'Price must be > 0').required('Price is required'),
  stock: yup.number().min(0, 'Stock must be >= 0').required('Stock is required'),
})

export default function ProductDetail() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const isEdit = id && id !== 'new'

  useEffect(() => {
    if (isEdit) {
      ;(async () => {
        try {
          setLoading(true)
          const data = await getProduct(Number(id))
          form.setFieldsValue({ name: data.name, price: data.price, stock: data.stock })
        } catch (e) {
          notification.error({ message: 'Failed to load product' })
        } finally {
          setLoading(false)
        }
      })()
    }
  }, [id, isEdit, form])

  const onSubmit = async () => {
    try {
      const values = await form.validateFields()
      await schema.validate(values, { abortEarly: false })
      setLoading(true)
      await createProduct(values)
      notification.success({ message: 'Product saved' })
      navigate('/products')
    } catch (err: any) {
      if (err?.inner) {
        const errors = err.inner.map((e: any) => ({ name: e.path, errors: [e.message] }))
        form.setFields(errors)
      } else if (!isEdit) {
        notification.error({ message: err.message || 'Failed to save' })
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <>
      <PageHeader title={isEdit ? 'Product Detail' : 'New Product'} />
      {isEdit && (
        <Alert
          type="info"
          showIcon
          style={{ marginBottom: 16 }}
          message="Editing is view-only in this demo (no update endpoint)."
        />
      )}
      <Form layout="vertical" form={form} disabled={loading || !!isEdit} style={{ maxWidth: 560 }}>
        <Form.Item label="Name" name="name" rules={[{ required: true, message: 'Name is required' }]}>
          <Input placeholder="Enter name" />
        </Form.Item>
        <Form.Item label="Price" name="price" rules={[{ required: true }]}>
          <InputNumber min={0.01} step={0.01} style={{ width: '100%' }} placeholder="Enter price" />
        </Form.Item>
        <Form.Item label="Stock" name="stock" rules={[{ required: true }]}>
          <InputNumber min={0} step={1} style={{ width: '100%' }} placeholder="Enter stock" />
        </Form.Item>
        {!isEdit && (
          <Space>
            <Button type="primary" onClick={onSubmit} loading={loading}>
              Save
            </Button>
            <Button onClick={() => navigate(-1)}>Cancel</Button>
          </Space>
        )}
      </Form>
    </>
  )
}
