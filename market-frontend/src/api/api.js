const apiOrigin = import.meta.env.VITE_API_ORIGIN?.replace(/\/$/, '') || ''
const BASE = `${apiOrigin}/api`

async function get(path) {
  try {
    const response = await fetch(BASE + path)
    if (!response.ok) {
      throw new Error(`HTTP ${response.status}`)
    }
    return await response.json()
  } catch (err) {
    console.error('API error:', err)
    return { success: false, error: err.message }
  }
}

const api = {
  crypto: () => get('/crypto'),
  stocksID: () => get('/stocks/id'),
  stocksUS: () => get('/stocks/us'),
  gold: () => get('/gold'),
  oil: () => get('/oil'),
  forex: () => get('/forex'),
  reksadana: () => get('/reksadana'),
  obligasi: () => get('/obligasi'),
  summary: () => get('/summary'),
  history: (type, symbol) => get(`/history/${type}/${symbol}`)
}

export { api }
export default api

