import { Typography } from 'antd'

type Props = {
  title: string
  extra?: React.ReactNode
}

export default function PageHeader({ title, extra }: Props) {
  return (
    <div style={{ display: 'flex', alignItems: 'center', gap: 12, marginBottom: 16 }}>
      <Typography.Title level={3} style={{ margin: 0, flex: 1 }}>
        {title}
      </Typography.Title>
      <div>{extra}</div>
    </div>
  )
}
