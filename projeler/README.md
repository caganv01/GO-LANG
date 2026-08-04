# Projeler

Her proje kendi Go modülüdür (kendi `go.mod` dosyası). Faz sırasına göre:

| # | Proje | Faz | Ne öğretir |
|---|-------|-----|------------|
| 01 | port-tarayici | 1 | net paketi, CLI arg, temel tipler |
| 02 | http-prober | 2 | net/http, json, dosya I/O, test |
| 03 | paralel-tarayici | 3 | goroutine, channel, worker pool, context |
| 04 | tarama-api | 4 | REST API, database/sql, PostgreSQL |

Yeni proje başlatma:
```
cd projeler/01-port-tarayici
go mod init port-tarayici
```
