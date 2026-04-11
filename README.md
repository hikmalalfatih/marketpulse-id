<div align="center">

# ◈ MarketPulse ID

### Dashboard Pasar Keuangan Indonesia Real-Time

**Pantau crypto, saham IDX & Wall Street, emas, minyak, kurs valas, reksadana, dan obligasi — semua dalam satu tampilan.**

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![React](https://img.shields.io/badge/React-18-61DAFB?style=flat-square&logo=react&logoColor=black)](https://reactjs.org)
[![Vite](https://img.shields.io/badge/Vite-5-646CFF?style=flat-square&logo=vite&logoColor=white)](https://vitejs.dev)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-Codeby--Javier-181717?style=flat-square&logo=github)](https://github.com/Codeby-Javier)

![MarketPulse ID Preview](https://raw.githubusercontent.com/Codeby-Javier/marketpulse-id/main/preview.png)

</div>

---

## 🎯 Tentang Proyek

**MarketPulse ID** adalah aplikasi dashboard keuangan real-time yang dirancang khusus untuk investor dan trader Indonesia. Proyek ini hadir untuk menjawab satu masalah nyata: **data pasar keuangan Indonesia tersebar di banyak platform berbeda**, memaksa pengguna membuka puluhan tab hanya untuk memantau portofolio mereka.

Dengan MarketPulse ID, semua data tersedia dalam **satu dashboard terpadu** — dari harga Bitcoin hingga yield obligasi SBN, dari kurs USD/IDR hingga NAB reksadana terpopuler.

### Masalah yang Dipecahkan

| Masalah | Solusi MarketPulse ID |
|---|---|
| Data crypto, saham, dan komoditas tersebar di banyak platform | Satu dashboard terpadu dengan 8 kategori aset |
| Tidak ada platform gratis yang menggabungkan data IDX + crypto + obligasi | Agregasi data dari 5+ sumber publik secara otomatis |
| Platform finansial Indonesia umumnya lambat dan berat | Backend Go dengan in-memory cache, respons < 100ms |
| Biaya langganan platform data finansial mahal | 100% gratis, open source, zero API key berbayar |
| Tidak ada konteks harga dalam Rupiah untuk aset global | Semua harga ditampilkan dalam USD dan IDR secara bersamaan |

---

## ✨ Fitur Utama

- **📊 Ringkasan Pasar** — Snapshot semua aset dalam satu halaman, top gainers & losers crypto
- **🪙 Kripto** — Top 20 aset berdasarkan market cap, sparkline 7 hari, data 1j/24j/7h
- **🇮🇩 Saham IDX** — 15 blue chip BEI (BBCA, BBRI, BMRI, TLKM, ASII, dll) dengan harga IDR real-time
- **🇺🇸 Saham US** — 10 saham Wall Street (AAPL, MSFT, NVDA, TSLA, META, dll)
- **⚡ Emas & Perak** — Harga XAU/XAG per troy oz dan per gram dalam USD & IDR
- **🛢️ Minyak & Gas** — Brent Crude, WTI Crude, Natural Gas
- **💱 Kurs Valas** — USD ke IDR, EUR, GBP, SGD, JPY, CNY, AUD, MYR, SAR
- **📈 Reksadana** — Performa NAB/unit, return 1 bulan, YTD, 1 tahun
- **📋 Obligasi** — SBN, ORI, SR, PBS dengan kupon, yield, dan harga pasar
- **📉 Grafik Historis** — Chart 90 hari interaktif untuk setiap aset (klik baris untuk membuka)
- **🔄 Auto-refresh** — Data diperbarui otomatis setiap 60 detik tanpa reload halaman
- **📡 Ticker Bar** — Scrolling ticker real-time di bagian atas dengan harga crypto terkini

---

## 🏗️ Tech Stack

### Backend — Go 1.24

```
market-api/
├── main.go              # Entry point, routing, graceful shutdown
├── config/config.go     # Konfigurasi tickers, port, timeout
├── cache/cache.go       # In-memory cache dengan sync.RWMutex
├── models/models.go     # Struct definisi semua tipe data
├── middleware/cors.go   # CORS middleware
├── handlers/            # HTTP handler per endpoint
│   ├── crypto.go
│   ├── stocks.go
│   ├── gold.go
│   ├── oil.go
│   ├── forex.go
│   ├── reksadana.go
│   ├── obligasi.go
│   ├── history.go
│   ├── summary.go
│   └── response.go
└── scraper/             # Data fetcher per sumber
    ├── client.go        # Shared HTTP client
    ├── crypto.go        # CoinGecko API
    ├── stocks_id.go     # Yahoo Finance (IDX)
    ├── stocks_us.go     # Yahoo Finance (US)
    ├── gold.go          # GoldPrice.org
    ├── oil.go           # Yahoo Finance (futures)
    ├── forex.go         # ExchangeRate API
    ├── reksadana.go     # Kontan (HTML scraping)
    └── obligasi.go      # DJPPR Kemenkeu
```

**Keunggulan arsitektur backend:**

- **Zero external dependencies** — hanya standard library Go (`net/http`, `encoding/json`, `sync`, `time`, dll)
- **In-memory cache** dengan TTL berbeda per endpoint (3 menit crypto, 30 menit forex, 60 menit reksadana)
- **Parallel fetching** di endpoint `/api/summary` menggunakan goroutines + `sync.WaitGroup`
- **Graceful shutdown** dengan `os/signal` — tidak ada request yang terpotong saat restart
- **Fallback data** — jika semua sumber eksternal gagal, API tetap mengembalikan data referensi
- **Stale cache** — jika fetch gagal tapi cache ada, data lama dikembalikan dengan flag `stale: true`

### Frontend — React 18 + Vite 5

```
market-frontend/src/
├── App.jsx              # Root component, routing antar section
├── api/api.js           # HTTP client wrapper
├── hooks/
│   ├── useMarketData.js # Fetch semua data paralel, auto-refresh 60s
│   └── useHistory.js    # Fetch data historis per aset
├── utils/format.js      # Formatter IDR, USD, market cap, volume
├── layout/
│   ├── Header.jsx       # Status pasar + jam WIB real-time
│   ├── Sidebar.jsx      # Navigasi dengan active state
│   ├── TickerBar.jsx    # Scrolling ticker dari data live
│   └── Footer.jsx       # Link sosial media
├── sections/            # Halaman per kategori aset
│   ├── SummarySection.jsx
│   ├── CryptoSection.jsx
│   ├── StocksIDSection.jsx
│   ├── StocksUSSection.jsx
│   ├── CommoditySection.jsx
│   ├── ForexSection.jsx
│   ├── ReksadanaSection.jsx
│   └── ObligasiSection.jsx
├── charts/
│   ├── MiniSparkline.jsx  # Sparkline 7 hari (Recharts)
│   └── PriceChart.jsx     # Chart historis 90 hari (Recharts)
└── ui/
    ├── AssetModal.jsx   # Modal chart historis on-click
    ├── Badge.jsx        # Badge perubahan harga +/- berwarna
    ├── Skeleton.jsx     # Loading skeleton animasi
    └── ErrorBanner.jsx  # Error state dengan tombol retry
```

**Keunggulan arsitektur frontend:**

- **Plain CSS only** — tidak ada Tailwind, MUI, atau Chakra. CSS variables untuk design system konsisten
- **Recharts** sebagai satu-satunya library eksternal untuk visualisasi data
- **Tab focus refresh** — data otomatis diperbarui saat pengguna kembali ke tab setelah 5 menit
- **Click-to-chart** — klik baris mana saja untuk membuka modal grafik historis 90 hari
- **Monospace numbers** — semua angka menggunakan `font-variant-numeric: tabular-nums` untuk alignment sempurna
- **Responsive** — sidebar collapse di tablet, bottom-friendly di mobile

### Data Sources

| Kategori | Sumber | Cache TTL |
|---|---|---|
| Crypto (Top 20) | CoinGecko Public API | 3 menit |
| Saham IDX (15 ticker) | Yahoo Finance v8 | 5 menit |
| Saham US (10 ticker) | Yahoo Finance v8 | 5 menit |
| Emas & Perak | GoldPrice.org | 10 menit |
| Minyak (Brent/WTI/Gas) | Yahoo Finance Futures | 5 menit |
| Kurs Valas | ExchangeRate API | 30 menit |
| Reksadana | Kontan (HTML scraping) | 60 menit |
| Obligasi SBN | DJPPR Kemenkeu | 60 menit |
| Historis (semua aset) | CoinGecko + Yahoo Finance | 15 menit |

---

## 🚀 Cara Menjalankan

### Prasyarat

- [Go 1.24+](https://golang.org/dl/)
- [Node.js 18+](https://nodejs.org/)

### 1. Clone Repository

```bash
git clone https://github.com/Codeby-Javier/marketpulse-id.git
cd marketpulse-id
```

### 2. Jalankan Backend

```bash
cd market-api
go run .
```

Backend akan berjalan di `http://localhost:8080`

```
┌─────────────────────────────────────┐
│  MarketPulse ID Backend v1.0        │
│  Listening on :8080                 │
│  Ready to serve financial data      │
└─────────────────────────────────────┘
```

### 3. Jalankan Frontend

Buka terminal baru:

```bash
cd market-frontend
npm install
npm run dev
```

Frontend akan berjalan di `http://localhost:3000`

---

## 📡 API Endpoints

| Method | Endpoint | Deskripsi | Cache |
|---|---|---|---|
| GET | `/health` | Health check | - |
| GET | `/api/crypto` | Top 20 crypto by market cap | 3 menit |
| GET | `/api/stocks/id` | 15 saham IDX blue chip | 5 menit |
| GET | `/api/stocks/us` | 10 saham Wall Street | 5 menit |
| GET | `/api/gold` | Harga emas & perak | 10 menit |
| GET | `/api/oil` | Brent, WTI, Natural Gas | 5 menit |
| GET | `/api/forex` | Kurs USD ke 9 mata uang | 30 menit |
| GET | `/api/reksadana` | Daftar reksadana populer | 60 menit |
| GET | `/api/obligasi` | SBN, ORI, SR, PBS | 60 menit |
| GET | `/api/history/{type}/{symbol}` | Data historis 90 hari | 15 menit |
| GET | `/api/summary` | Semua data dalam satu respons | 5 menit |

**Contoh respons:**

```json
{
  "success": true,
  "data": [...],
  "cached": false,
  "source": "coingecko",
  "updated_at": "2025-04-11T14:30:00Z"
}
```

**History endpoint contoh:**
- `/api/history/crypto/bitcoin`
- `/api/history/stock_id/BBCA`
- `/api/history/stock_us/AAPL`
- `/api/history/gold/gold`
- `/api/history/oil/brent`

---

## 🎨 Design System

Dashboard menggunakan dark theme dengan prinsip **data-dense, no-fluff**:

```css
--bg-primary:  #0b0f14   /* Background utama */
--bg-card:     #141920   /* Card/panel */
--bg-header:   #0d1117   /* Header & sidebar */
--green:       #0ecb81   /* Harga naik */
--red:         #f6465d   /* Harga turun */
--gold:        #f0b90b   /* Aksen emas */
--accent:      #1890ff   /* Navigasi aktif */
--font-mono:   'Courier New', Consolas  /* Semua angka */
```

Terinspirasi dari tampilan terminal trading profesional — tidak ada gradien berlebihan, tidak ada glassmorphism, tidak ada animasi yang mengganggu fokus.

---

## 📁 Struktur Proyek

```
marketpulse-id/
├── market-api/          # Go backend
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   ├── cache/
│   ├── models/
│   ├── middleware/
│   ├── handlers/
│   ├── scraper/
│   └── utils/
└── market-frontend/     # React frontend
    ├── index.html
    ├── vite.config.js
    ├── package.json
    └── src/
        ├── api/
        ├── hooks/
        ├── utils/
        ├── layout/
        ├── sections/
        ├── charts/
        └── ui/
```

---

## 🤝 Kontribusi

Pull request sangat disambut. Untuk perubahan besar, buka issue terlebih dahulu untuk mendiskusikan apa yang ingin diubah.

---

## 👤 Author

**Javier** — [@ibnu.jz](https://instagram.com/ibnu.jz)

- Instagram: [@ibnu.jz](https://instagram.com/ibnu.jz)
- GitHub: [@Codeby-Javier](https://github.com/Codeby-Javier)

---

## 📄 Lisensi

[MIT](LICENSE) © 2025 Codeby-Javier

---

<div align="center">
  <sub>Dibuat dengan ☕ dan Go + React · Data untuk tujuan informasi, bukan saran investasi</sub>
</div>
