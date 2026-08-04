# 🗺️ Go Yol Haritası — Gün Gün (8 Hafta)

**Varsayım:** Günde 1-2 saat, tutarlı çalışma. C/Java geçmişin var, giriş konularını atlıyoruz.
**Kural:** Her gün küçük de olsa **kod yaz**. Her fazı bir projeyle kapat. AI'a kod yazdırma — takıldığında "neden böyle" diye sor.

**İşaretleme:** `[ ]` → yapılacak, `[x]` → bitti. Her günün sonunda ilgili kutuyu doldur.

---

## ⚠️ Baştan Aklında Tut — 5 Klasik Tuzak
Bunları en baştan bilirsen çoğu hatadan kurtulursun:
1. **Error handling'i atlama.** Go'da hata bir değerdir, `if err != nil` yazmaktan kaçma. Dilin kalbi bu.
2. **Array değil slice kullan.** `[5]int` değil, neredeyse her zaman `[]int`.
3. **Pointer'ı anla** (C'den tanıdık ama Go otomatik dereference yapar, `->` yok).
4. **Gereksiz goroutine açma.** "Her şeyi paralel yapayım" tuzağına düşme; concurrency çözdüğü sorun kadar değerli.
5. **Idiomatic yaz.** Syntax bilmek yetmez; `gofmt`, `go vet` ve Effective Go senin pusulan.

---

