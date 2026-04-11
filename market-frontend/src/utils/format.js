// Indonesian number formatting (dot thousands)
function formatIDR(value, decimals = 0) {
  const fixed = value.toFixed(decimals).replace('.', ',')
  const parts = fixed.split(',')
  const intPart = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, '.')
  return `Rp ${intPart}${parts[1] ? ',' + parts[1] : ''}`
}

// US number formatting (comma thousands)
function formatUSD(value, decimals = 2) {
  return `$${value.toLocaleString('en-US', { minimumFractionDigits: decimals, maximumFractionDigits: decimals })}`
}

function formatPrice(value, currency = 'USD', decimals = 2) {
  return currency === 'IDR' ? formatIDR(value, decimals) : formatUSD(value, decimals)
}

function formatChange(value) {
  const sign = value >= 0 ? '▲' : '▼'
  const colorClass = value >= 0 ? 'positive' : 'negative'
  const val = Math.abs(value).toFixed(2)
  return `<span class="${colorClass}">${sign} ${val}%</span>`
}

function getChangeClass(value) {
  return value >= 0 ? 'positive' : 'negative'
}

function formatMarketCap(value) {
  const suffixes = ['', 'K', 'M', 'B', 'T']
  let i = 0
  while (value >= 1000 && i < suffixes.length - 1) {
    value /= 1000
    i++
  }
  return `$${value.toFixed(0)}${suffixes[i]}`
}

function formatVolume(value) {
  const suffixes = ['', 'K', 'M', 'B']
  let i = 0
  while (value >= 1000 && i < suffixes.length - 1) {
    value /= 1000
    i++
  }
  return value.toLocaleString('en-US', { maximumFractionDigits: 1 }) + suffixes[i]
}

function formatTimestamp(ts) {
  return new Date(ts * 1000).toLocaleDateString('id-ID', { 
    day: 'numeric', 
    month: 'short', 
    year: 'numeric' 
  })
}

function formatNumber(value, decimals = 2) {
  return value.toLocaleString('en-US', { 
    minimumFractionDigits: 0, 
    maximumFractionDigits: decimals 
  })
}

export {
  formatIDR,
  formatUSD,
  formatPrice,
  formatChange,
  getChangeClass,
  formatMarketCap,
  formatVolume,
  formatTimestamp,
  formatNumber
}

