import { useEffect, useMemo, useState } from 'react'
import { Alert, Button, Form, InputNumber, Select, Space, Typography, notification } from 'antd'
import PageHeader from '@components/PageHeader'
import { useNavigate, useParams } from 'react-router-dom'
import { createOrder, getOrder } from '@api/orders'
import { getCustomers } from '@api/customers'
import { getProducts } from '@api/products'
import type { Customer, ID, Product } from '@utils/types'
import * as yup from 'yup'

const itemSchema = yup.object({
  product_id: yup.number().required(),
  quantity: yup.number().moreThan(0, 'Quantity must be > 0').required(),
  price: yup.number().moreThan(0, 'Price must be > 0').required(),
})
const schema = yup.object({
  customer_id: yup.number().required('Customer is required'),
  items: yup.array().of(itemSchema).min(1, 'Add at least one item'),
})

export default function OrderDetail() {
  const { id } = useParams()
  const isEdit = id && id !== 'new'
  const navigate = useNavigate()
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)

  const [customers, setCustomers] = useState<Customer[]>([])
  const [products, setProducts] = useState<Product[]>([])

  useEffect(() => {
    ;(async () => {
      try {
        setLoading(true)
        const [cs, ps] = await Promise.all([getCustomers(), getProducts()])
        setCustomers(cs)
        setProducts(ps)
        if (isEdit) {
          const o = await getOrder(Number(id))
          form.setFieldsValue({
            customer_id: o.customer_id,
            items: o.items.map((it) => ({ product_id: it.product_id, quantity: it.quantity, price: it.price })),
          })
        }
      } catch (e) {
        notification.error({ message: 'Failed to load data' })
      } finally {
        setLoading(false)
      }
    })()
  }, [id, isEdit, form])

  const items = Form.useWatch('items', form) as { product_id?: ID; quantity?: number; price?: number }[] | undefined
  const total = useMemo(() => {
    return (items || []).reduce((sum, it) => sum + (it.quantity || 0) * (it.price || 0), 0)
  }, [items])

  const productPriceMap = useMemo(() => Object.fromEntries(products.map((p) => [p.id, p.price])), [products])

  const onProductChange = (index: number, productId: ID) => {
    // set price to default product price when selecting
    const price = productPriceMap[productId]
    const curr = form.getFieldValue(['items', index]) || {}
    form.setFieldsValue({ items: { [index]: { ...curr, product_id: productId, price } } })
  }

  const onSubmit = async () => {
    try {
      const values = await form.validateFields()
      await schema.validate(values, { abortEarly: false })
      setLoading(true)
      await createOrder({
        customer_id: values.customer_id,
        order_date: new Date().toISOString(),
        status: 'pending',
        items: values.items,
      })
      notification.success({ message: 'Order placed' })
      navigate('/orders')
    } catch (err: any) {
      if (err?.inner) {
        const errors = err.inner.map((e: any) => ({ name: e.path, errors: [e.message] }))
        form.setFields(errors)
      } else if (!isEdit) {
        notification.error({ message: err.message || 'Failed to place order' })
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <>
      <PageHeader title={isEdit ? 'Order Detail' : 'New Order'} />
      {isEdit && (
        <Alert type="info" showIcon style={{ marginBottom: 16 }} message="Editing is view-only in this demo (no update endpoint)." />
      )}

      <Form layout="vertical" form={form} disabled={loading || !!isEdit} style={{ maxWidth: 800 }}>
        <Form.Item label="Customer" name="customer_id" rules={[{ required: true, message: 'Select customer' }]}>
          <Select
            placeholder="Select customer"
            options={customers.map((c) => ({ label: c.name, value: c.id }))}
            showSearch
            optionFilterProp="label"
          />
        </Form.Item>

        <Form.List name="items" initialValue={[{ product_id: undefined, quantity: 1, price: 0 }]}>
          {(fields, { add, remove }) => (
            <div>
              {fields.map((field, index) => (
                <Space key={field.key} style={{ display: 'flex', marginBottom: 8 }} align="baseline">
                  <Form.Item
                    {...field}
                    style={{ minWidth: 260 }}
                    name={[field.name, 'product_id']}
                    rules={[{ required: true, message: 'Select product' }]}
                  >
                    <Select
                      placeholder="Select product"
                      options={products.map((p) => ({ label: p.name, value: p.id }))}
                      onChange={(v) => onProductChange(index, v)}
                      showSearch
                      optionFilterProp="label"
                      style={{ minWidth: 240 }}
                    />
                  </Form.Item>
                  <Form.Item {...field} name={[field.name, 'quantity']} rules={[{ required: true }]}> 
                    <InputNumber min={1} step={1} placeholder="Qty" />
                  </Form.Item>
                  <Form.Item {...field} name={[field.name, 'price']} rules={[{ required: true }]}> 
                    <InputNumber min={0.01} step={0.01} placeholder="Price" />
                  </Form.Item>
                  {fields.length > 1 && (
                    <Button danger onClick={() => remove(field.name)}>Remove</Button>
                  )}
                </Space>
              ))}
              <Button onClick={() => add({ quantity: 1 })}>Add item</Button>
            </div>
          )}
        </Form.List>

        <Typography.Paragraph style={{ marginTop: 12 }}>
          <strong>Total:</strong> ${total.toFixed(2)}
        </Typography.Paragraph>

        {!isEdit && (
          <Space>
            <Button type="primary" onClick={onSubmit} loading={loading}>
              Place Order
            </Button>
            <Button onClick={() => navigate(-1)}>Cancel</Button>
          </Space>
        )}
      </Form>
    </>
  )
}
