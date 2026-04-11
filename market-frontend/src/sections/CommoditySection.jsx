import { useState } from 'react'
import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { Badge } from '../ui/Badge.jsx'
import { AssetModal } from '../ui/AssetModal.jsx'
import { formatIDR, formatUSD } from '../utils/format.js'

function CommoditySection({ gold, oil, loading, oilOnly }) {
  const [showGoldChart, setShowGoldChart] = useState(false)
  const [selectedOil, setSelectedOil] = useState(null)

  const getOilHistorySymbol = (item) => {
    const name = item.name?.toLowerCase() || ''
    const sym = item.symbol?.toLowerCase() || ''
    if (name.includes('brent') || sym.includes('bz')) return 'brent'
    if (name.includes('wti') || sym.includes('cl')) return 'wti'
    if (name.includes('gas') || sym.includes('ng')) return 'natgas'
    return 'brent'
  }

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title">
            <span className="section-icon">{oilOnly ? '🛢️' : '⚡'}</span>
            {oilOnly ? 'Minyak & Gas' : 'Emas, Perak & Minyak'}
          </div>
          <div className="section-subtitle">Harga komoditas global real-time</div>
        </div>
      </div>
      <div className="section-content">
        {loading ? <Skeleton rows={3} type="card" /> : (
          <div className="commodity-layout">
            {/* Gold & Silver */}
            {!oilOnly && gold && gold.price_usd_per_oz > 0 && (
              <div className="commodity-group">
                <div className="commodity-group-title">Logam Mulia</div>
                <div className="commodity-cards">
                  <div className="commodity-card clickable-row" onClick={() => setShowGoldChart(true)}>
                    <div className="commodity-card-header">
                      <span className="commodity-icon">🥇</span>
                      <div>
                        <div className="commodity-card-name">Emas (XAU)</div>
                        <div className="commodity-card-sub">per troy oz</div>
                      </div>
                      {gold.change_percent !== 0 && (
                        <div className="commodity-badge">
                          <Badge value={gold.change_percent} />
                        </div>
                      )}
                    </div>
                    <div className="commodity-price-usd">${formatUSD(gold.price_usd_per_oz)}</div>
                    <div className="commodity-price-idr">{formatIDR(gold.price_idr_per_oz)}</div>
                    <div className="commodity-detail-row">
                      <span className="muted">Per gram</span>
                      <span className="num">{formatIDR(gold.price_idr_per_gram)}</span>
                    </div>
                    <div className="commodity-chart-hint">Klik untuk lihat grafik →</div>
                  </div>

                  {gold.silver_usd_per_oz > 0 && (
                    <div className="commodity-card">
                      <div className="commodity-card-header">
                        <span className="commodity-icon">🥈</span>
                        <div>
                          <div className="commodity-card-name">Perak (XAG)</div>
                          <div className="commodity-card-sub">per troy oz</div>
                        </div>
                      </div>
                      <div className="commodity-price-usd">${formatUSD(gold.silver_usd_per_oz)}</div>
                      <div className="commodity-detail-row">
                        <span className="muted">Per gram</span>
                        <span className="num">{formatIDR(gold.silver_idr_per_gram)}</span>
                      </div>
                    </div>
                  )}
                </div>
              </div>
            )}

            {/* Oil & Gas */}
            {oil && oil.length > 0 && (
              <div className="commodity-group">
                <div className="commodity-group-title">Energi</div>
                <div className="commodity-cards">
                  {oil.map((item) => (
                    <div
                      key={item.symbol}
                      className="commodity-card clickable-row"
                      onClick={() => setSelectedOil(item)}                    >
                      <div className="commodity-card-header">
                        <span className="commodity-icon">🛢️</span>
                        <div>
                          <div className="commodity-card-name">{item.name}</div>
                          <div className="commodity-card-sub">{item.unit}</div>
                        </div>
                        <div className="commodity-badge">
                          <Badge value={item.change_percent} />
                        </div>
                      </div>
                      <div className="commodity-price-usd">${formatUSD(item.price_usd)}</div>
                      {item.price_idr > 0 && (
                        <div className="commodity-price-idr">{formatIDR(item.price_idr)}</div>
                      )}
                      <div className="commodity-chart-hint">Klik untuk lihat grafik →</div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        )}
      </div>

      {showGoldChart && (
        <AssetModal
          type="gold"
          symbol="gold"
          title="Emas (XAU/USD) — 90 Hari"
          color="#f0b90b"
          onClose={() => setShowGoldChart(false)}
        />
      )}

      {selectedOil && (
        <AssetModal
          type="oil"
          symbol={getOilHistorySymbol(selectedOil)}
          title={`${selectedOil.name} — 90 Hari`}
          color="#3b82f6"
          onClose={() => setSelectedOil(null)}
        />
      )}
    </div>
  )
}

export { CommoditySection }
