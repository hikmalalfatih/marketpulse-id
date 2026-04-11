import { useState, useEffect } from 'react'
import api from '../api/api.js'

export function useHistory({ type, symbol, enabled = true }) {
  const [state, setState] = useState({
    data: [],
    loading: false,
    error: null
  })

  useEffect(() => {
    if (!enabled || !type || !symbol) return

    setState(prev => ({ ...prev, loading: true, error: null }))

    api.history(type, symbol).then(res => {
      if (res.success) {
        setState({ data: res.data || [], loading: false, error: null })
      } else {
        setState({ data: [], loading: false, error: res.error })
      }
    }).catch(err => {
      setState({ data: [], loading: false, error: err.message })
    })

  }, [type, symbol, enabled])

  return state
}

