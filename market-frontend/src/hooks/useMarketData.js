import { useState, useEffect, useCallback, useRef } from 'react'
import { api } from '../api/api.js'

const REFRESH_INTERVAL = 60 * 1000 // 60s
const STALE_THRESHOLD = 5 * 60 * 1000 // 5 min

export function useMarketData() {
  const [state, setState] = useState({
    crypto: [],
    stocksID: [],
    stocksUS: [],
    gold: {},
    oil: [],
    forex: {},
    reksadana: [],
    obligasi: [],
    loading: true,
    error: null,
    lastUpdated: null,
  })

  const timerRef = useRef(null)
  const lastFetchRef = useRef(0)

  const fetchAll = useCallback(async (silent = false) => {
    if (!silent) {
      setState(prev => ({ ...prev, loading: true, error: null }))
    }

    try {
      const [cryptoRes, stocksIDRes, stocksUSRes, goldRes, oilRes, forexRes, reksadanaRes, obligasiRes] =
        await Promise.all([
          api.crypto(),
          api.stocksID(),
          api.stocksUS(),
          api.gold(),
          api.oil(),
          api.forex(),
          api.reksadana(),
          api.obligasi(),
        ])

      lastFetchRef.current = Date.now()

      setState({
        crypto:    cryptoRes.success    ? (cryptoRes.data    || []) : [],
        stocksID:  stocksIDRes.success  ? (stocksIDRes.data  || []) : [],
        stocksUS:  stocksUSRes.success  ? (stocksUSRes.data  || []) : [],
        gold:      goldRes.success      ? (goldRes.data      || {}) : {},
        oil:       oilRes.success       ? (oilRes.data       || []) : [],
        forex:     forexRes.success     ? (forexRes.data     || {}) : {},
        reksadana: reksadanaRes.success ? (reksadanaRes.data || []) : [],
        obligasi:  obligasiRes.success  ? (obligasiRes.data  || []) : [],
        loading:   false,
        error:     null,
        lastUpdated: new Date().toISOString(),
      })
    } catch (err) {
      setState(prev => ({
        ...prev,
        loading: false,
        error: `Gagal memuat data: ${err.message}`,
      }))
    }
  }, [])

  // Initial fetch
  useEffect(() => {
    fetchAll(false)
  }, [fetchAll])

  // Auto-refresh every 60s
  useEffect(() => {
    timerRef.current = setInterval(() => {
      fetchAll(true)
    }, REFRESH_INTERVAL)
    return () => clearInterval(timerRef.current)
  }, [fetchAll])

  // Refresh on tab focus if data is stale
  useEffect(() => {
    const onFocus = () => {
      if (Date.now() - lastFetchRef.current > STALE_THRESHOLD) {
        fetchAll(true)
      }
    }
    window.addEventListener('focus', onFocus)
    return () => window.removeEventListener('focus', onFocus)
  }, [fetchAll])

  return { ...state, refresh: () => fetchAll(false) }
}
