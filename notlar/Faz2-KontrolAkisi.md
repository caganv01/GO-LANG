# 📚 Faz 2 — Kontrol Akışı (for, if-else, switch, defer)

> Gün 8-14 — Döngüler, koşullar, seçim yapısı, erteleme.

---

## 1️⃣ FOR — Go'nun Tek Döngüsü

**Ne işe yarar?** Go'da **bir tane döngü var: `for`**. While ve do-while yoktur — `for` hepsi yapar.

### A) Klasik For (C/Java gibi)

**Syntax:**
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**Açıklama:**
- `i := 0` — başlat
- `i < 10` — koşul (doğru oldukça döngü çalışır)
- `i++` — her tur sonunda artır
- `{}` — yapılacaklar

**Pratik (port tarama):**
```go
for port := 1; port <= 65535; port++ {
    // her portu kontrol et
}
```

> **Not:** C/Java'daki **parantezler** (`for (...)`) Go'da **yoktur**. Doğrudan `for` başlarsın. Ama `{}` **zorunlu**.

---

### B) While Gibi For (koşul sadece)

Go'da while yoktur; ama `for` + sadece koşul kullanırsın:

**Syntax:**
```go
i := 0
for i < 10 {          // koşul
    fmt.Println(i)
    i++
}
```

**Açıklama:**
- Başlatma ve artış kodda — `for` sadece koşula bakar
- Koşul doğu oldukça döngü çalışır
- Koşul yanlış olunca çıkar

**Pratik (bağlantı testi):**
```go
for hedefAktifMi() {    // fonksiyon true dönerse devam et
    fmt.Println("Hedef canlı")
    time.Sleep(2 * time.Second)
}
```

---

### C) Sonsuz For (do-while yerine)

**Syntax:**
```go
for {
    fmt.Println("Sonsuz döngü")
    // `break` ile çıkmalısın
}
```

**Pratik (sonsuz tarama + break ile çıkış):**
```go
for {
    port++
    if port > 65535 {
        break  // çık
    }
    // port taraması
}
```

---

### D) Range (diziler üzerinde)

**Syntax:**
```go
hedefler := []string{"1.1.1.1", "8.8.8.8", "9.9.9.9"}
for i, hedef := range hedefler {
    fmt.Println(i, hedef)  // 0 1.1.1.1 | 1 8.8.8.8 | ...
}
```

**Açıklama:**
- `i` = index (0, 1, 2...)
- `hedef` = değer
- Range otomatik döngü sayısını belirler

> **Ders:** For döngüsü Go'da çok yönlü — klasik, while, sonsuz, range hepsi birer `for` varyasyonu.

---

## 2️⃣ IF-ELSE — Koşul İfadeleri

**Ne işe yarar?** Belirli koşullar sağlanırsa kod çalışır.

### A) Basit If

**Syntax:**
```go
if port == 443 {
    fmt.Println("HTTPS")
}
```

**Önemli:**
- **Parantez yok** (`if port == 443` — `if (port == 443)` değil)
- **`{}` zorunlu** (tek satırlı bile)

**Pratik:**
```go
if status > 0 && status < 300 {   // VE bağlacı
    fmt.Println("Başarılı")
}

if status == 404 || status == 500 {  // VEYA bağlacı
    fmt.Println("Hata")
}

if !hedefAktif {  // DEĞİL bağlacı
    fmt.Println("Hedef kapalı")
}
```

---

### B) If-Else

**Syntax:**
```go
if port == 443 {
    fmt.Println("HTTPS")
} else if port == 80 {
    fmt.Println("HTTP")
} else {
    fmt.Println("Diğer")
}
```

---

### C) 🎯 If + Kısa Statement (Go'ya Özgü!)

**Ne işe yarar?** If'in koşulundan **hemen öncesinde** bir işlem yapabilirsin.

**Syntax:**
```go
if sonuc := hedefTaraması("1.1.1.1"); sonuc == "açık" {
    fmt.Println("Port açık:", sonuc)
}
// Buradan sonuc değişkenine erişemezsin (kapsam bitmiş)
```

