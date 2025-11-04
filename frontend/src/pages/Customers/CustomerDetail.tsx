import { Alert, Button, Form, Input, Space, notification } from 'antd'
import { useNavigate, useParams } from 'react-router-dom'
import { useEffect, useState } from 'react'
import PageHeader from '@components/PageHeader'
import { createCustomer, getCustomer } from '@api/customers'
import * as yup from 'yup'

const schema = yup.object({
  name: yup.string().required('Name is required'),
  email: yup.string().email('Invalid email').required('Email is required'),
})

export default function CustomerDetail() {
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
          const data = await getCustomer(Number(id))
          form.setFieldsValue({ name: data.name, email: data.email })
        } catch (e) {
          notification.error({ message: 'Failed to load customer' })
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
      await createCustomer(values)
      notification.success({ message: 'Customer saved' })
      navigate('/customers')
    } catch (err: any) {
      if (err?.inner) {
        // map yup errors into antd form errors
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
      <PageHeader title={isEdit ? 'Customer Detail' : 'New Customer'} />
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
        <Form.Item label="Email" name="email" rules={[{ type: 'email' }, { required: true }]}>
          <Input placeholder="Enter email" />
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
