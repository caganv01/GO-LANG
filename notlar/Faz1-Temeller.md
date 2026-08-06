# 📚 Faz 1 — Go Temelleri (Adım Adım, Syntax Açık)

> Kişisel çalışma notu. A Tour of Go üzerinden öğrenilen konular, kendi cümlelerimle + C/Java karşılaştırmalı. Güvenlik/oyun senaryolarıyla pekiştirilmiş.

---

## 1️⃣ Exported vs Unexported (Erişim Kuralı)

**Ne işe yarar?** Go'da erişim (public/private) **ismin ilk harfi** ile belirtilir.

| **Harf** | **Adı** | **Anlamı** | **Örnek** |
|---|---|---|---|
| **Büyük** | Exported | Paket dışından erişilir (public) | `Pi`, `Println`, `User` |
| **küçük** | Unexported | Sadece paketi içinde (private) | `pi`, `age`, `secret` |

**Syntax:**
```go
var Pi = 3.14159      // Büyük P → dışarıdan erişilir ✅
var pi = 3.14159      // küçük p → kendi paketi içinde ✅
// Başka paketten math.pi istersen hata: unexported name ❌
```

> **Ders:** C/Java'daki `public`/`private` anahtar kelimeleri yerine Go bunu isimle çözer — daha az kelime, daha açık.

---

## 2️⃣ `fmt` Paketi — Çıktı Paketi

**Ne işe yarar?** Biçimlendirilmiş çıktı (ekrana yazdırma).

**Karşılıkları:**
- C: `stdio.h` (`printf`)
- Java: `System.out.println()`
- Go: `fmt.Printf()`, `fmt.Println()`

**Syntax (en sık kullanılanlar):**
```go
fmt.Println("Merhaba")              // Otomatik satır sonu
fmt.Printf("Port: %d\n", 443)       // Biçimli (format string)
result := fmt.Sprintf("IP: %s", ip) // String'e döndürür, basmaz
```

**Format verb'leri (% ile başlayanlar):**

