import { useMemo, useState } from 'react'
import { Button, Flex, Input } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import DataTable from '@components/DataTable'
import PageHeader from '@components/PageHeader'
import { useQuery } from '@tanstack/react-query'
import { getCustomers } from '@api/customers'
import type { Customer } from '@utils/types'
import { useNavigate } from 'react-router-dom'

export default function CustomersPage() {
  const navigate = useNavigate()
  const [search, setSearch] = useState('')
  const { data = [], isLoading } = useQuery({ queryKey: ['customers'], queryFn: getCustomers })

  const filtered = useMemo(() => {
    const q = search.trim().toLowerCase()
    if (!q) return data
    return data.filter((c) => c.name.toLowerCase().includes(q) || c.email.toLowerCase().includes(q))
  }, [data, search])

  return (
    <>
      <PageHeader
        title="Customers"
        extra={
          <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/customers/new')}>
            New
          </Button>
        }
      />
      <Flex style={{ marginBottom: 12 }} gap={8}>
        <Input.Search
          placeholder="Search customers"
          allowClear
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          style={{ maxWidth: 320 }}
        />
      </Flex>
      <DataTable<Customer>
        rowKey="id"
        loading={isLoading}
        dataSource={filtered}
        pagination={{ pageSize: 10 }}
        onRow={(record) => ({ onClick: () => navigate(`/customers/${record.id}`) })}
        columns={[
          { title: 'ID', dataIndex: 'id', width: 80 },
          { title: 'Name', dataIndex: 'name' },
          { title: 'Email', dataIndex: 'email' },
        ]}
      />
    </>
  )
}
