import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { Badge } from '../ui/Badge.jsx'
import { MiniSparkline } from '../charts/MiniSparkline.jsx'
import { formatIDR, formatUSD, formatMarketCap } from '../utils/format.js'

function SummarySection({ crypto, stocksID, stocksUS, forex, gold, oil, loading, onNavigate }) {
  const idrRate = forex?.rates?.IDR || 0
  const btc = crypto.find(c => c.id === 'bitcoin') || crypto[0]
  const eth = crypto.find(c => c.id === 'ethereum') || crypto[1]
  const brent = oil?.find(o => o.name?.toLowerCase().includes('brent')) || oil?.[0]

  const topGainers = [...crypto].sort((a, b) => b.change_24h - a.change_24h).slice(0, 5)
  const topLosers  = [...crypto].sort((a, b) => a.change_24h - b.change_24h).slice(0, 5)

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">📊</span>Ringkasan Pasar</div>
          <div className="section-subtitle">Snapshot semua aset terkini</div>
        </div>
      </div>
      <div className="section-content">
        {loading ? <Skeleton rows={3} type="card" /> : (
          <>
            {/* Quick stats row */}
            <div className="summary-stats">
              {idrRate > 0 && (
                <div className="stat-pill" onClick={() => onNavigate('forex')} style={{cursor:'pointer'}}>
                  <span className="stat-label">USD/IDR</span>
                  <span className="stat-value num">{formatIDR(idrRate)}</span>
                </div>
              )}
              {gold?.price_usd_per_oz > 0 && (
                <div className="stat-pill" onClick={() => onNavigate('commodity')} style={{cursor:'pointer'}}>
                  <span className="stat-label">Emas/oz</span>
                  <span className="stat-value num">${formatUSD(gold.price_usd_per_oz)}</span>
                </div>
              )}
              {brent && (
                <div className="stat-pill" onClick={() => onNavigate('oil')} style={{cursor:'pointer'}}>
                  <span className="stat-label">Brent</span>
                  <span className="stat-value num">${formatUSD(brent.price_usd)}/bbl</span>
                </div>
              )}
            </div>

            {/* Main cards */}
            <div className="summary-grid">
              {btc && (
                <div className="summary-card" onClick={() => onNavigate('crypto')} style={{cursor:'pointer'}}>
                  <div className="summary-card-header">
                    {btc.image_url && <img src={btc.image_url} alt="BTC" className="summary-coin-img" />}
                    <span className="summary-card-label">Bitcoin</span>
                  </div>
                  <div className="summary-card-value num">${formatUSD(btc.price)}</div>
                  <div className="summary-card-idr num muted">{formatIDR(btc.price * idrRate)}</div>
                  <div className="summary-card-change"><Badge value={btc.change_24h} /></div>
                  <MiniSparkline
                    data={(btc.sparkline_7d || []).map(v => ({ close: v }))}
                    color={btc.change_24h >= 0 ? '#0ecb81' : '#f6465d'}
                  />
                </div>
              )}

              {eth && (
                <div className="summary-card" onClick={() => onNavigate('crypto')} style={{cursor:'pointer'}}>
                  <div className="summary-card-header">
                    {eth.image_url && <img src={eth.image_url} alt="ETH" className="summary-coin-img" />}
                    <span className="summary-card-label">Ethereum</span>
                  </div>
                  <div className="summary-card-value num">${formatUSD(eth.price)}</div>
                  <div className="summary-card-idr num muted">{formatIDR(eth.price * idrRate)}</div>
                  <div className="summary-card-change"><Badge value={eth.change_24h} /></div>
                  <MiniSparkline
                    data={(eth.sparkline_7d || []).map(v => ({ close: v }))}
                    color={eth.change_24h >= 0 ? '#0ecb81' : '#f6465d'}
                  />
                </div>
              )}

              {gold?.price_usd_per_oz > 0 && (
                <div className="summary-card" onClick={() => onNavigate('commodity')} style={{cursor:'pointer'}}>
                  <div className="summary-card-header">
                    <span style={{fontSize:'20px'}}>🥇</span>
                    <span className="summary-card-label">Emas (XAU)</span>
                  </div>
                  <div className="summary-card-value num">${formatUSD(gold.price_usd_per_oz)}/oz</div>
                  <div className="summary-card-idr num muted">{formatIDR(gold.price_idr_per_gram)}/gram</div>
                  <div className="summary-card-change"><Badge value={gold.change_percent} /></div>
                </div>
              )}

              {brent && (
                <div className="summary-card" onClick={() => onNavigate('oil')} style={{cursor:'pointer'}}>
                  <div className="summary-card-header">
                    <span style={{fontSize:'20px'}}>🛢️</span>
                    <span className="summary-card-label">Brent Crude</span>
                  </div>
                  <div className="summary-card-value num">${formatUSD(brent.price_usd)}/bbl</div>
                  <div className="summary-card-change"><Badge value={brent.change_percent} /></div>
                </div>
              )}

              {stocksID.slice(0, 2).map(s => (
                <div key={s.symbol} className="summary-card" onClick={() => onNavigate('stocks-id')} style={{cursor:'pointer'}}>
                  <div className="summary-card-header">
                    <span className="asset-symbol" style={{fontSize:'13px'}}>{s.symbol}</span>
                    <span className="summary-card-label">{s.name?.split(' ').slice(0,2).join(' ')}</span>
                  </div>
                  <div className="summary-card-value num">{formatIDR(s.price)}</div>
                  <div className="summary-card-change"><Badge value={s.change_percent} /></div>
                </div>
              ))}
            </div>

            {/* Top movers */}
            {crypto.length > 0 && (
              <div className="movers-grid">
                <div className="movers-panel">
                  <div className="movers-title positive">▲ Top Gainers 24j</div>
                  <table className="section-table">
                    <tbody>
                      {topGainers.map(c => (
                        <tr key={c.id}>
                          <td>
                            <div className="asset-cell">
                              {c.image_url && <img src={c.image_url} alt={c.name} className="asset-img" />}
                              <span className="asset-symbol">{c.symbol?.toUpperCase()}</span>
                            </div>
                          </td>
                          <td className="num">${formatUSD(c.price, c.price < 1 ? 4 : 2)}</td>
                          <td><Badge value={c.change_24h} /></td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
                <div className="movers-panel">
                  <div className="movers-title negative">▼ Top Losers 24j</div>
                  <table className="section-table">
                    <tbody>
                      {topLosers.map(c => (
                        <tr key={c.id}>
                          <td>
                            <div className="asset-cell">
                              {c.image_url && <img src={c.image_url} alt={c.name} className="asset-img" />}
                              <span className="asset-symbol">{c.symbol?.toUpperCase()}</span>
                            </div>
                          </td>
                          <td className="num">${formatUSD(c.price, c.price < 1 ? 4 : 2)}</td>
                          <td><Badge value={c.change_24h} /></td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  )
}

export { SummarySection }
