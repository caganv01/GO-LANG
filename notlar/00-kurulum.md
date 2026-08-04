# Faz 0 — Kurulum & Toolchain

## Go'yu kurma (Windows)
1. https://go.dev/dl/ adresinden **Windows MSI** installer'ı indir (`goX.XX.X.windows-amd64.msi`).
2. Çift tıkla, kur. Installer PATH'i otomatik ayarlar.
3. **Yeni bir terminal aç** (PATH'in yenilenmesi için) ve doğrula:
   ```
   go version
   ```

> Not: GOPATH ile uğraşma, üçüncü parti "version manager" kurma. Go 1.21+ toolchain
> versiyon seçimini kendi hallediyor. Sade tut.

## Öğrenilecek temel komutlar
| Komut | Ne yapar |
|-------|----------|
| `go run main.go` | Derleyip anında çalıştırır (dev sırasında) |
| `go build` | Binary (.exe) üretir |
| `go mod init <isim>` | Yeni modül başlatır (go.mod dosyası) |
| `go mod tidy` | Bağımlılıkları temizler/ekler |
| `go test ./...` | Tüm testleri çalıştırır |
| `gofmt -w .` | Kodu Go standardına göre formatlar |
| `go vet ./...` | Statik hata avı |

## İlk çalışan program (hello)
```go
package main

import "fmt"

func main() {
    fmt.Println("Merhaba Go")
}
```
Çalıştır: `go run main.go`

## Durum
- [ ] Go kuruldu (`go version` çalışıyor)
- [ ] İlk `hello` programı çalıştı
- [ ] `go`, `git` komutları terminalden erişilebilir