**Açıklama:**
- `sonuc := hedefTaraması(...)` — önce bu fonksiyon çalışır
- Koşul kontrol edilir
- If bloğunun sona erdiğinde `sonuc` ölür (scope: sadece if-else içi)

**Neden faydalı?**
1. **Kapsam temiz** — değişken sadece ihtiyaç duyduğun yerde yaşar
2. **Tek satırda işlem + kontrol** — C gibi dillerde iki satır yazardın

**Pratik:**
```go
if err := dosyayıAc("scan.txt"); err != nil {
    fmt.Println("Dosya açılmadı:", err)
    return
} else {
    fmt.Println("Dosya açıldı")
}
```

> **Ders:** Bu yapı `&&` (VE) değildir — bir **ön-işlem + koşul** yapısıdır. Go'ya özgü.

---

## 3️⃣ SWITCH — Seçim Yapısı

**Ne işe yarar?** Uzun `if-else-if` zincirini daha temiz yazmanın yolu.

### A) Basit Switch

**Syntax:**
```go
gun := time.Now().Weekday()

switch gun {
case time.Monday:
    fmt.Println("Pazartesi")
case time.Tuesday:
    fmt.Println("Salı")
default:
    fmt.Println("Diğer gün")
}
```

**Önemli:**
- **Otomatik `break`** — C/Java'da `break` yazman lazım, Go yapıyor
- Her `case` bulununca sadece o çalışır, sonrasına bakmaz
- `default` — hiçbiri tutmazsa

---

### B) Switch + Koşul (Şart)

**Syntax:**
```go
port := 443

switch port {
case 80:
    fmt.Println("HTTP")
case 443:
    fmt.Println("HTTPS")
case 8080:
    fmt.Println("Geliştirme")
default:
    fmt.Println("Bilinmeyen")
}
```

---

### C) Switch Üzerinde Koşul (İleri)

**Syntax:**
```go
x := 5

switch {
case x < 0:
    fmt.Println("Negatif")
case x == 0:
    fmt.Println("Sıfır")
case x > 0:
    fmt.Println("Pozitif")
}
```

**Açıklama:**
- `switch` sonuna değer yazmazsın
- Her `case`'de tam koşul yazarsın
- If-else zinciri gibi çalışır ama daha temiz

---

### D) Fallthrough (İsteğe Bağlı)

**Syntax:**
```go
status := 200

switch status {
case 200:
    fmt.Println("Tamam")
    fallthrough  // bir sonraki case'e de git!
case 201:
    fmt.Println("Oluşturuldu")
default:
    fmt.Println("Hata")
}
```

**Çıktı:** "Tamam" + "Oluşturuldu" (fallthrough olmadan sadece "Tamam" yazardı)

> **Ders:** Go otomatik break yapıyor. Eğer istersen `fallthrough` ile bir sonraki case'e gidebilirsin (nadir kullanım).

---

### E) Switch + Kısa Statement

**Syntax:**
```go
switch ip := hedefiGetir(); ip {
case "1.1.1.1":
    fmt.Println("Cloudflare")
case "8.8.8.8":
    fmt.Println("Google")
default:
    fmt.Println("Bilinmeyen")
}
```

---

## 4️⃣ DEFER — Fonksiyon Bittiğinde Çalıştır

**Ne işe yarar?** Fonksiyon bittiğinde çalışması gereken kod — ama koddaki yerinde yazarsın.

### Temel Mantık

**Syntax:**
```go
func dosyayıOku() {
    dosya := aç("scan.txt")
    
    defer dosya.Kapat()  // Fonksiyon bittiğinde çalış!
    
    fmt.Println(dosya.Oku())
    // fonksiyon bitiyor → defer çalışır → dosya kapatılır
}
```

**Akış:**
1. `dosya := aç(...)` — dosyayı aç
2. `defer dosya.Kapat()` — kodu ertelemeye al (henüz çalıştırma)
3. `fmt.Println(dosya.Oku())` — dosyayı oku
4. Fonksiyon bitiyor → ertelenen `dosya.Kapat()` çalışır

---

### Neden Faydalı?

