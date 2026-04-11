import { useState } from 'react'
import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { Badge } from '../ui/Badge.jsx'
import { formatIDR } from '../utils/format.js'

const TYPES = ['Semua', 'Saham', 'Pendapatan Tetap', 'Campuran', 'Pasar Uang']

function ReksadanaSection({ data, loading }) {
  const [filter, setFilter] = useState('Semua')

  const filtered = filter === 'Semua'
    ? data
    : data.filter(d => d.type === filter)

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">📈</span>Reksadana</div>
          <div className="section-subtitle">Performa reksadana populer Indonesia</div>
        </div>
        <div className="section-meta">{data.length > 0 && `${data.length} produk`}</div>
      </div>
      <div className="section-content">
        <div className="filter-tabs">
          {TYPES.map(t => (
            <button
              key={t}
              className={`filter-tab${filter === t ? ' active' : ''}`}
              onClick={() => setFilter(t)}
            >
              {t}
            </button>
          ))}
        </div>

        {loading ? <Skeleton rows={8} /> : (
          <table className="section-table">
            <thead>
              <tr>
                <th>Nama Reksadana</th>
                <th>Jenis</th>
                <th>NAB/Unit</th>
                <th>1 Bulan</th>
                <th>YTD</th>
                <th>1 Tahun</th>
              </tr>
            </thead>
            <tbody>
              {filtered.map((item, i) => (
                <tr key={i}>
                  <td>
                    <div className="asset-name">{item.name}</div>
                    {item.data_source === 'sample' && (
                      <span className="badge-sample">sample</span>
                    )}
                  </td>
                  <td>
                    <span className={`type-badge type-${item.type?.toLowerCase().replace(/\s+/g, '-')}`}>
                      {item.type}
                    </span>
                  </td>
                  <td className="num">{formatIDR(item.nab_per_unit)}</td>
                  <td><Badge value={item.return_1m} /></td>
                  <td><Badge value={item.return_ytd} /></td>
                  <td><Badge value={item.return_1y} /></td>
                </tr>
              ))}
              {filtered.length === 0 && (
                <tr>
                  <td colSpan={6} className="muted" style={{textAlign:'center', padding:'24px'}}>
                    Tidak ada data untuk kategori ini
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        )}
      </div>
    </div>
  )
}

export { ReksadanaSection }
