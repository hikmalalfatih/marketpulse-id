import {
  AreaChart, Area, XAxis, YAxis, CartesianGrid,
  Tooltip, ResponsiveContainer
} from 'recharts'

function PriceChart({ data = [], title, color = '#0ecb81' }) {
  if (data.length === 0) {
    return <div className="chart-empty">Tidak ada data grafik</div>
  }

  const CustomTooltip = ({ active, payload }) => {
    if (!active || !payload?.length) return null
    const p = payload[0].payload
    return (
      <div className="custom-tooltip">
        <div className="tooltip-date">{p.date}</div>
        <div className="tooltip-price">{p.close?.toLocaleString('en-US', { maximumFractionDigits: 2 })}</div>
        {p.open > 0 && <div className="tooltip-row"><span>Open</span><span>{p.open?.toLocaleString('en-US', { maximumFractionDigits: 2 })}</span></div>}
        {p.high > 0 && <div className="tooltip-row"><span>High</span><span>{p.high?.toLocaleString('en-US', { maximumFractionDigits: 2 })}</span></div>}
        {p.low  > 0 && <div className="tooltip-row"><span>Low</span><span>{p.low?.toLocaleString('en-US', { maximumFractionDigits: 2 })}</span></div>}
      </div>
    )
  }

  const gradId = `grad-${color.replace('#', '')}`

  return (
    <div className="chart-container">
      {title && <div className="chart-title">{title}</div>}
      <ResponsiveContainer width="100%" height={280}>
        <AreaChart data={data} margin={{ top: 8, right: 8, left: 0, bottom: 0 }}>
          <defs>
            <linearGradient id={gradId} x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%"   stopColor={color} stopOpacity={0.25} />
              <stop offset="100%" stopColor={color} stopOpacity={0.0}  />
            </linearGradient>
          </defs>
          <CartesianGrid vertical={false} strokeDasharray="3 3" stroke="var(--border)" />
          <XAxis
            dataKey="date"
            tickFormatter={d => {
              const dt = new Date(d)
              return dt.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
            }}
            tickLine={false}
            axisLine={false}
            tick={{ fontSize: 11, fill: 'var(--text-muted)' }}
            interval="preserveStartEnd"
          />
          <YAxis
            tickFormatter={v => v >= 1000 ? `${(v/1000).toFixed(0)}K` : v.toFixed(2)}
            tickLine={false}
            axisLine={false}
            tick={{ fontSize: 11, fill: 'var(--text-muted)' }}
            width={56}
          />
          <Tooltip content={<CustomTooltip />} />
          <Area
            type="monotone"
            dataKey="close"
            stroke={color}
            fill={`url(#${gradId})`}
            strokeWidth={2}
            isAnimationActive={false}
            dot={false}
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  )
}

export { PriceChart }
