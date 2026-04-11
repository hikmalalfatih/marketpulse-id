import { useState } from 'react'
import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { formatNumber } from '../utils/format.js'

const JENIS = ['Semua', 'SUN', 'ORI', 'SR', 'PBS']

function ObligasiSection({ data, loading }) {
  const [filter, setFilter] = useState('Semua')

  const filtered = filter === 'Semua'
    ? data
    : data.filter(d => d.jenis === filter)

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">📋</span>Obligasi</div>
          <div className="section-subtitle">Surat Berharga Negara (SBN) Indonesia</div>
        </div>
        <div className="section-meta">{data.length > 0 && `${data.length} instrumen`}</div>
      </div>
      <div className="section-content">
        <div className="filter-tabs">
          {JENIS.map(j => (
            <button
              key={j}
              className={`filter-tab${filter === j ? ' active' : ''}`}
              onClick={() => setFilter(j)}
            >
              {j}
            </button>
          ))}
        </div>

        {loading ? <Skeleton rows={6} /> : (
          <table className="section-table">
            <thead>
              <tr>
                <th>Kode</th>
                <th>Nama</th>
                <th>Jenis</th>
                <th>Kupon</th>
                <th>Jatuh Tempo</th>
                <th>Harga Pasar</th>
                <th>Yield</th>
              </tr>
            </thead>
            <tbody>
              {filtered.map((item, i) => (
                <tr key={i}>
                  <td><span className="asset-symbol">{item.kode}</span></td>
                  <td>
                    <div className="asset-name">{item.nama}</div>
                    {item.tenor && <span className="muted" style={{fontSize:'11px'}}>{item.tenor}</span>}
                  </td>
                  <td>
                    <span className={`type-badge type-${item.jenis?.toLowerCase()}`}>{item.jenis}</span>
                  </td>
                  <td className="num positive">{formatNumber(item.kupon, 2)}%</td>
                  <td className="muted">{item.jatuh_tempo}</td>
                  <td className="num">{formatNumber(item.harga_pasar, 2)}</td>
                  <td className="num">{formatNumber(item.yield, 2)}%</td>
                </tr>
              ))}
              {filtered.length === 0 && (
                <tr>
                  <td colSpan={7} className="muted" style={{textAlign:'center', padding:'24px'}}>
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

export { ObligasiSection }
