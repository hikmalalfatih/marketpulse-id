import './TickerBar.css'

function TickerBar({ crypto = [], forex = {} }) {
  const idrRate = forex?.rates?.IDR || 0

  const items = []

  // Top 10 crypto
  crypto.slice(0, 10).forEach(c => {
    const up = c.change_24h >= 0
    const arrow = up ? '▲' : '▼'
    const cls = up ? 'tick-up' : 'tick-down'
    const price = c.price >= 1
      ? `$${c.price.toLocaleString('en-US', { maximumFractionDigits: 2 })}`
      : `$${c.price.toFixed(6)}`
    items.push(
      <span key={c.id} className="ticker-item">
        <span className="tick-sym">{c.symbol?.toUpperCase()}</span>
        <span className="tick-price">{price}</span>
        <span className={cls}>{arrow}{Math.abs(c.change_24h).toFixed(2)}%</span>
      </span>
    )
  })

  // USD/IDR
  if (idrRate > 0) {
    items.push(
      <span key="usdidr" className="ticker-item">
        <span className="tick-sym">USD/IDR</span>
        <span className="tick-price">Rp {idrRate.toLocaleString('id-ID', { maximumFractionDigits: 0 })}</span>
      </span>
    )
  }

  if (items.length === 0) {
    return (
      <div className="ticker-bar">
        <div className="ticker-content-static">
          <span className="tick-sym">maley exchange</span>
          <span className="tick-price">Memuat data pasar...</span>
        </div>
      </div>
    )
  }

  // Duplicate for seamless loop
  const doubled = [...items, ...items]

  return (
    <div className="ticker-bar">
      <div className="ticker-track">
        <div className="ticker-content">
          {doubled}
        </div>
      </div>
    </div>
  )
}

export { TickerBar }
