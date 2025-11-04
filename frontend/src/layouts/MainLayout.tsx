import { Layout, Menu, theme } from 'antd'
import { ShoppingCartOutlined, TeamOutlined, AppstoreOutlined } from '@ant-design/icons'
import { Link, useLocation } from 'react-router-dom'
import { ReactNode, useMemo } from 'react'

const { Header, Sider, Content } = Layout

export default function MainLayout({ children }: { children: ReactNode }) {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()
  const location = useLocation()

  const selectedKey = useMemo(() => {
    if (location.pathname.startsWith('/customers')) return 'customers'
    if (location.pathname.startsWith('/products')) return 'products'
    if (location.pathname.startsWith('/orders')) return 'orders'
    return 'customers'
  }, [location.pathname])

  return (
    <Layout>
      <Sider breakpoint="lg" collapsedWidth="0">
        <div style={{ padding: 16, color: 'white', fontWeight: 600 }}>ShopLite</div>
        <Menu theme="dark" mode="inline" selectedKeys={[selectedKey]}
          items={[
            { key: 'customers', icon: <TeamOutlined />, label: <Link to="/customers">Customers</Link> },
            { key: 'products', icon: <AppstoreOutlined />, label: <Link to="/products">Products</Link> },
            { key: 'orders', icon: <ShoppingCartOutlined />, label: <Link to="/orders">Orders</Link> },
          ]}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }} />
        <Content style={{ margin: '24px 16px 0' }}>
          <div style={{ padding: 24, minHeight: 360, background: colorBgContainer, borderRadius: borderRadiusLG }}>
            {children}
          </div>
        </Content>
      </Layout>
    </Layout>
  )
}
