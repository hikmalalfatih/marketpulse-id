import { LineChart, Line, ResponsiveContainer } from 'recharts'

function MiniSparkline({ data = [], color = '#0ecb81', height = 32 }) {
  if (!data || data.length < 2) {
    return <div style={{ height, width: 80, background: 'transparent' }} />
  }

  return (
    <ResponsiveContainer width={80} height={height}>
      <LineChart data={data} margin={{ left: 0, right: 0, top: 2, bottom: 2 }}>
        <Line
          type="monotone"
          dataKey="close"
          stroke={color}
          strokeWidth={1.5}
          dot={false}
          isAnimationActive={false}
        />
      </LineChart>
    </ResponsiveContainer>
  )
}

export { MiniSparkline }
