import './Sections.css'
import { Skeleton } from '../ui/Skeleton.jsx'
import { formatIDR, formatNumber } from '../utils/format.js'

const CURRENCIES = [
  { code: 'IDR', name: 'Rupiah Indonesia',     flag: '🇮🇩', invert: true  },
  { code: 'EUR', name: 'Euro',                  flag: '🇪🇺', invert: false },
  { code: 'GBP', name: 'Poundsterling',         flag: '🇬🇧', invert: false },
  { code: 'SGD', name: 'Dollar Singapura',      flag: '🇸🇬', invert: false },
  { code: 'MYR', name: 'Ringgit Malaysia',      flag: '🇲🇾', invert: false },
  { code: 'JPY', name: 'Yen Jepang',            flag: '🇯🇵', invert: false },
  { code: 'CNY', name: 'Yuan China',            flag: '🇨🇳', invert: false },
  { code: 'AUD', name: 'Dollar Australia',      flag: '🇦🇺', invert: false },
  { code: 'SAR', name: 'Riyal Arab Saudi',      flag: '🇸🇦', invert: false },
]

function ForexSection({ data, loading }) {
  // data = { from: "USD", rates: { IDR: 16384, EUR: 0.92, ... }, updated_at: "..." }
  const rates = data?.rates || {}
  const idrRate = rates['IDR'] || 0

  return (
    <div className="section">
      <div className="section-header">
        <div>
          <div className="section-title"><span className="section-icon">💱</span>Kurs Valuta Asing</div>
          <div className="section-subtitle">Nilai tukar terhadap USD · {data?.updated_at ? new Date(data.updated_at).toLocaleTimeString('id-ID') : ''}</div>
        </div>
      </div>
      <div className="section-content">
        {loading ? <Skeleton rows={5} /> : (
          <>
            {/* USD/IDR highlight */}
            {idrRate > 0 && (
              <div className="forex-highlight">
                <span className="forex-highlight-label">🇺🇸 1 USD =</span>
                <span className="forex-highlight-value num">{formatIDR(idrRate)}</span>
              </div>
            )}

            <div className="forex-grid">
              {CURRENCIES.filter(c => c.code !== 'IDR').map(curr => {
                const rate = rates[curr.code]
                if (!rate) return null

                // Convert: 1 unit of curr.code = ? IDR
                const toIDR = idrRate / rate

                return (
                  <div key={curr.code} className="forex-card">
                    <div className="forex-card-header">
                      <span className="forex-flag">{curr.flag}</span>
                      <div>
                        <div className="forex-code">{curr.code}</div>
                        <div className="forex-name">{curr.name}</div>
                      </div>
                    </div>
                    <div className="forex-rate-row">
                      <span className="forex-label">1 {curr.code}</span>
                      <span className="forex-value num">
                        {idrRate > 0 ? formatIDR(toIDR) : `${formatNumber(rate, 4)} USD`}
                      </span>
                    </div>
                    <div className="forex-rate-row muted">
                      <span className="forex-label">1 USD</span>
                      <span className="forex-value num">{formatNumber(rate, curr.code === 'JPY' ? 2 : 4)} {curr.code}</span>
                    </div>
                  </div>
                )
              })}
            </div>
          </>
        )}
      </div>
    </div>
  )
}

export { ForexSection }
