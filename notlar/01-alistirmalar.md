# Faz 1 — Alıştırmalar (Gün 1 pekiştirme)

> Kural: Kodu **kendim** yazacağım. Sadece Gün 1 kavramları (var/:=, type-last, zero value,
> scope, named return, byte/rune, explicit conversion, type inference, const).
> Henüz `if`/`for`/`switch` YOK — hepsi bunlarsız çözülebilir.
> Her görev için `projeler/alistirmalar/` altında ayrı `.go` dosyası aç, `go run` ile çalıştır.

---

## Görev 1 — Hedef Bilgi Kartı (var + := + zero value)
Bir tarama hedefinin bilgilerini tut ve yazdır:
- `hedefIP` (string) → `:=` ile
- `port` (int) → `var` ile
- `aktifMi` (bool) → değer **atama**, zero value'yu gör
- `fmt.Printf` ile tek satırda düzgün yazdır (`%s`, `%d`, `%t`)

**Başarı kriteri:** Çıktı → `Hedef: 10.10.10.5:443 | Aktif: false`
**İpucu:** `%t` bool için.

---

## Görev 2 — Çoklu Tanım & Type-Last (var bloğu)
Tek `var (...)` bloğunda tanımla:
- `minPort, maxPort int = 1, 65535`
- `protokol string = "tcp"`
Sonra hepsini yazdır.

**Başarı kriteri:** `Protokol: tcp | Aralik: 1-65535`
**İpucu:** Faktör tanımlama (blok) notundaki gibi.

---

## Görev 3 — Explicit Conversion (int ↔ float)
Bir tarama süresi hesabı:
- `toplamPort := 1000` (int)
- `saniyeBasina := 40.0` (float64)
- Süreyi hesapla: `toplamPort / saniyeBasina` → **doğrudan çalışmaz!** Neden?
- Doğru şekilde dönüştürüp `sure` (float64) bul, sonra bir de `int(sure)` ile tam saniyeyi yazdır.

**Başarı kriteri:** `Sure: 25.00 sn (yaklasik 25 sn)`
**İpucu:** Go int'i float'a otomatik çevirmez → `float64(toplamPort)`. `int()` ondalığı keser.

---

## Görev 4 — byte & rune (alias'lar)
- Bir karakter al: `harf := 'A'` → bunun **rune (int32)** sayısal değerini yazdır.
- `parola := "gizli"` → bunu `[]byte`'a çevirip yazdır (ham byte dizisi görünsün).

**Başarı kriteri:** `'A' = 65` ve parola byte dizisi olarak `[103 105 122 108 105]` gibi.
**İpucu:** `int('A')` sayıyı verir. `[]byte("...")` string'i byte dizisine çevirir.

---

## Görev 5 — const & Değişmezlik
- `const MaxPort = 65535` tanımla (paket seviyesinde).
- `main` içinde `const zamanAsimi = 30` tanımla (fonksiyon içi).
- Yorum satırı olarak: `MaxPort` değiştirmeyi dene, **neden hata verdiğini** yaz.
- Yorum satırı olarak: `const x := 5` neden yanlış, açıkla.

**Başarı kriteri:** Program çalışıyor + iki yorum açıklaması doğru.

---

## Görev 6 — Named Return + Naked Return (BİRLEŞTİRME)
Bir fonksiyon yaz: `hedefOzeti(ip string, port int)` →
- İsimli dönüş değerleri: `(adres string, guvenli bool)`
- Gövdede: `adres = ip` (basitçe), `guvenli = false` ata
- Sonda **naked return** (`return`) kullan
- `main`'den çağır, sonucu yazdır

**Başarı kriteri:** `Adres: 1.1.1.1:80 | Guvenli: false`, fonksiyon `return` ile bitiyor (çıplak).
**İpucu:** `func hedefOzeti(ip string, port int) (adres string, guvenli bool) { ... return }`

---

## 🏆 Bonus — Exported/Unexported (kavram)
Aynı dosyada bir `struct` tanımla:
```go
type Hedef struct {
    IP    string   // exported
    port  int      // unexported
}
```
- Yorum olarak: Bu struct başka bir paketten import edilseydi, hangi alana erişilebilir, hangisine erişilemezdi? Neden?

**Başarı kriteri:** Doğru açıklama (IP erişilir çünkü büyük harf, port erişilemez çünkü küçük).

---

## Kontrol
- [+] Görev 1  - [+] Görev 2  - [+] Görev 3
- [+] Görev 4  - [+] Görev 5  - [+] Görev 6  - [ ] Bonus