**Sorun (defer olmasa):**
```go
func dosyayıOku() {
    dosya := aç("scan.txt")
    
    veri := dosya.Oku()
    if veri == "" {
        return  // ❌ dosya kapalı kalır! (kapama unuttuk)
    }
    
    fmt.Println(veri)
    dosya.Kapat()  // bu satıra ulaşamayız
}
```

**Çözüm (defer ile):**
```go
func dosyayıOku() {
    dosya := aç("scan.txt")
    defer dosya.Kapat()  // ne olursa olsun, sonunda kapanacak
    
    veri := dosya.Oku()
    if veri == "" {
        return  // ✅ defer otomatik çalışır, dosya kapanır
    }
    
    fmt.Println(veri)
    // fonksiyon sona erince de defer çalışır
}
```

> **Ders:** Defer = "kapı çıkarken ışığı söndür". Çıkarken nereye çıkarsan çık, ışık söneceğini biliyorsun.

---

### Defer + Değişken Değeri

**Önemli:** Defer'ın argümanları **hemen** değerlendirilir.

**Syntax:**
```go
isim := "Ali"

defer fmt.Println("Hoşça kalın,", isim)  // "Ali" o an hafızaya alınır

isim = "Veli"
fmt.Println("Merhaba,", isim)

// Çıktı:
// Merhaba, Veli
// Hoşça kalın, Ali  (ilk andaki değeri kullanır!)
```

---

### Birden Çok Defer (LIFO - Ters Sıra)

**Syntax:**
```go
func main() {
    defer fmt.Println("1. Çıkılıyor")
    defer fmt.Println("2. Işık söndürülüyor")
    defer fmt.Println("3. Kapı kilitliliyor")
    
    fmt.Println("Fonksiyon çalışıyor...")
}

// Çıktı:
// Fonksiyon çalışıyor...
// 3. Kapı kilitliliyor
// 2. Işık söndürülüyor
// 1. Çıkılıyor
```

**Açıklama:** Defer'lar **ters sırada** çalışır (en son yazılan ilk çalışır).

---

## 📊 Kontrol Akışı — Özet Tablosu

| Yapı | Syntax | Ne zaman |
|---|---|---|
| **For (klasik)** | `for i := 0; i < 10; i++` | C/Java gibi sayacı kontrol |
| **For (while)** | `for i < 10` | Koşul sağlandıkça |
| **For (sonsuz)** | `for { break }` | Sonsuz, break ile çık |
| **For (range)** | `for i, v := range arr` | Dizi/slice üzerinde |
| **If** | `if koşul {}` | Parantez yok, `{}` zorunlu |
| **If + statement** | `if x := f(); x > 0 {}` | Ön-işlem + koşul |
| **Switch** | `switch val { case ... }` | Otomatik break |
| **Switch + fallthrough** | `case 1: ... fallthrough` | Bir sonraki case'e git |
| **Defer** | `defer f()` | Fonksiyon bitince çalış |

---

## 🎯 Pratik Örnek — Tarama Programı

```go
package main

import (
    "fmt"
    "time"
)

func taramaYap(baslaPort, bitisPort int) {
    fmt.Println("Tarama başlıyor...")
    
    for port := baslaPort; port <= bitisPort; port++ {
        // Koşul kontrol
        if port%100 == 0 {
            fmt.Println("İlerleme:", port)
        }
        
        // Switch ile port tipi belirle
        switch port {
        case 80:
            fmt.Println("Port 80: HTTP")
        case 443:
            fmt.Println("Port 443: HTTPS")
        case 22:
            fmt.Println("Port 22: SSH")
        default:
            // çoğu port buraya düşer
        }
    }
    
    fmt.Println("Tarama bitti")
}

func main() {
    defer fmt.Println("Program kapanıyor...")
    
    taramaYap(1, 1000)
    
    // Program bıittiğinde defer çalışır
}
```

---

## ✅ Faz 2 — Bitti
- [x] For döngüsü (klasik, while, sonsuz, range)
- [x] If-else (parantez yok, `{}` zorunlu)
- [x] If + kısa statement (Go'ya özgü)
- [x] Switch (otomatik break, fallthrough)
- [x] Defer (temizlik, LIFO, değişken değeri)

**Sırada:** Slice, Map, Struct (Faz 2 devam)
