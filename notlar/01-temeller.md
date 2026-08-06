# 📚 Faz 1 — Go Temelleri (Adım Adım)

> Kişisel çalışma notu. Go — Faz 1 Temeller. A Tour of Go üzerinden öğrenilen konular, kendi cümlelerimle + C/Java karşılaştırmalı. Güvenlik/oyun senaryolarıyla pekiştirilmiş.

---

## 1️⃣ Exported vs Unexported (Erişim Kuralı)

Go'da Java'daki gibi `private`/`public`/`protected` **yoktur**. Erişimi **ismin ilk harfi** belirler:

| Yazım | Anlamı | Erişim |
|---|---|---|
| **Büyük** harfle başlar (`Pi`, `Println`) | exported | Paket **dışından** erişilebilir (public) |
| **küçük** harfle başlar (`pi`, `age`) | unexported | Sadece **kendi paketi** içinde (private) |

```go
type User struct {
    Name string  // exported — dışarıdan erişilir
    age  int     // unexported — sadece bu paket içinde
}
```

> **Ders:** `math.pi` hata verir çünkü küçük `p` → pakete kilitli. `math.Pi` çalışır. Bu kural değişken, fonksiyon, struct, struct alanı — **her isim** için geçerli. C/Java'daki `public` anahtar kelimesi yerine Go bunu isimle halleder.

---

## 2. `fmt` Paketi

`fmt` = "format" kısaltması. Go'nun **biçimlendirilmiş giriş/çıkış (I/O)** standart kütüphanesidir.

- C'deki `stdio.h`'a (`printf`, `scanf`) benzer.
- Java'daki `System.out` yapısına benzer.
- Görevi: programın terminalle / kullanıcıyla konuşmasını sağlamak.

```go
fmt.Println("Merhaba")             // satır + otomatik yeni satır
fmt.Printf("Port: %d\n", 443)      // biçimli çıktı (%d, %s, %f...)
```

---

## 3. Değişken Tanımlama: `var` ve `:=`

### `:=` — Kısa Değişken Tanımlama (short variable declaration)
Go'nun en sevilen kısayolu. Arka planda **iki iş** yapar: değişkeni **oluşturur (declare)** + türünü otomatik belirleyip **değer atar (initialize)**.

```go
hedefIP := "192.168.1.1"   // Go türü otomatik string yapar
// eşdeğeri:
var hedefIP string = "192.168.1.1"
```

> **Kural:** `:=` **sadece fonksiyon içinde** kullanılır. Ayrıca satır bu operatörle başlayamaz — dosya seviyesinde `var` şart.

### Type-Last: Türün en sonda olması
Çoğu dilde önce tür yazılır (`int x`). Go'da **önce isim, sonra tür**:

```go
var x, y int        // x ve y int
var port, thread int = 80, 4   // port=80, thread=4 (initializer, birebir eşleşir)
```

> **Ders:** Başlangıç değeri (initializer) varsa **tür yazılmayabilir**; Go değere bakıp türü çıkarır. Yoksa **zero value** atanır.

### Zero Value (Sıfır Değeri)
Değer atamazsan Go otomatik atar:

| Tip | Zero value |
|---|---|
| Sayılar (`int`, `float`) | `0` |
| String | `""` (boş) |
| Bool | `false` |

---

## 4. Kapsam (Scope): Paket vs Fonksiyon

| Seviye | Nerede tanımlanır | Erişim |
|---|---|---|
| **Paket (global)** | Herhangi bir `func` **dışında**, dosya üstünde | O dosyadaki tüm fonksiyonlar görür |
| **Fonksiyon (local)** | `{}` **içinde** | Sadece o fonksiyon içinde yaşar |

### Satır başı anahtar kelimeleri
Fonksiyon **dışında** bir satır şu kelimelerden biriyle başlamak zorundadır, yoksa **Syntax Error**:

`var` · `func` · `import` · `type` · `const`

