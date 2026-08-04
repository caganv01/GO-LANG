# GO-LANG — Go Öğrenme Yolculuğum 🐹

Proje odaklı, faz faz Go öğrenme deposu. Her fazı bir projeyle kapatıyorum ve kodu kendim yazıyorum.
Odak: güvenlik / pentest tooling (nuclei, ffuf, httpx gibi araçların yazıldığı dil).

## 📚 Kaynaklar
- [A Tour of Go](https://go.dev/tour/) — interaktif temel tur (buradan başla)
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic Go (en kritik doküman)
- [Go by Example](https://gobyexample.com/) — çalıştırılabilir örnekler (referans)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) — TDD ile Go
- [The Go Blog](https://go.dev/blog/) — concurrency & error handling yazıları

## 🗺️ Roadmap & İlerleme

| Faz | Konu | Proje | Durum |
|-----|------|-------|-------|
| 0 | Kurulum & toolchain | — | 🔄 Devam ediyor |
| 1 | Temeller (struct, slice, interface, error, defer) | Port tarayıcı (CLI) | ⬜ |
| 2 | Stdlib, CLI & test (net/http, os, flag, json, testing) | HTTP prober / subdomain enumerator | ⬜ |
| 3 | Concurrency (goroutine, channel, select, context) | Paralel tarayıcı (worker pool) | ⬜ |
| 4 | Web/API + DB (net/http, database/sql, PostgreSQL) | Tarama sonuçları API + DB | ⬜ |

**Durum işaretleri:** ⬜ başlanmadı · 🔄 devam · ✅ bitti

## 📂 Yapı
```
GO-LANG/
├── notlar/        → faz faz markdown notlar (öğrendiklerim)
└── projeler/      → her proje kendi Go modülü
```

## ⏱️ Hedef
Günde 1-2 saat: rahat kod yazma ~2-3 ay, iş görür seviye ~4-6 ay.

## Başlangıç: 2026-08-04
