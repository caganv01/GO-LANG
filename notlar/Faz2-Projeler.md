# 🎯 Faz 2 — Proje Görevleri (Gündelik + Siber Güvenlik)

> Faz 1-2'yi kullanarak 4 proje. Her birini `projeler/` altında kendi klasöründe yap.
> Kodu **sen** yazacaksın, takıldığında "nasıl yapacağım" sorabilirsin.

---

## Proje 1 — Şifre Gücü Kontrolörü 🔐

**Senaryo:** Bir başladı şifre oluştururken, o şifrenin güvenli olup olmadığını kontrol eden bir tool.

**Hedef:**
```
Şifrenizi girin: 123
Şifre Gücü: ZAYıF (Sadece rakam)

Şifrenizi girin: MyPassword2024!
Şifre Gücü: GÜÇLÜ (Harf + Rakam + Sembol)
```

**Yapılacaklar:**
1. Kullanıcıdan şifre al (fmt.Scanln)
2. Şifre **uzunluğuna** göre kontrol et:
   - 0-5 → Çok kısa ❌
   - 6-8 → Zayıf ⚠️
   - 9-12 → Orta 📊
   - 13+ → Güçlü ✅
3. Şifrede ne var kontrol et:
   - Sadece harf → "sadece harf"
   - Harf + rakam → "harf + rakam"
   - Harf + rakam + sembol → "çok güçlü"
4. Switch kullanarak sonucu yazdır

**Başarı Kriteri:**
```go
// Komut satırı:
// go run password.go
// Şifrenizi girin: Test123!
// Uzunluk: 9 | İçerik: Harf+Rakam+Sembol
// Sonuç: ORTA GÜÇ (9-12 karakter + karışık içerik)
```

**İpuçları:**
- `strings.Contains()` kullan (paket: `strings`)
- `len()` ile uzunluk bul
- `rune` ile her karakteri kontrol et
- Switch + if birlikte kullan

**Klasör:** `projeler/02-password-checker/`

---

## Proje 2 — IP Tarama Raporu 📡

**Senaryo:** Güvenlik testi sırasında taradığın IP'ler ve portları bir dosyaya yazıp raporlayan tool.

**Hedef:**
```
Hedef: 192.168.1.1 | Açık Portlar: 80, 443, 22
Hedef: 10.0.0.5   | Açık Portlar: 3306, 5432
---
Toplam Hedef: 2
Toplam Açık Port: 5
Ortalama: 2.5
```

**Yapılacaklar:**
1. **Hedef IP'ler** tanımla (const olarak 3-4 tane):
   ```go
   const hedef1 = "192.168.1.1"
   const hedef2 = "10.0.0.5"
   ```
2. **Açık portlar** (simulate):
   ```go
   var hedef1Portlari []int = []int{80, 443, 22}
   var hedef2Portlari []int = []int{3306, 5432}
   ```
3. **For + range** ile her hedefi tara:
   ```go
   for i, port := range hedef1Portlari {
       // her portu raporla
   }
   ```
4. **Defer** ile sonunda özet yazılsın (program bittiğinde)
5. **Toplam, ortalama** hesapla

**Başarı Kriteri:**
```
Tarama Raporu
---
192.168.1.1 → 80, 443, 22
10.0.0.5 → 3306, 5432
---
Toplam Portlar: 5
Ortalama/Hedef: 2.5
[Program kapanıyor...]  ← defer çalıştı
```

