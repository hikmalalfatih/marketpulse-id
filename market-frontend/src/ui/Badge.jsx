function Badge({ value, suffix = '%' }) {
  const isPositive = value >= 0
  const absValue = Math.abs(value).toFixed(2)
  const arrow = isPositive ? '▲' : '▼'
  const className = isPositive ? 'positive' : 'negative'

  return (
    <span className={`badge ${className}`}>
      {arrow} {absValue}{suffix}
    </span>
  )
}

export { Badge }

