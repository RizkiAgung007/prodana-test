🚀 Prodana System - Management Dashboard with Gemini AI
Prodana System adalah aplikasi manajemen data internal yang mengintegrasikan sistem CRUD (Create, Read, Update, Delete) untuk User dan Produk. Aplikasi ini diperkuat dengan fitur Artificial Intelligence (Gemini AI) untuk pembuatan deskripsi produk secara otomatis, profesional, dan persuasif.

Dirancang dengan arsitektur modern menggunakan Golang pada sisi backend dan React pada sisi frontend untuk menjamin performa tinggi serta antarmuka yang responsif.

✨ Fitur Utama
- 🔐 Authentication & Authorization: Login menggunakan JWT (JSON Web Token) dengan 3 level  akses: Admin, Editor, dan Viewer.
- 👥 User Management: Manajemen data pengguna (Nama, Email, Role) dengan otentikasi password yang aman (Bcrypt).
- 📦 Product Management: Katalog produk lengkap dengan Harga, Stok, dan Deskripsi.
- ✨ AI Product Describer: Integrasi Google Gemini AI untuk men-generate deskripsi produk yang menarik dan profesional hanya berdasarkan nama produk.
- 📱 Responsive & Modern UI: Antarmuka bersih menggunakan Tailwind CSS dengan fitur Sticky Footer dan Fixed Header.

🛠️ Tech Stack
1. Backend (Golang)
    - Framework: net/http (Standard Library)
    - ORM: GORM
    - Database: PostgreSQL / SQLite
    - Security: Bcrypt & JWT
    - AI Integration: Google Generative AI (Gemini 1.5 Flash) via REST API

2. Frontend (React)
    - Framework: React.js (Vite)
    - Routing: React Router DOM v6
    - Styling: Tailwind CSS
    - HTTP Client: Axios

⚙️ Instalasi & Konfigurasi
1. Prasyarat
    - Go (Minimal v1.21.4)
    - Node.js & npm
    - Google Gemini API Key

2. Setup Backend
    - Masuk ke direktori backend
        cd backend

    - Install dependencies
        go mod tidy

    - Set Environment Variable (Windows PowerShell)
        $env:APIKEY="YOUR_GEMINI_API_KEY"

    - Jalankan server
        go run main.go

3. Setup Frontend
    - Masuk ke direktori frontend
        cd frontend

    - Install dependencies
        npm install

    - Jalankan aplikasi
        npm run dev

📂 Struktur Proyek
PRODANA-TEST/
├── backend/
│   ├── cmd/                # Entry point aplikasi
│   ├── internal/
│   │   ├── config/         # Konfigurasi Database & Env
│   │   ├── handlers/       # Logika API (termasuk AI Handler)
│   │   ├── middleware/     # Auth & CORS Middleware
│   │   ├── models/         # Schema Database
│   │   └── utils/          # Helper (JWT, Password)
│   └── main.go
└── frontend/
    ├── src/
    │   ├── components/     # UI Components (Header, Footer, ProtectedRoute)
    │   ├── Auth/          # Data Login (hooks, pages)
    │   ├── Dashboard/          # Data User (components, hooks, pages)
    │   ├── Product/          # Data Product (components, hooks, pages)
    │   └── services/       # API Services (api, utils)

Fitur       Admin   Editor  Viewer
Lihat Data   ✅      ✅     ✅
Tambah Data  ✅      ❌     ❌
Edit Data    ✅      ✅     ❌
Hapus Data   ✅      ✅     ❌
Generate AI  ✅      ❌     ❌

👤 Developer
Rizki Agung Dermawan

Catatan Penting
- Pastikan backend berjalan di port :8080.
- Jika menggunakan database lokal, pastikan kredensial di internal/config/database.go sudah sesuai.
- Untuk fitur AI, pastikan koneksi internet stabil untuk menghubungi endpoint Google API.
