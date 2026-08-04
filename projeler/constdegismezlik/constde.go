package main

const MaxPort = 65535

func main() {
	const zamanAsimi = 30
	// Max port const olarak tanımlandıgı ıcın alındıgı bellek alaından bır daha degıstırılemez degerı
	/*1. := Operatörü Sadece Değişkenler (Variables) İçindir

	:= sembolü aslında Go'da var kelimesinin ve tip çıkarımının birleştirilmiş bir kısayoludur. Yani sen x := 10 yazdığında derleyici arka planda bunu var x int = 10 olarak okur.

	var (variable - değişken) kelimesi, doğası gereği sonradan değiştirilebilen bir bellek alanını ifade eder. Sabitler (const) ise asla değiştirilemez. Bu yüzden, değişken oluşturmak için tasarlanmış bir kısayolu, sabit oluşturmak için kullanamazsın. Mantıksal bir çelişki oluşturur.
	2. Derleme Zamanı (Compile-Time) vs. Çalışma Zamanı (Run-Time)

	Sabitler (const), program daha çalışmadan önce, derleme aşamasında (compile-time) kesinleşir ve uygulamanın içine doğrudan gömülür.

	:= operatörü ise fonksiyonların içinde, program çalışırken (run-time) dinamik olarak hafızada yer açılan yerel değişkenler için kullanılır. Go derleyicisi, derleme aşamasında o değerin kesin ve değişmez bir sabit olduğunu garanti altına almak için açıkça const anahtar kelimesini görmek zorundadır.
	3. Kod Okunabilirliği ve Kesinlik

	Go, kodu okuyan kişinin ne olduğunu anında anlamasını hedefleyen, belirsizlikleri sevmeyen bir dildir.

	Eğer const için de := kullanabilseydik, kodu okuyan birisi şu satıra baktığında:
	HedefPort := 443
	Bunun ileride değiştirilebilecek bir değişken mi yoksa asla değişmeyecek bir sabit mi olduğunu ilk bakışta anlayamazdı. Go, seni const HedefPort = 443 yazmaya zorlayarak bu kafa karışıklığını tamamen ortadan kaldırır.

	Kısa Özet:

	    := demek: "Bana yeni bir değişken (var) oluştur, ileride değerini değiştirebilirim."

	    const demek: "Bu değer program kapanana kadar asla değişmeyecek, bunu betona kazı."*/
}