```go
hedefIP := "1.1.1.1"   // ❌ fonksiyon dışında HATA (:= ile başlayamaz)
var hedefIP = "1.1.1.1" // ✅
```

---

## 5. Fonksiyonlar: İsimli Dönüş & Naked Return

### İsimlendirilmiş Dönüş Değerleri (Named Return Values)
Dönüş değerlerine **imzada isim** verebilirsin. Go bunları fonksiyonun başında `var` ile tanımlanmış hazır değişkenler gibi kabul eder.

### Çıplak Dönüş (Naked Return)
Dönüş isimlendirilmişse, sonda `return x, y` yazmana gerek yok — sadece `return`. Go isimli değişkenlerin son değerlerini döndürür.

```go
func topla(a, b int) (toplam int) {
    toplam = a + b
    return          // naked return — toplam'ı döndürür
}
```

> **Ders:** Naked return kısa fonksiyonlarda okunur; uzun fonksiyonlarda okunabilirliği bozar, dikkatli kullan.

---

## 6. Temel Tipler (Basic Types)

### Blok Halinde Tanımlama (Factored Declaration)
Tıpkı `import` gibi, `var`/`const` da blok halinde yazılabilir — kod temiz kalır:

```go
var (
    AktifMi   bool   = false
    MaksDeger uint64 = 1<<64 - 1
    HedefHost string = "127.0.0.1"
)
```

