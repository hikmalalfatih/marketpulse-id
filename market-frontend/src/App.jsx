import { useState } from 'react'
import './App.css'
import { Header } from './layout/Header.jsx'
import { Sidebar } from './layout/Sidebar.jsx'
import { TickerBar } from './layout/TickerBar.jsx'
import { Footer } from './layout/Footer.jsx'
import { useMarketData } from './hooks/useMarketData.js'
import { SummarySection } from './sections/SummarySection.jsx'
import { CryptoSection } from './sections/CryptoSection.jsx'
import { StocksIDSection } from './sections/StocksIDSection.jsx'
import { StocksUSSection } from './sections/StocksUSSection.jsx'
import { CommoditySection } from './sections/CommoditySection.jsx'
import { ForexSection } from './sections/ForexSection.jsx'
import { ReksadanaSection } from './sections/ReksadanaSection.jsx'
import { ObligasiSection } from './sections/ObligasiSection.jsx'
import { ErrorBanner } from './ui/ErrorBanner.jsx'

const NAV_ITEMS = [
  { id: 'summary',   icon: '📊', label: 'Ringkasan' },
  { id: 'crypto',    icon: '🪙', label: 'Kripto' },
  { id: 'stocks-id', icon: '🇮🇩', label: 'Saham IDX' },
  { id: 'stocks-us', icon: '🇺🇸', label: 'Saham US' },
  { id: 'commodity', icon: '⚡', label: 'Emas & Perak' },
  { id: 'oil',       icon: '🛢️', label: 'Minyak & Gas' },
  { id: 'forex',     icon: '💱', label: 'Kurs Valas' },
  { id: 'reksadana', icon: '📈', label: 'Reksadana' },
  { id: 'obligasi',  icon: '📋', label: 'Obligasi' },
]

function App() {
  const [activeSection, setActiveSection] = useState('summary')

  const {
    crypto, stocksID, stocksUS, gold, oil, forex,
    reksadana, obligasi, loading, error, refresh
  } = useMarketData()

  const renderSection = () => {
    switch (activeSection) {
      case 'summary':
        return <SummarySection
          crypto={crypto} stocksID={stocksID} forex={forex}
          gold={gold} oil={oil} loading={loading}
          onNavigate={setActiveSection}
        />
      case 'crypto':
        return <CryptoSection data={crypto} loading={loading} />
      case 'stocks-id':
        return <StocksIDSection data={stocksID} loading={loading} />
      case 'stocks-us':
        return <StocksUSSection data={stocksUS} loading={loading} />
      case 'commodity':
        return <CommoditySection gold={gold} oil={oil} loading={loading} />
      case 'oil':
        return <CommoditySection gold={gold} oil={oil} loading={loading} oilOnly />
      case 'forex':
        return <ForexSection data={forex} loading={loading} />
      case 'reksadana':
        return <ReksadanaSection data={reksadana} loading={loading} />
      case 'obligasi':
        return <ObligasiSection data={obligasi} loading={loading} />
      default:
        return null
    }
  }

  return (
    <div className="app">
      <Header />
      <TickerBar crypto={crypto} forex={forex} />
      <div className="main-layout">
        <Sidebar
          items={NAV_ITEMS}
          active={activeSection}
          onSelect={setActiveSection}
        />
        <main className="main-content">
          {error && <ErrorBanner message={error} onRetry={refresh} />}
          {renderSection()}
        </main>
      </div>
      <Footer />
    </div>
  )
}

export default App
