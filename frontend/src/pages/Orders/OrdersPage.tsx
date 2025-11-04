import { useMemo, useState } from 'react'
import { Button, Flex, Input, Tag } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import DataTable from '@components/DataTable'
import PageHeader from '@components/PageHeader'
import { useQuery } from '@tanstack/react-query'
import { getOrders } from '@api/orders'
import type { Order } from '@utils/types'
import { useNavigate } from 'react-router-dom'
import { formatDate } from '@utils/format'

export default function OrdersPage() {
  const navigate = useNavigate()
  const [search, setSearch] = useState('')
  const { data = [], isLoading } = useQuery({ queryKey: ['orders'], queryFn: getOrders })

  const filtered = useMemo(() => {
    const q = search.trim().toLowerCase()
    if (!q) return data
    return data.filter((o) =>
      (o.customer?.name || '').toLowerCase().includes(q) || (o.status || '').toLowerCase().includes(q)
    )
  }, [data, search])

  return (
    <>
      <PageHeader
        title="Orders"
        extra={
          <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/orders/new')}>
            New
          </Button>
        }
      />
      <Flex style={{ marginBottom: 12 }} gap={8}>
        <Input.Search
          placeholder="Search orders by customer or status"
          allowClear
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          style={{ maxWidth: 420 }}
        />
      </Flex>
      <DataTable<Order>
        rowKey="id"
        loading={isLoading}
        dataSource={filtered}
        pagination={{ pageSize: 10 }}
        onRow={(record) => ({ onClick: () => navigate(`/orders/${record.id}`) })}
        columns={[
          { title: 'ID', dataIndex: 'id', width: 80 },
          { title: 'Customer', dataIndex: ['customer', 'name'] },
          { title: 'Status', dataIndex: 'status', render: (v) => <Tag>{v}</Tag> },
          { title: 'Date', dataIndex: 'order_date', render: (v) => formatDate(v) },
        ]}
      />
    </>
  )
}