### Mimariye Bağlı Tipler: `int`, `uint`, `uintptr`
Bu tipler sistem 32-bit ise **32 bit**, 64-bit ise **64 bit** genişliğindedir (C'deki `size_t` mantığı).

> **Neden normalde `int`?** İşlemci kendi word size'ını (64-bit) en hızlı işler. Özel bir kısıtlaman yoksa döngü/hesapta direkt `int` kullan.

### Özel Boyutlu Tipler & Takma Adlar (Aliases)
"Özel sebep" varsa boyutlu tip kullan:

| Tip | Takma adı | Ne zaman |
|---|---|---|
| `byte` | = `uint8` (0-255) | Ham veri, `[]byte`, ağ paketi, shellcode, kriptografi |
| `rune` | = `int32` | Tek Unicode karakter (Türkçe, emoji, UTF-8). Java'daki `char` gibi |
| `int8`/`int16` | — | Bellek optimizasyonu (örn. HP değeri 255'i geçmiyorsa) |

> **Güvenlik notu:** TCP/UDP paketleyici yazarken veya buffer'dan ham veri okurken veriler bellekte byte byte dizilir → `[]byte` zorunlu.

---

## 7. Tip Dönüşümü: Explicit (Açık) Conversion

C **örtülü (implicit)** dönüşüm yapar, Go **yapmaz**. Go'da her dönüşüm `T(v)` formülüyle **elle** yazılır.

```go
// C: int -> float sessizce olur (veri kaybı riski)
// Go:
can := 100
gercekCan := float64(can)   // ✅ açıkça dönüştür
// gercekCan := can          ❌ HATA — Go otomatik çevirmez
```

`int(v)` dönüşümü ondalığı **yuvarlamaz, direkt keser (truncate)**:

```go
var konumX float64 = 105.89
pixelX := int(konumX)   // 105 (0.89 atılır)
```

> **Ders:** Go'nun bu katılığı bug önler. C'de o `0.89` sessizce kaybolur, sen "karakter neden duvara giriyor?" diye günlerce bug ararsın. Go seni açıkça dönüştürmeye zorlar.

---

## 8. Tip Çıkarımı (Type Inference)

### Sağ tarafın tipi belliyse → kopyalar
```go
var i int = 100
j := i        // j de int (i'nin tipini alır)
```

### Ham sabitse → varsayılan (default) tip
Sağda sadece ham sayı varsa Go varsayılan tipi atar:

| Değer | Varsayılan tip |
|---|---|
| `42` (tam sayı) | `int` |
| `3.142` (ondalık) | `float64` (hassasiyet kaybını önlemek için en geniş) |
| `0.8 + 0.5i` (karmaşık) | `complex128` |

> **Ders:** `:=` her zaman en optimize tip değildir. Değer asla 255'i geçmeyecekse `hiz := 120` sana koca bir `int` (64-bit) verir; `var optimizeHiz uint8 = 120` sadece 8-bit yer kaplar. Bellek kritikse kontrolü ele al.

---

## 9. Sabitler (`const`)

Program çalışırken **asla değişmeyen** değerler. `var` yerine `const` ile tanımlanır, derleme anında (compile-time) koda gömülür.

**Kurallar:**
1. Sonradan **değiştirilemez** (`OyunAdi = "x"` → HATA).
2. Sadece **basit tipler**: character, string, boolean, numeric. `slice`, `map`, `struct` **const olamaz**.
3. `:=` **kullanılamaz** (o değişkenlere özel bir kısayol; sabit değişken değildir).

```go
const HedefPort int = 443
const TaramaModuAktif = true

func main() {
    const ZamanAsimi = 30   // fonksiyon içinde de const olur
    // ZamanAsimi := 30      ❌ HATA — const'ta := yok
}
```

### Sayısal Sabitler & Yüksek Hassasiyet
Go derleyicisinin sabit matematik motoru çok güçlüdür — sabitler devasa/yüksek hassasiyette tutulabilir. Ama bu sabiti standart bir `int` bekleyen fonksiyona verirsen **taşar (overflow)** ve **derlenmeden** hata verir (C/Java runtime'da çökerdi).

> **Ders:** Sabit çok büyükse ve `int`e (max 64-bit) sığmıyorsa Go compile-time'da yakalar → runtime sürprizi olmaz.

### Sabite atama denemesi → derleme hatası
`MaxPort = 100` yazarsan Go şu hatayı verir (runtime'da değil, **derlenmez**):
```
cannot assign to MaxPort (neither addressable nor a map index expression)
```
Sebep: const bir bellek hücresi değildir; değeri derleme anında doğrudan koda gömülür, o yüzden "atanacak bir yer" yoktur.

### Neden `const` ile `:=` kullanılamaz? (3 sebep)
1. **`:=` = değişken kısayolu.** `:=`, arka planda `var` + tip çıkarımıdır (`x := 10` → `var x int = 10`). `var` doğası gereği **değiştirilebilir** bir bellek alanıdır; `const` ise **değişmez**. Değişken üretmek için tasarlanmış kısayolu sabit için kullanmak mantıksal çelişkidir.
2. **Compile-time vs run-time.** `const` derleme aşamasında kesinleşip koda gömülür. `:=` ise fonksiyon içinde, program çalışırken (run-time) hafızada yer açılan yerel değişkenler içindir. Derleyici, değerin kesin/değişmez olduğunu garanti altına almak için açıkça `const` görmek ister.
3. **Okunabilirlik.** `HedefPort := 443` satırına bakan biri bunun değişken mi sabit mi olduğunu anlayamaz. Go seni `const HedefPort = 443` yazmaya zorlayarak belirsizliği kaldırır.

> **Kısa özet:**
> - `:=` → "yeni bir **değişken** oluştur, değerini sonra değiştirebilirim"
> - `const` → "bu değer program kapanana kadar **asla değişmeyecek**, betona kaz"

---

## Faz 1 — Öğrenilenler Özeti
- [x] Exported/unexported (ilk harf kuralı)
- [x] `fmt`, `:=`, `var`, type-last, zero value
- [x] Scope (paket vs fonksiyon), satır başı keyword kuralı
- [x] Named return + naked return
- [x] Basic types, `byte`/`rune` alias, mimariye bağlı `int`
- [x] Explicit conversion `T(v)` (Go otomatik çevirmez)
- [x] Type inference & default tipler
- [x] `const` kuralları + overflow compile-time yakalama