**İpuçları:**
- `const` ile sabit IP'ler tanımla
- `var hedefler []string` ile bir dizi yap
- Nested for (IP'ler + portlar)
- `defer` + `fmt.Println()` sonunda özet için
- Toplam port sayısı + hedef sayısı = ortalama

**Klasör:** `projeler/03-scan-report/`

---

## Proje 3 — İşletim Sistemi Dedektörü 🖥️

**Senaryo:** Bir sistemin kullanıcı agent'ine bakarak ne OS kullandığını tahmin eden tool.

**Hedef:**
```
User Agent: Mozilla/5.0 (Windows NT 10.0)
Sonuç: Windows 10 ✓

User Agent: Mozilla/5.0 (iPhone OS 15_0)
Sonuç: iOS 15 ✓

User Agent: Mozilla/5.0 (Linux; U; Android 11)
Sonuç: Android 11 ✓
```

**Yapılacaklar:**
1. **Örnek User Agent'ler** tanımla (const):
   ```go
   const ua1 = "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"
   const ua2 = "Mozilla/5.0 (iPhone OS 15_0) AppleWebKit"
   const ua3 = "Mozilla/5.0 (X11; Linux x86_64)"
   ```
2. **Switch** ile her UA'yı kontrol et:
   - "Windows" varsa → Windows
   - "iPhone" varsa → iOS
   - "Linux" varsa → Linux
   - "Android" varsa → Android
   - Yoksa → Bilinmeyen
3. **If + short statement** kullanarak version bilgisi çıkar
4. Sonucu yaz

**Başarı Kriteri:**
```go
// User Agent'te "Windows NT 10.0" → "Windows 10"
// User Agent'te "iPhone OS 15" → "iOS 15"
// User Agent'te "Linux" → "Linux (Unknown Version)"
```

**İpuçları:**
- `strings.Contains(ua, "Windows")` ile kontrol et
- `strings.Index()` ile pozisyon bul
- Switch + case'ler farklı OS'ler için
- If-else ile version parsing

**Klasör:** `projeler/04-ua-detector/`

---

## Proje 4 — Zamanlanmış Güvenlik Denetimi 🔔

**Senaryo:** Her X saniyede sistemin belli durumlarını kontrol et, sonunda rapor ver.

**Hedef:**
```
[00:00] Kontrol 1: Firewall AÇIK ✓
[01:00] Kontrol 2: Şüpheli bağlantı yok ✓
[02:00] Kontrol 3: Dosya bütünlüğü: TAMAM ✓
---
Toplam Kontrol: 3
Hata Sayısı: 0
Durum: SİSTEM SAĞLIKLI ✅

[Temizlik işlemleri yapılıyor...]  ← defer
[Denetim sonlandırıldı]
```

**Yapılacaklar:**
1. **Kontrol listesi** (sabitler):
   ```go
   const kontrol1 = "Firewall durumu"
   const kontrol2 = "SSH erişimi"
   const kontrol3 = "Dosya izinleri"
   ```
2. **For döngüsü** (3 tur, her turda 1 kontrol):
   ```go
   for i := 1; i <= 3; i++ {
       // kontrol et, sonucu yaz
       time.Sleep(1 * time.Second)  // simülasyon
   }
   ```
3. **Switch** ile kontrol türüne göre mesaj ver
4. **Sayaçlar** (başarılı, başarısız):
   ```go
   var basarili, basarisiz int
   ```
5. **Defer** ile temizlik mesajını ver
6. Sonunda özet (toplam, başarı oranı)

**Başarı Kriteri:**
```
Güvenlik Denetimi Başlıyor...
[01] Firewall → AÇIK ✓
[02] SSH Port → AÇIK ✓
[03] Şifreler → GÜÇLÜ ✓
---
Denetim Tamamlandı
Başarılı: 3/3
Durum: SAĞLIKLI
[Cleanup: Temizlik yapılıyor...]  ← defer
```

**İpuçları:**
- `time.Sleep()` için `"time"` paketini import et
- For + i++ ile turları say
- Her turda if-else ile rastgele başarı/başarısızlık ver:
  ```go
  if rand.Intn(2) == 0 {
      fmt.Println("✓")
      basarili++
  }
  ```
- Defer'ı main'in başında yaz
- `float64(basarili) / float64(toplam) * 100` ile yüzde hesapla

**Klasör:** `projeler/05-security-audit/`

---

## 📋 Proje Zorluk Sırası

| # | Proje | Faz 1 Kullanımı | Faz 2 Kullanımı | Zorluk |
|---|-------|---|---|---|
| 1 | Password Checker | var, const, if-else | if-else, switch, for | ⭐ Kolay |
| 2 | Scan Report | const, slice | for, range, defer | ⭐⭐ Orta |
| 3 | UA Detector | const, string ops | switch, if-statement | ⭐⭐ Orta |
| 4 | Security Audit | const, int ops | for, defer, switch | ⭐⭐⭐ Zor |

---

## 🚀 Başlama Adımları

1. **Proje 1**'den başla (Password Checker)
2. Kodu **tam sen yaz** — AI'a "nasıl yazabilirim" sor, kodu sen yapıştır
3. `go run` ile çalıştır
4. Başarı kriterini karşılıyor mu kontrol et
5. Tamamlayınca **commit + push**:
   ```bash
   git add projeler/02-password-checker/
   git commit -m "Proje 1: Password Checker tamam"
   git push
   ```
6. Sonraki projeye geç

---

## ✅ Her Proje Tamamlandığında
- [ ] Proje 1 — Password Checker
- [ ] Proje 2 — Scan Report
- [ ] Proje 3 — UA Detector
- [ ] Proje 4 — Security Audit

**Tamamlayınca:** Faz 3'e (Slice, Map, Struct) geçeriz.

---

## 💡 Bonus Zorluk Seviyeleri
Her proje için ileri sürüm:
- **Password Checker v2:** Dosyadan şifre listesi oku, rapor ver
- **Scan Report v2:** Gerçek portları (net paketini) test et
- **UA Detector v2:** Tarayıcı sürümünü de çıkar
- **Security Audit v2:** Denetim sonuçlarını dosyaya kaydet
