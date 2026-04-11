import { useState } from 'react'
import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { Badge } from '../ui/Badge.jsx'
import { AssetModal } from '../ui/AssetModal.jsx'
import { formatIDR } from '../utils/format.js'

function StocksIDSection({ data, loading }) {
  const [selected, setSelected] = useState(null)

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">🇮🇩</span>Saham Indonesia (IDX)</div>
          <div className="section-subtitle">Bursa Efek Indonesia — Blue Chip</div>
        </div>
        <div className="section-meta">{data.length > 0 && `${data.length} emiten`}</div>
      </div>
      <div className="section-content">
        {loading ? <Skeleton rows={8} /> : (
          <table className="section-table">
            <thead>
              <tr>
                <th>Kode</th>
                <th>Nama</th>
                <th>Harga (IDR)</th>
                <th>Perubahan</th>
                <th>%</th>
              </tr>
            </thead>
            <tbody>
              {data.map((item) => (
                <tr key={item.symbol} onClick={() => setSelected(item)} className="clickable-row">
                  <td><span className="asset-symbol">{item.symbol}</span></td>
                  <td><div className="asset-name">{item.name || item.symbol}</div></td>
                  <td className="num">{formatIDR(item.price)}</td>
                  <td className={`num ${item.change >= 0 ? 'positive' : 'negative'}`}>
                    {item.change >= 0 ? '+' : ''}{formatIDR(Math.abs(item.change))}
                  </td>
                  <td><Badge value={item.change_percent} /></td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>

      {selected && (
        <AssetModal
          type="stock_id"
          symbol={selected.symbol}
          title={`${selected.symbol} — ${selected.name || ''}`}
          color={selected.change_percent >= 0 ? '#0ecb81' : '#f6465d'}
          onClose={() => setSelected(null)}
        />
      )}
    </div>
  )
}

export { StocksIDSection }