## 📅 HAFTA 1 — Temeller I: Syntax & Toolchain
> Kaynak: [A Tour of Go](https://go.dev/tour/) + [Go by Example](https://gobyexample.com/)

- [ ] **Gün 1** — Toolchain: `go run`, `go build`, `go mod init`, `gofmt`. Hello world çalıştır. Tour: *Packages, imports, functions, variables*.
- [ ] **Gün 2** — Temel tipler: int/float/string/bool/rune, sabitler, `fmt` paketi (Printf/Sprintf). Tour: *Basic types, constants*.
- [ ] **Gün 3** — Kontrol akışı: `if`, `for` (Go'da tek döngü var!), `switch`, `defer`. Tour: *Flow control statements*.
- [ ] **Gün 4** — Struct & pointer: struct tanımı, alan erişimi, `&` ve `*`. Tour: *Pointers, structs*.
- [ ] **Gün 5** — 🔑 **Slice** (array değil!): `make`, `append`, `len`, `cap`, dilimleme. Tour: *Arrays, slices*.
- [ ] **Gün 6** — Map & range: `make(map[...]...)`, ekleme/silme, `range` ile dönme. Tour: *Maps, range*.
- [ ] **Gün 7** — Pekiştirme: [Effective Go](https://go.dev/doc/effective_go)'nun ilk yarısını oku. Mini alıştırma: bir metindeki **kelime frekansı sayacı** yaz (map + slice pratiği).

## 📅 HAFTA 2 — Temeller II: Metot, Interface, Error → PROJE 1
- [ ] **Gün 8** — Metotlar: receiver kavramı, value vs pointer receiver. Tour: *Methods*.
- [ ] **Gün 9** — 🔑 **Interface** (implicit! Java'dan en büyük fark), empty interface, type assertion/switch. Tour: *Interfaces*.
- [ ] **Gün 10** — 🔑 **Error handling**: error değerdir, `errors.New`, `fmt.Errorf` + `%w` ile wrap, `errors.Is`/`errors.As`. Go blog: *Error handling and Go*.
- [ ] **Gün 11** — `defer` derinlemesine, `panic`/`recover` (nadir kullanılır, ne zaman?).
- [ ] **Gün 12** — CLI girdisi: `os.Args`, `flag` paketi; `net` paketine ilk bakış (`net.DialTimeout`).
- [ ] **Gün 13-14** — 🚀 **PROJE 1: Port Tarayıcı**. Tek host + port aralığı al, TCP ile açık portları bul. Kodu **sen yaz**, ben mantığı anlatırım.

## 📅 HAFTA 3 — Standart Kütüphane & Test Kültürü → PROJE 2
- [ ] **Gün 15** — 🔑 **testing** paketi: table-driven test (Go idiomu). Port tarayıcıya test yaz. `go test ./...`.
- [ ] **Gün 16** — `encoding/json`: marshal/unmarshal, struct tag'leri.
- [ ] **Gün 17** — `net/http` **client**: GET isteği, `resp.Body` okuma, status code.
- [ ] **Gün 18** — Dosya I/O: `os.Open`, `bufio.Scanner` ile satır satır okuma.
- [ ] **Gün 19** — Düzgün CLI yapısı: `flag` (veya `cobra`), alt komutlar.
- [ ] **Gün 20-21** — 🚀 **PROJE 2: HTTP Prober**. Domain listesi oku → her birine istek at → canlı mı + status code raporla. (httpx'in minyatürü.)

## 📅 HAFTA 4 — Concurrency I: Temel Yapı Taşları
> Kaynak: [Go Blog — Concurrency Patterns](https://go.dev/blog/pipelines)
- [ ] **Gün 22** — `goroutine` nedir, `go` keyword, `sync.WaitGroup`.
- [ ] **Gün 23** — 🔑 **Channel**: unbuffered vs buffered, send/receive, `close`, `range` ile okuma.
- [ ] **Gün 24** — `select`, timeout pattern (`time.After`).
- [ ] **Gün 25** — 🔑 **context**: `WithTimeout`, `WithCancel` — iptal ve deadline.
- [ ] **Gün 26** — `sync.Mutex`, `sync.Once`; **race detector**: `go run -race`.
- [ ] **Gün 27-28** — Effective Go concurrency bölümü + küçük denemeler (channel ile üretici/tüketici).

## 📅 HAFTA 5 — Concurrency II: Pattern'ler → PROJE 3
- [ ] **Gün 29-30** — 🔑 **Worker pool** pattern (sınırlı goroutine ile iş kuyruğu).
- [ ] **Gün 31** — Fan-out / fan-in.
- [ ] **Gün 32-35** — 🚀 **PROJE 3: Paralel Tarayıcı**. Proje 2'yi worker pool ile paralelleştir, `context` ile iptal ekle, `-race` ile test et. **Hız farkını ölç** (seri vs paralel).

## 📅 HAFTA 6-7 — Web / API + Veritabanı → PROJE 4
> Kaynak: [Let's Go — Alex Edwards](https://lets-go.alexedwards.net/)
- [ ] **Gün 36-38** — `net/http` **server**: handler, routing, `http.ServeMux`, middleware mantığı.
- [ ] **Gün 39-41** — `database/sql` + PostgreSQL: bağlantı, sorgu, prepared statement, migration.
- [ ] **Gün 42-45** — 🔑 REST API tasarımı: JSON endpoint'ler, hata yanıtları, basit auth.
- [ ] **Gün 46-49** — 🚀 **PROJE 4: Tarama Sonuçları API + DB**. Tarama çıktılarını DB'ye yaz, sorgulanabilir REST API ile sun. (Supabase deneyimin burada işe yarar.)

## 📅 HAFTA 8 — Pekiştirme & Idiomatic Go
- [ ] **Gün 50-52** — [*100 Go Mistakes*](https://100go.co/) sitesinden seçmeler oku, kendi kodunda ara.
- [ ] **Gün 53-54** — Tüm projelerini `go vet` + `staticcheck` ile tara, refactor et.
- [ ] **Gün 55-56** — Bir aracını genişlet veya küçük bir açık kaynak issue'ya bak. Kendi Go stilini oturt.

---

## ✅ Faz Sonu Kontrol — "Öğrendim" Demek İçin
- [ ] `if err != nil` yazmak refleks oldu, hataları wrap'liyorum
- [ ] Slice/map/interface'i düşünmeden kullanıyorum
- [ ] Table-driven test yazabiliyorum
- [ ] goroutine + channel + context ile güvenli paralel kod yazabiliyorum
- [ ] `net/http` ile hem client hem server yazabiliyorum
- [ ] Başkasının Go kodunu açıp 30 saniyede ne yaptığını anlıyorum

**Hedef:** Rahat kod yazma ~2-3 ay · İş görür/idiomatic seviye ~4-6 ay.
