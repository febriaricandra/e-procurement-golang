# Variabel
APP_NAME=my-procurement-system
APP_PORT=8080
DB_NAME=procurement_db
DB_USER=postgres
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=5432

# Target default
.DEFAULT_GOAL := help

# Help: Menampilkan semua perintah yang tersedia
help:
	@echo "Gunakan: make <target>"
	@echo ""
	@echo "Target:"
	@echo "  build         : Build aplikasi"
	@echo "  run           : Jalankan aplikasi"
	@echo "  test          : Jalankan semua unit test"
	@echo "  migrate       : Jalankan migrasi database"
	@echo "  clean         : Hapus file binary dan cache"
	@echo "  docker-build  : Build Docker image"
	@echo "  docker-run    : Jalankan aplikasi di Docker"
	@echo "  help          : Tampilkan pesan bantuan ini"

# Build: Build aplikasi
build:
	@echo "🚀 Membangun aplikasi..."
	go build -o bin/$(APP_NAME) cmd/api/main.go
	@echo "✅ Aplikasi berhasil dibangun: bin/$(APP_NAME)"

# Run: Jalankan aplikasi
run:
	@echo "🚀 Menjalankan aplikasi..."
	APP_PORT=$(APP_PORT) DB_NAME=$(DB_NAME) DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) go run ./cmd/api/main.go

# Test: Jalankan semua unit test
test:
	@echo "🧪 Menjalankan unit test..."
	go test -v ./...

# Migrate: Jalankan migrasi database
migrate:
	@echo "🔄 Menjalankan migrasi database..."
	go run ./cmd/api/main.go -migrate

# Clean: Hapus file binary dan cache
clean:
	@echo "🧹 Membersihkan file binary dan cache..."
	rm -rf bin/$(APP_NAME)
	go clean -testcache
	@echo "✅ Pembersihan selesai."