import { useEffect } from 'react'
import { useHistory } from '../hooks/useHistory.js'
import { PriceChart } from '../charts/PriceChart.jsx'
import { Skeleton } from './Skeleton.jsx'
import './UI.css'

function AssetModal({ type, symbol, title, color, onClose }) {
  const { data, loading, error } = useHistory({ type, symbol, enabled: true })

  // Close on Escape
  useEffect(() => {
    const handler = (e) => { if (e.key === 'Escape') onClose() }
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [onClose])

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-panel" onClick={e => e.stopPropagation()}>
        <div className="modal-header">
          <div className="modal-title">{title}</div>
          <button className="modal-close" onClick={onClose}>✕</button>
        </div>
        <div className="modal-body">
          {loading && <Skeleton rows={4} />}
          {error && <div className="muted" style={{padding:'16px'}}>Gagal memuat data historis: {error}</div>}
          {!loading && !error && data.length > 0 && (
            <PriceChart data={data} title="" color={color} />
          )}
          {!loading && !error && data.length === 0 && (
            <div className="muted" style={{padding:'16px', textAlign:'center'}}>Tidak ada data historis</div>
          )}
        </div>
      </div>
    </div>
  )
}

export { AssetModal }
