# Fullstack Blog-App (Go + Next.js)

## Backend Dependencies

```bash
# Tambahkan dependensi utama (pin versi stabil terbaru per 11 Agu 2025)
go get github.com/go-chi/chi/v5@v5.2.2                # router HTTP
go get github.com/go-chi/cors@v1.2.2                  # CORS middleware
go get gorm.io/gorm@v1.30.1                           # ORM
go get gorm.io/driver/mysql@v1.5.7                    # contoh versi terbaru (cek halaman pkg) :contentReference[oaicite:18]{index=18}
go get github.com/spf13/viper@v1.20.1                 # config
go get go.uber.org/zap@v1.27.0                        # logging
go get github.com/golang-jwt/jwt/v5@v5.2.1            # JWT
go get github.com/go-playground/validator/v10@v10.27.0# validasi
go get github.com/google/uuid@v1.6.0                  # UUID
go get github.com/casbin/casbin/v2@v2.113.0           # RBAC Casbin
go get github.com/glebarez/sqlite@v1.11.0             # opsional: driver SQLite murni Go
go get github.com/stretchr/testify@v1.10.0            # testing

# Tooling (opsional): live reload & migrasi (pakai binary saat run)
go install github.com/air-verse/air@v1.62.0           # air (PATH harus mengarah ke %USERPROFILE%\go\bin) :contentReference[oaicite:19]{index=19}
go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2  # CLI migrasi
```

## Install Frontend (Nextjs) pada folder web

```bash
npm create next-app@latest . -- --ts --tailwind --eslint --app --turbopack
```

## Frontend Dependencies

```bash
# TanStack Query v5 + Devtools
npm i @tanstack/react-query@5 @tanstack/query-devtools@5
# shadcn/ui + dark mode helper
npx shadcn@latest init
npm i next-themes
```
