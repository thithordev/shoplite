import dayjs from 'dayjs'

export const formatDate = (iso?: string) => (iso ? dayjs(iso).format('YYYY-MM-DD HH:mm') : '')
export const formatMoney = (n?: number) => (typeof n === 'number' ? `$${n.toFixed(2)}` : '$0.00')
