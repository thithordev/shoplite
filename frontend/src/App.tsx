import { ConfigProvider, Layout } from 'antd'
import { AppRoutes } from '@routes/AppRoutes'
import MainLayout from '@layouts/MainLayout'

export default function App() {
  return (
    <ConfigProvider theme={{ token: { colorPrimary: '#1677ff' } }}>
      <Layout style={{ minHeight: '100vh' }}>
        <MainLayout>
          <AppRoutes />
        </MainLayout>
      </Layout>
    </ConfigProvider>
  )
}
