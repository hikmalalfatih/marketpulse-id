import { useState } from 'react'
import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { Badge } from '../ui/Badge.jsx'
import { MiniSparkline } from '../charts/MiniSparkline.jsx'
import { AssetModal } from '../ui/AssetModal.jsx'
import { formatUSD, formatMarketCap, formatVolume } from '../utils/format.js'

function CryptoSection({ data, loading }) {
  const [selected, setSelected] = useState(null)

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">🪙</span>Kripto</div>
          <div className="section-subtitle">Top 20 aset kripto berdasarkan market cap</div>
        </div>
        <div className="section-meta">{data.length > 0 && `${data.length} aset`}</div>
      </div>
      <div className="section-content">
        {loading ? <Skeleton rows={10} /> : (
          <table className="section-table">
            <thead>
              <tr>
                <th style={{width:32}}>#</th>
                <th>Aset</th>
                <th>Harga (USD)</th>
                <th>1j</th>
                <th>24j</th>
                <th>7h</th>
                <th>Market Cap</th>
                <th>Volume 24j</th>
                <th>7 Hari</th>
              </tr>
            </thead>
            <tbody>
              {data.map((item, i) => (
                <tr key={item.id} onClick={() => setSelected(item)} className="clickable-row">
                  <td className="muted num">{i + 1}</td>
                  <td>
                    <div className="asset-cell">
                      {item.image_url && (
                        <img src={item.image_url} alt={item.name} className="asset-img" />
                      )}
                      <div>
                        <div className="asset-name">{item.name}</div>
                        <span className="asset-symbol">{item.symbol?.toUpperCase()}</span>
                      </div>
                    </div>
                  </td>
                  <td className="num">${formatUSD(item.price, item.price < 1 ? 6 : 2)}</td>
                  <td><Badge value={item.change_1h} /></td>
                  <td><Badge value={item.change_24h} /></td>
                  <td><Badge value={item.change_7d} /></td>
                  <td className="num muted">${formatMarketCap(item.market_cap)}</td>
                  <td className="num muted">${formatVolume(item.volume_24h)}</td>
                  <td>
                    <MiniSparkline
                      data={(item.sparkline_7d || []).map(v => ({ close: v }))}
                      color={item.change_7d >= 0 ? '#0ecb81' : '#f6465d'}
                    />
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>

      {selected && (
        <AssetModal
          type="crypto"
          symbol={selected.id}
          title={`${selected.name} (${selected.symbol?.toUpperCase()})`}
          color={selected.change_24h >= 0 ? '#0ecb81' : '#f6465d'}
          onClose={() => setSelected(null)}
        />
      )}
    </div>
  )
}

export { CryptoSection }