| Verb | Ne için |
|---|---|
| `%s` | String (metin) |
| `%d` | Integer (tam sayı) |
| `%f` | Float (ondalık sayı) |
| `%t` | Boolean (true/false) |
| `%v` | Herhangi bir değer (otomatik) |
| `%x` | Hexadecimal (16'lı sayı) |

> **Ders:** `Printf` satırın sonuna `\n` koymadığın sürece sonraki çıktı yapışır — hatırlat.

---

## 3️⃣ Değişken Tanımlama — `:=` vs `var` (ÇOK ÖNEMLİ)

Go'da **iki farklı** değişken tanımlama yöntemi var. Hangisini ne zaman kullanacağını anlamak **kritik**.

### A) Kısa Tanımlama: `:=` (Short Variable Declaration)

**Ne işe yarar?** Değişkeni oluştur + türünü otomatik belirle + değer ata — **tek satırda**.

**Kuralı:** ⚠️ **SADECE FONKSİYON İÇİNDE KULLAN.**

**Syntax:**
```go
func main() {
    x := 5              // int (Go otomatik belirler)
    y := "merhaba"      // string (Go otomatik belirler)
    z := 3.14           // float64 (Go otomatik belirler)
    
    // Bu ise HATAAAAA:
    // x := 5          → fonksiyon dışında = Syntax Error
}
```

**Ne olur arka planda?**
```go
x := 5
// Gerçekten bu yapılır:
// var x int = 5  (Go derleyicisi otomatik belirler)
```

**Neden `:=`?** Yazı yazarken hızlı (3 karakter), `var x int = 5` yerine.

---

### B) Klasik Tanımlama: `var` (Açık Tip)

**Ne işe yarar?** Değişken oluştur, türünü **sen belirt** (opsiyonal olarak).

**Kural:** Fonksiyon dışında **ve** içinde kullanılabilir.

**Syntax — Tek tanım:**
```go
var x int = 5         // Açıkça int dedik
var y string = "ip"   // Açıkça string dedik
var z = 3.14          // Tip yazmadık, Go float64 bulur (type inference)
```

**Syntax — Çoklu tanım (factored declaration):**
```go
var (
    minPort    int    = 1
    maxPort    int    = 65535
    protokol   string = "tcp"
)
// Paket seviyesinde yazıldı (fonksiyon dışında)
```

**Syntax — Tür yazılmazsa (type inference):**
```go
var x = 5          // Go bunu int yapar (5 tam sayı olduğu için)
var y = "merhaba"  // Go bunu string yapar
```

---

### C) Zero Value (Değer Atamadığında Ne Olur?)

Eğer `var` kullanıp değer **atamadan** bırakırsan Go **otomatik bir sıfır değeri** atar:

```go
var x int       // x = 0 (tam sayılar için sıfır)
var s string    // s = "" (metin için boş string)
var b bool      // b = false (boolean için yanlış)
var f float64   // f = 0.0 (ondalık için sıfır)
```

> **Ders:** `:=` **her zaman değer ister**; `var` değer yoksa zero value atar.

---

### ⭐ `:=` vs `var` — Özet Tablo

| Yön | `:=` | `var` |
|---|---|---|
| **Yazım** | `x := 5` | `var x int = 5` |
| **Nerede kullan** | Sadece fonksiyon içinde | Her yerde |
| **Tür belirtmek** | İsteğe bağlı (otomatik) | İsteğe bağlı |
| **Değer ister mi** | Zorunlu | Opsiyonel (yoksa zero value) |
| **Hız** | Daha kısa yazmak | Açık olmak, global değişken |

**Pratik kod:**
```go
package main

import "fmt"

var globalX int = 100  // Paket seviyesi, var ile

func main() {
    x := 5              // Fonksiyon içi, := ile (hızlı)
    var y int = 10      // Fonksiyon içi, var ile (açık)
    
    fmt.Println(globalX, x, y)  // 100 5 10
}
```

---

## 4️⃣ Type-Last (Türün Sonda Yazılması)

**Fark ne?** C, Java: `int x` | Go: `x int`

**Neden?** Go'yu okumayı daha doğal yaptı — "x adında bir integer" → `x int`.

**Syntax:**
```go
// ❌ YANLIIŞ (C/Java stili):
int x;
float y;

// ✅ DOĞRU (Go stili):
var x int
var y float64
```

**Çoklu tanım (type-last faydalı olur):**
```go
var x, y int              // İkisi de int
var minPort, maxPort int  // İkisi de int (bir kez yazarsın)
var port int, host string // Port int, host string

// Çoklu ve değer:
var x, y int = 1, 2       // x=1, y=2
```

> **Ders:** Type-last'e alışınca `var a, b, c int` yazıp üç satırdan bir satırda tanımlarsın.

---

## 5️⃣ Kapsam (Scope) — Nerede Yaşarsa

**Değişkenler iki yerde yaşar:**

| Seviye | Nerede | Erişim | Yaşam |
|---|---|---|---|
| **Paket (Global)** | `func` dışında, dosya üstü | O paket içindeki tüm fonksiyon | Program kapanana kadar |
| **Fonksiyon (Local)** | `func` içinde, `{}` arası | Sadece o fonksiyon | Fonksiyon bitene kadar |

**Syntax:**
```go
package main

var globalX = 100  // Paket seviyesi — main ve diğer fonksiyonlar görür

func fonk1() {
    var localX = 5  // Local — sadece fonk1 içinde yaşar
}

func fonk2() {
    // localX ?  → Erişemez, hata: undefined (tanımlanmamış)
    // globalX ? → Erişir, 100 değeri
}
```

---

## 6️⃣ Fonksiyon Satırı Başı Kuralı

Fonksiyon **dışında** bir satır **mutlaka** şu anahtar kelimelerden biriyle başlamalı:

- `var` — değişken
- `func` — fonksiyon
- `import` — paket
- `type` — kendi tipi
- `const` — sabit

**Syntax (HATALAR):**
```go
package main

x := 5              // ❌ HATA! Paket seviyesinde := kullanamazsın

var x = 5           // ✅ DOĞRU

ipAdresi = "1.1.1.1" // ❌ HATA! Satır bir keyword ile başlamalı

var ipAdresi = "1.1.1.1"  // ✅ DOĞRU

func main() {
    y := 10         // ✅ Fonksiyon içinde := tamam
}
```

---

## 7️⃣ Named Return + Naked Return

**Ne işe yarar?** Fonksiyonun dönüş değerlerine imzada isim ver, sonda sadece `return` yaz.

### Sıradan dönüş (named değil):
```go
func topla(a, b int) (int, int) {
    return a + b, a - b
}
```

### Named return (isimli):
```go
func topla(a, b int) (toplam int, fark int) {
    // toplam ve fark zaten var, var ile tanımlananlar gibi
    toplam = a + b
    fark = a - b
    return  // ← Naked return — sonuçları otomatik döndürür
}
```

**Pratik:**
```go
func hedefOzeti(ip string, port int) (adres string, guvenli bool) {
    adres = ip
    guvenli = false
    return  // Bu adres ve guvenli değerlerini döndürür
}

func main() {
    a, g := hedefOzeti("1.1.1.1", 80)
    fmt.Println(a, g)  // 1.1.1.1 false
}
```

> **Ders:** Kısa fonksiyonlarda naked return okunur; uzun fonksiyonlarda açık `return adres, guvenli` daha iyidir.

---

## 8️⃣ Temel Tipler (Basic Types)

### Blok Tanımlama (Factored Declaration):
```go
var (
    yapı      bool       = false
    maxPort   uint64     = 65535
    hedefIP   string     = "127.0.0.1"
    hiz       float64    = 120.5
)
```

### Mimariye Bağlı Tipler (`int`, `uint`):
Sistem 32-bit ise 32 bit, 64-bit ise 64 bit.

```go
var x int   // Sistemin native size'ı (genellikle 64-bit günümüzde)
```

### Özel Boyutlu Tipler:
```go
var kanal uint8   // 1 byte (0-255 arası)
var port uint16   // 2 byte

// Takma adlar:
byte   = uint8      // Ağ paketi, ham veri
rune   = int32      // Unicode karakter ('A', 'Ş', '😀')
```

---

## 9️⃣ Explicit Conversion — Tür Dönüştürme

**Kural:** Go tip dönüştürmeyi **otomatik yapmaz**. Sen `T(value)` ile yaparsın.

**Syntax:**
```go
var x int = 100
var y float64 = float64(x)    // ✅ Açıkça dönüştür

// var y = x  // ❌ HATA! Go otomatik çevirmez

// int'ten float'a:
sonuc := float64(1000) / 40.0  // ✅ İki taraf float64
tam := int(sonuc)              // ✅ float64 → int (kesme yapılır)
```

**int() kesme (truncate) yapar, yuvarlamaz:**
```go
var konum float64 = 105.89
pixel := int(konum)  // 105 (0.89 atılır, 106 değil)
```

> **Ders:** C'de sessizce kaybolan `0.89` Go'da seni açıkça dönüştürmeye zorlayarak bulur.

---

## 🔟 Type Inference (Tip Çıkarımı)

**Sağ tarafın tipi belliyse:**
```go
var x int = 100
y := x          // y de int (x'in tipini alır)
```

**Ham sayıysa, varsayılan tip:**
```go
x := 42              // int (tam sayı olduğu için)
y := 3.14            // float64 (ondalık olduğu için — hassasiyet)
z := 0.5 + 0.3i      // complex128 (karmaşık sayı)
```

> **Ders:** `:=` size sığması için en geniş tipi seçer. Bellek kritikse `var x uint8 = 42` ile elle belirt.

---

## 1️⃣1️⃣ Sabitler (`const`)

**Kural:** Derleme anında koda gömülür, **asla değişmez**.

**Syntax:**
```go
const MaxPort = 65535      // Paket seviyesi
const TaramaModuAktif = true

func main() {
    const ZamanAsimi = 30  // Fonksiyon içi de olabilir
    // MaxPort = 100       ❌ HATA! Sabit değiştirilemez
}
```

### `:=` ile `const` kullanılamaz (3 sebep):
1. **`:=` = değişken kısayolu** (arka planda `var`). Sabit değişmez, çelişki.
2. **Compile-time garanti:** Derleyici açıkça `const` görmek ister.
3. **Okunabilirlik:** `const` okunca kodun `HedefPort` değişmeyeceği belli.

**Karşılaştırma:**
```go
const SafBir = 443         // Asla değişmez, betona kazı

var AyarlanabilirPort int = 443  // Sonra değiştirebilirim
```

**Zero value yoktur:** const her zaman değer ister.

### Sabite atama denemesi → derleme hatası
`MaxPort = 100` yazarsan Go şu hatayı verir (runtime'da değil, **derlenmez**):
```
cannot assign to MaxPort (neither addressable nor a map index expression)
```
Sebep: const bir bellek hücresi değildir; değeri derleme anında doğrudan koda gömülür, o yüzden "atanacak bir yer" yoktur.

### Neden `const` ile `:=` kullanılamaz? (3 sebep derinlemesine)
1. **`:=` = değişken kısayolu.** Arka planda `var` + tip çıkarımıdır (`x := 10` → `var x int = 10`). `var` değiştirilebilir, `const` değişmez → değişken kısayolunu sabit için kullanmak mantıksal çelişki.
2. **Compile-time vs run-time.** `const` derlemede kesinleşip gömülür; `:=` run-time'da hafızada yer açan yerel değişken içindir. Derleyici garanti için açıkça `const` görmek ister.
3. **Okunabilirlik.** `HedefPort := 443` değişken mi sabit mi belli değil; `const HedefPort = 443` belirsizliği kaldırır.

> **Kısa özet:** `:=` → "değişken, sonra değiştirebilirim" · `const` → "asla değişmez, betona kaz"

---

## 📊 Hepsini Bir Satırda Özet

| Konsept | Syntax | Ne zaman |
|---|---|---|
| Exported | `Büyük` harf | Dışarıdan erişilsin |
| Unexported | `küçük` harf | Paketi içinde kalsın |
| Kısa tanımlama | `x := 5` | Fonk. içinde, hızlı |
| Açık tanımlama | `var x int = 5` | Açık olmak, global |
| Type-last | `var x, y int` | Birden çok, aynı tip |
| Naked return | `return` (named ise) | Kısa fonksiyon |
| Conversion | `float64(x)` | Tür değiştirmek |
| Sabit | `const x = 5` | Değişmeyecek değer |

---

## ✅ Faz 1 — Bitti
- [x] Exported/unexported (ilk harf kuralı)
- [x] `fmt`, `:=`, `var`, type-last, zero value
- [x] Scope (paket vs fonksiyon), satır başı keyword kuralı
- [x] Named return + naked return
- [x] Basic types, `byte`/`rune` alias, mimariye bağlı `int`
- [x] Explicit conversion `T(v)` (Go otomatik çevirmez)
- [x] Type inference & default tipler
- [x] `const` kuralları + overflow compile-time yakalama
