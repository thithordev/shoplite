import { useMemo, useState } from 'react'
import { Button, Flex, Input } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import DataTable from '@components/DataTable'
import PageHeader from '@components/PageHeader'
import { useQuery } from '@tanstack/react-query'
import { getProducts } from '@api/products'
import type { Product } from '@utils/types'
import { useNavigate } from 'react-router-dom'
import { formatMoney } from '@utils/format'

export default function ProductsPage() {
  const navigate = useNavigate()
  const [search, setSearch] = useState('')
  const { data = [], isLoading } = useQuery({ queryKey: ['products'], queryFn: getProducts })

  const filtered = useMemo(() => {
    const q = search.trim().toLowerCase()
    if (!q) return data
    return data.filter((p) => p.name.toLowerCase().includes(q))
  }, [data, search])

  return (
    <>
      <PageHeader
        title="Products"
        extra={
          <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/products/new')}>
            New
          </Button>
        }
      />
      <Flex style={{ marginBottom: 12 }} gap={8}>
        <Input.Search
          placeholder="Search products"
          allowClear
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          style={{ maxWidth: 320 }}
        />
      </Flex>
      <DataTable<Product>
        rowKey="id"
        loading={isLoading}
        dataSource={filtered}
        pagination={{ pageSize: 10 }}
        onRow={(record) => ({ onClick: () => navigate(`/products/${record.id}`) })}
        columns={[
          { title: 'ID', dataIndex: 'id', width: 80 },
          { title: 'Name', dataIndex: 'name' },
          { title: 'Price', dataIndex: 'price', render: (v) => formatMoney(v) },
          { title: 'Stock', dataIndex: 'stock' },
        ]}
      />
    </>
  )
}
