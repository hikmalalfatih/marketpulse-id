import { useState, useEffect } from 'react'
import './Header.css'

function Header() {
  const [marketStatus, setMarketStatus] = useState('Tutup')
  const [timeString, setTimeString] = useState('')
  const [dateString, setDateString] = useState('')

  useEffect(() => {
    const updateTime = () => {
      const now = new Date()
      const hours = now.getHours()
      const minutes = now.getMinutes()
      const dayNames = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']
      const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agu', 'Sep', 'Okt', 'Nov', 'Des']

      const dayName = dayNames[now.getDay()]
      const day = now.getDate()
      const month = monthNames[now.getMonth()]
      const year = now.getFullYear()
      const time = `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')} WIB`

      setDateString(`${dayName}, ${day} ${month} ${year} · ${time}`)

      // Market open: 09:00-16:00 Mon-Fri WIB
      const isWeekday = now.getDay() >= 1 && now.getDay() <= 5
      const isMarketHours = hours >= 9 && hours < 16
      setMarketStatus(isWeekday && isMarketHours ? 'Buka' : 'Tutup')
    }

    updateTime()
    const interval = setInterval(updateTime, 1000)
    return () => clearInterval(interval)
  }, [])

  const statusClass = marketStatus === 'Buka' ? 'market-open' : 'market-closed'

  return (
    <header className="header">
      <a className="logo" href='http://maley.vercel.app'>
        ◈ about dev
      </a>
      <div className={`market-status ${statusClass}`}>
        🟢 Pasar {marketStatus}
      </div>
      <div className="clock">
        {dateString}
      </div>
    </header>
  )
}

export { Header }

