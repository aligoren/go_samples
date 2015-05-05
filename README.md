<!--
.. title: Merhaba Go
.. slug: merhaba-go
.. date: 2015-05-05 21:32:27 UTC+03:00
.. tags: go, golang, programlama, programming
.. category: go, golang, programlama, programming
.. description: Golang Introduction. Go Dili
-->

Bir süredir Go dili ile ilgili makaleler okuyorum. Aslında ben de yeni yeni bu dili tam anlamıyla öğrenmekteyim. Okula başlamadan önce çalıştığım yerdeki arkadaşlarım bütün dilleri araştırmamdan şikayetçi oluyorlardı.

Bunları araştırmamda tabii ki sebep vardı. Nasıl işlerler ya da bu dilin tasarımcısı nasıl düşünür ne gibi yazılımcı ipuçları alabilirim diye merak ettim. Tek anlayamadığım dilse **type system** olayını hala daha anlayamadığım Go dili oldu. Düşünce yapım hoş olmayabilir ama bir şeyi bilmiyor ya da anlamıyorsan onu öğrenerek bu eksikliğini giderebilirsin diye düşündüm. Özellikle hala daha **type system** olayında karmaşık düşüncelere sahibim. Ekşi Sözlük'te dil hakkındaki bir entry'de [Go Object Oriented Design](http://nathany.com/good/) makalesini gördüm. Kafamı karıştıran **type system** olayı hakkında açıklamaları var. Dil nasıl ortaya çıkmış yine bunu dil geliştiricisi **Rob Pike** amca [Less is exponentially more](http://commandcenter.blogspot.com.tr/2012/06/less-is-exponentially-more.html) yazısında kaleme almış. <!-- TEASER_END -->

Go üzerinde hala anlamadığım olaylar var. Mesela paket yönetim olayı. Kendi bilgisayarınızda yazılan kütüphaneleri linkleme olayı sorun değil. Yani bir projede **N** kadar **.go** dosyası ile çalışın bunların hepsini çağırabilirsiniz. Fakat bu dil ayrıca github üzerinden direkt olarak yayınlanmış olan kütüphaneleri link olarak çağırma fırsatı da sunmakta. Burada nasıl çalıştığını ve neden böyle bir yönteme başvurduğunu anlamadım açıkcası.

#Syntax'a Söylenerek Merhaba Dünya

Örnek bir merhaba olayını kontrol ederken dahi zorlandım aslında. Örneğin şu go kodunu yorumlayalım:

~~~~{.go}
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
~~~~

Bu go kodunda yukarıdaki package kısmını belirtmediyseniz yani bu bir main paketidir demediyseniz ne run ne de build edebilirsiniz. Yeri gelmişken ondan da bahsedelim.

Go dilinde programları run ve build olmak üzere iki işleme tabi tutabilirsiniz. Run işlemi projeyi derleme yapmadan sadece çalıştırır. Örneğin;

`go run main.go`

Derleme yapmak içinse build komutunun olması gerekmekte. Yani;

`go build main.go`

Aslında dosyayı belirtmesek dahi main paketini arayacağı için yine derleme işlemi yapacaktır. Bu işlemleri Sublime ile daha basit hale getirebilirsiniz. Bu konu hakkında bir yazı [yazmıştım](http://aligoren.github.io/posts/sublime-text-golang-ide.html).

#Paketleri Çağırma ve Kullanılmaları

Devam edecek olursak paketleri çağırma olayı iki şekilde işlemekte. Yani multiple ve single olarak çağırma işlemi var. Çoğu ve tekil olarak olduğunu belirteyim. Yukarıdaki Hello World kodunda tek bir paket **import** ile çağırılmış durumda. Örnek bir paket şöyle çağırılır;

~~~~{.go}
import "fmt"
~~~~

Bu fmt paketi, C'nin stdio paketi ile aynı işlevi görmekte. Bu paket üzerinden stdin/out işlemleri yapılmakta. Çoklu paket çağırımına gelecek olursak eğer;

~~~~{.go}
import (
    "fmt"
    "paket/kutuphanesi"
)
~~~~

#Syntax Hatalarını Görmek

Çoklu kütüphaneleri böyle import edebilirsiniz. Fakat burada dikkat etmenizi istiyorum. Yukarıdaki import işleminde parantezlerin nerede olduğunun pek bir önemi olmasa da işlevlerde çok önemli bir yere sahipler. Örneğin şu kullanım yanlıştır, derleme anında hata verir.

~~~~{.go}
func main()
{
    fmt.Println("Hello World")
}
~~~~

***Programın verdiği hata:** syntax error: unexpected semicolon or newline before {*

Buradan da anlıyoruz ki dil içinde bazı alanlarda syntax aşırı kısıtlayıcı özelliklere sahip. İlk süslü parantez asla yeni satırda başlayamaz. Mutlaka işlevin tanımlandığı alanla aynı satırda olmalı.

#Küçük-Büyük Harf Bağımlılığına Dikkat Edin

Ayrıca küçük-büyük harf bağımlılığı var. Örneğin `math.pi` ile `math.Pi` arasında fark vardır. Go dilinin işlevlerini anladıysanız doğru yazımın `math.Pi` olduğunu da anlamışsınızdır. Diğer türlü derleme esnasında hata alınır.


#Go Dilinin Basit Tipleri

Hazır math falan demişken tiplere de gelelim. Diğer dillerde olduğu gibi uint, int, float, string, bool gibi sayısal, metinsel ve karar tipleri var. Hatta karmaşık sayılar için complex dahi var. Bunlar da kendi aralarında **uint8, uint16, uint32, uint64, int8..64, rune(int32), float32, float64, complex64, complex128** şeklinde ayrılmaktalar. Burada `uint8 == byte` olarak söylenebilir. Ayrıca burada makina bağımlı üç sayı tipi de var. Bunlar da **uint, int, uintptr**.

#Tipler Üzerine Basit Bir Matematiksel İşlem

Basit olarak bir matematiksel işlem, şöyle ifade edilebilir;

~~~~{.go}
fmt.Println("1 + 1 =", 1+1) // integer sayılar
fmt.Println("1.0 + 1.0 =", 1.0+1.0) // float sayılar
~~~~

#Stringler, String Birleştirme ve Len

Tip belirtmeksizin bu şekilde basit bir işlemi gerçekleştirebilirsiniz. Stringler de yine aynı şekilde diğer dillerde yapılan birleştirme işlemleri gibi yapılmakta. Örneğin python'da var olan `len` burada da var.

~~~~{.go}
fmt.Println(len("Merhaba Dunya")) // string boyutunu belirtiyor :)
fmt.Println("Merhaba Dunya"[1]) // 101
fmt.Println("Merhaba " + "Dunya") // string birleştirme
~~~~

Stringlerin kullanımı bu şekilde basit olabilmekte. Yukarıdaki örnekte fark ettiğiniz bir şey vardır belki 101 olayı yani `[1]` bu o stringi bir dizi gibi görmekte. Dizinin ikinci elementini alıyor ve onu byte olarak tekrar tanımlıyor. Go dilinde, tekrardan `byte == integer` dememin bir sakıncası yok.

#Boolean Karar Tipi

Tiplerle ilgili son olarak `boolean` karar tiplerinden bahsetmek gerekli.

~~~~{.go}
fmt.Println("True && True =", true && true)
fmt.Println("True && False =", true && false)
fmt.Println("True || True =", true || true)
fmt.Println("True || False =", true || false)
fmt.Println("!true =", !true)
~~~~

kullanımı bu şekildedir. Boolean adını **George Boole** amcadan almıştır. Bir boolean değeri 1 bit özel sayı türü, true ya da false durumlarını temsil etmek için kullanılır. Üç mantıksal operatör, boolean değerleri ile birlikte kullanılır.

**Boolean Operatörleri:** [&&: and], [||: or], [!: not]

Yukarıdaki bool operatörleri diğer dillerle zaten aynı. Çok da yabancılık çekmeyeceğinizi düşünerek basit şekilde yazdım. Yukarıdaki kodu çalıştırarak sonucu görebilirsiniz. Eğer bilgisayarınızda Go kurulu değilse [ideone](https://ideone.com/) bu iş için ideal online compiler ve IDE görevini görmekte.

#Go Dilinde Değişkenler ve Scala

Değişkenleri tanımlarken aklıma scala gelmiyor değil. Çünkü scala'da değişken tanımlamaları go ile aynı şekilde. Örnek bir scala programında değişkenleri şöyle tanımlarsınız:

~~~~{.scala}
var degisken : String = "Degistirilebilir"
val iDegisken : String = "Degistirilemez"
~~~~

Yukarıdaki kod parçacığı scala'da şöyle çalışır `degisken` isimli değişken tanımlanmıştır. Daha sonra değiştirilebilir şekilde. `iDegisken` ise immutable değişken olarak tanımlanır. Yani değiştirilemez değişken. Yukarıdaki kod satırına neden değindim diye soracak olursanız söyleyeyim go dilinde de aynı kullanım var. Yani **var** ve **val** mantığı ile çalışan değişkenler var. Fakat tamamen aynı değiller. Şu kodlar üzerinden yorumlayalım:

~~~~{.go}
var x string = "Hello"
fmt.Println(x)

x = "World"

fmt.Println(x)

y := "Hi"

fmt.Println(y)

y = "Are you okey?"

fmt.Println(y)

n :=  1 // n := x mean => as x value(string, float, int vs.)

fmt.Println(n) // 1
~~~~

Yukarıdaki koda bakarsak eğer **x** adında bir değişkenin türünün string olarak belirtildiğini ve bu şekilde tanımlandığını görüyoruz. Bir alttaki değişkende ise yani **y** değişkeninde ise bir tip belirtilmemiş olmasına rağmen bu değişkenin de aynı işlemleri yaptığını görüyoruz. Eğer y değişkenine tekrar atama işlemini `y := "falan"` şeklinde yapmış olsaydık bir derleme hatası alacaktık. Değişken tanımlamaları şöyle yapılır **var DEĞİŞKEN_ADI değişken_veri_tipi = DEĞİŞKEN_İÇERİĞİ**

#İşlev Değişkenleri ve Global Özelliği Olan Değişkenler

Peki `:=` ve `var` arasındaki fark nedir. Neden böyle bir kullanım gerekir? Eğer işlev alanlarının yani bildiğiniz `func` dışında bir tanımlama yapacaksa `var` kullanmalıyız. Eğer işlev alanı dışında bir değişkni `:=` şeklinde tanımlamak istersek hata alırız.

Ayrıca dışarıdan erişimli değişkenler oluşturabiliriz. Bunu ana işlev dışında tanımlayıp, diğer işlevler için kullanıma açabiliriz. Örneğin

~~~~{.go}
var islevDisi string = "Bu islev disinda herkese acik"

func main() {
    fmt.Println(islevDisi)
}

func baskeIslev() {
    fmt.Println(islevDisi)
}
~~~~

Eğer **islevDisi** değişkeni main içerisinde olsaydı **baskaIslev** fonksiyonunda kullanılamayacaktı. Kapsamı genel olması itibariyle bu tip kullanımlar **scope** olarak da adlandırılırlar.

#Sabitler ve Değiştirilemez Oluşları

Sabitler yani const kullanımı ise diğer dillerdeki mantıkla hareket eder. Bu tipler daha sonradan değiştirilme şansı olmayan tiplerdir.

~~~~{.go}
const degisken string = "Ben degistirilemem"

degisken = "Degistirmeyi denerim" // derleme zamanı hatası!
~~~~

Görüldüğü gibi değer bu değişkenlere tekrar atamalar yapmaya kalkarsak derleme zamanında hata alırız. Burada yapılan şey `var` anahtar kelimesini yerini `const` anahtar kelimesine bırakmıştır.

#Çoklu Değişken ve Çoklu Sabit Tanımlama

Diğer dillerde olduğu gibi burada da çoklu değişken tanımlama yapılabilir. Diğer dillerde bu özellik var mı bilmiyorum ama Go şöyle bir tanımlama işlemini destekler:

~~~~{.go}
var (
    a = 1
    b = 2
    text = "ben metinsel bir degiskenim bunlarsa sayisal arkadaslarim: "
)

fmt.Println(text, a, b)
~~~~

Çoklu değişken tanımlayıp tek seferde hepsini yazdırabilirsiniz. Yine sabitleri de çoklu olarak bu şekilde tanımlayabilirsiniz.

#Kontrol Yapıları: For Döngüsü

Kontrol yapıları dil açısından şu anlık söylemek istediğim son şey. Zaten yazı da epey bir uzadı. Diğer dillerde kullanılan kontrol yapıları burada da aynı şekilde kullanılmakta. Örneğin 1'den 10'a kadar giden sayıları for döngüsü kullanarak yazdıralım:

~~~~{.go}
i := 1
for i <= 10 {
    fmt.Println(i)
    i = i + 1
}
~~~~

Yukarıdaki kodda `i = i + 1` kullanımı olmasaydı ne olurdu programı kendiniz derlemeye çalışarak sonucu görün. Yukarıdaki for döngüsü bir kullanım çeşididir. Ayrıca normal en bilindik yöntemi ile de şöyle kullanılmakta:

~~~~{.go}
for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
~~~~

#Kontrol Yapıları: if-else ve switch-case

Çok tanıdık değil mi diğer dillerden bildiğiniz şekli. Hangisi size daha basit gelirse. Kontrol yapılarında son olarak belirtmek istediğim **if-else** ve **switch-case**

#Kontrol Yapıları: if-else

Diğer dillerde olduğu gibi bu dilde de bu yapıların mantıkları aynıdır. Yukarıdaki for döngüsünü tekrar kurmayacağım ama aşağıdaki kodların o döngü içerisinde olduğunu varsayalım:

~~~~{.go}
if i == 0 {
    fmt.Println("Sifir")
} else if i == 1 {
    fmt.Println("Bir")
}
~~~~

Böylece artık for döngüsünden gelen sayıları metinsel hale getirebilirsiniz. Çok uzadığı için yapıları kısa tuttum.

#Kontrol Yapıları: switch-case

Yine **switch-case** için de kullanım aynı şekilde. Yine bu kullanılan kodların da for döngüsü içerisinde olduğunu varsayalım:

~~~~{.go}
switch i {
    case 0: fmt.Println("Sifir")
    case 1: fmt.Println("Bir")
    default: fmt.Println("Bu sayiyi bilmiyorum")
}
~~~~

Kullanımı bu şekildedir. Yine diğer dillerde olduğu gibi `default` aranan koşulların sağlanmaması durumunda kendisini gösterir.

Anlatım bu kadardır.

**Yazı Kaynağı:** [http://aligoren.github.io/posts/merhaba-go.html](http://aligoren.github.io/posts/merhaba-go.html)
