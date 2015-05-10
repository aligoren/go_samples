<!--
.. title: Merhaba Go
.. slug: merhaba-go
.. date: 2015-05-05 21:32:27 UTC+03:00
.. tags: go, golang, programlama, programming
.. category: go, golang, programlama, programming
.. description: Golang Introduction. Go Dili
-->

Bir süredir Go dili ile ilgili makaleler okuyorum. Aslında ben de yeni yeni bu dili tam anlamıyla öğrenmekteyim. Okula başlamadan önce çalıştığım yerdeki arkadaşlarım bütün dilleri araştırmamdan şikayetçi oluyorlardı.

Bunları araştırmamda tabii ki sebep vardı. Nasıl işlerler ya da bu dilin tasarımcısı nasıl düşünür ne gibi yazılımcı ipuçları alabilirim diye merak ettim. Tek anlayamadığım dilse **type system** olayını hala daha anlayamadığım Go dili oldu. Düşünce yapım hoş olmayabilir ama bir şeyi bilmiyor ya da anlamıyorsan onu öğrenerek bu eksikliğini giderebilirsin diye düşündüm. Özellikle hala daha **type system** olayında karmaşık düşüncelere sahibim. Ekşi Sözlük'te dil hakkındaki bir entry'de [Go Object Oriented Design](http://nathany.com/good/) makalesini gördüm. Kafamı karıştıran **type system** olayı hakkında açıklamaları var. Dil nasıl ortaya çıkmış yine bunu dil geliştiricisi **Rob Pike** amca [Less is exponentially more](http://commandcenter.blogspot.com.tr/2012/06/less-is-exponentially-more.html) yazısında kaleme almış.

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

Go dilinde 3 basit tip vardır. Bu yazıda bu üç built-in tipten bahsedeceğim. Bu tipler: arrays, slices ve maps'dir.

Diğer dillerde olan map, array ve slices kavramı bu dilde de var. Tanımlanışları ve kullanışları itibariyle farklılıklar gösterseler de çalışma mantıkları aynı.

#Arrays

Örneğin go dilinde bir array şöyle tanımlanmakta ve şöyle değer ataması yapılmaktadır:

~~~~{.go}
var x [5]int

x[4] = 10
~~~~

Tanımlama işlemi yaparken yine **var** anahtar kelimesini kullanım dizi değişkeninin adını yazıp dizi boyutunu ve türünü belirtmemiz gerekiyor. Nasıl derlenir, ne yapılır işine girmiyorum daha önce [bahsetmiştim](http://aligoren.github.io/posts/merhaba-go.html).

Yukarıdaki kodda `[5]int` aslında `0,1,2,3,4` demektir. Yani dizilere aşinalığınız varsa dizilerin ilk değerinin sıfır olduğunu biliyorsunuz demektir.

Ayrıca bir önceki yazımda bahsettiğim işlev değişkenleri mantığını burada da uygulayabiliriz. Dizileri farklı olarak şöyle de tanımlayabiliriz:

~~~~{.go}
x := [5]float64{10, 15, 22, 35, 56}
~~~~

şeklinde tanımlayabildiğimiz gibi yukarıdaki tanımlamayı şöyle de yapabiliriz:

~~~~{.go}
x := [5]float64{
    10,
    15,
    22,
    35,
    56,
}
~~~~

Dizileri iterasyona uğratma işlemini şöyle yapıyoruz. Yukarıdaki dizimiz üzerinden bir örnek geliştirelim:

~~~~{.go}
var total float64 = 0

for i := 0; i < 5; i++ {
    total += x[i]
}

fmt.Println(total / 5)
~~~~

Yukarıda dizi boyutunu kendimiz manuel olarak belirledik. Yani `total / 5` işlemindeki 5 bizim dizi boyutumuzdu. Bunu biraz daha otomatik hale getirelim:

~~~~{.go}
for i := 0; i < len(x); i++ {
    total += x[i]
}

fmt.Println(total / len(x))
~~~~

Yukarıdaki kod bir üstteki kodun aynısı, çalışma şekli ve getireceği sonuçlar aynı. Ancak çalıştırıp denediğimizde `invalid operation: total / 5 (mismatched types float64 and int)` hatasını alırız. Bu hata dönüşüm hatasıdır. Bu hatadan kurtulmak için float64 dönüşümü yapalım. Unutmadan float64 dönüşümü her zaman olmamalıdır tiplerin kendilerine göre dönüşümleri vardır. Artık fixlenmiş hali şöyledir:

~~~~{.go}
fmt.Println(total / float64(len(x)))
~~~~

Artık bu kod çalışır hale geldi. Ayrıca dizileri C# ve diğer bazı dillerde olan **foreach** mantığı ile kullanabiliriz. Bunu şöyle yaparız:

~~~~{.go}
for i, val := range x {
    total += val
}
~~~~

Ancak bu kod için bir hata var. Söyledim mi bilmiyorum ancak go dilinde tanımlanan bir değişken, çağırılan bir kütüphane kullanılmazsa derleme hatası verir. Yukarıdaki kod için söyleyecek olursak **i** değişkeni kullanılmamaktadır. Bunu `i declared and not used` hatasından anlayabiliriz. Bu durumdan kurtulmak için yukarıdaki kodu şu şekilde değiştirelim:

~~~~{.go}
for _, val := range x {
    total += val
}
~~~~

Yukarıdaki kodu inceleyelim. Yukarıdaki kod **i** değişkenini kaybetmiş ve onun yerine **_** yani underscore(alt çizgi) almıştır. Peki neden alt çizgi kullanırız?

Alt çizgi, derleyiciye **bu değere ihtiyacım yok ama dilin işleyişi açısından da gerekli** mesajını iletir. Bu değişken aslında bir yineleyicidir. Kullanılmayacak değişkenleri underscore(alt çizgi) ile belirtebiliriz tabii bu değişkenler aslında ihtiyaç olmayan ama orada bulunması gereken değişkenlerdir. Bunu unutmadan ihtiyaca göre kullanmamız gerektiğini bilelim.

#Slices

Slices yani dilimler dizilerin birer parçasıdır. Dilimlerin kullanılışı da diziler gibidir. Belirli uzunlukları vardır. Dizilerin aksine dilimlerde uzunluklar değiştirilebilirler.

Bir dilim şöyle tanımlanabilir:

`var x []float64`

Burada boyutu sıfır olan bir dilim tanımlaması yaptık. Eğer kullanılması gereken bir dilim oluşturmak istiyorsak `make` built-in fonksiyonunu kullanarak bunu yapabiliriz:

~~~~{.go}
x := make([]float64, 5)
~~~~

Burada uzunluğu 5 olan bir array ile ilişkilendirilmiş float64 tipinde bir dilim oluştururuz.

Dilimler diziler ile ilişkili demiştim örneğin uzunluğu 5 olan bir dilim daha küçük de olabilir değiştirilebilirlikleri vardır. Ayrıca `make` fonksiyonu dilimin kapasitesini belirten üçüncü bir değer de alabilir.

~~~~{.go}
x := make([]float64, 5, 10) // benim kapasitem 10
~~~~

Bir diziyi şöyle tanımlayabiliriz demiştik:

~~~~{.go}
arr := [5]float64{1,2,3,4,5}
~~~~

Bu dizileri ayrıca `[low:high]` ifadesi ile kullanabiliriz. Bunu şöyle yaparız:

~~~~{.go}
x := arr[0:5]

fmt.Println(x) // [1 2 3 4 5]
~~~~

Aslında kendisi burada bir iterasyon işlemi yapmış oluyor.

#Slices Functions

Dilimlerin bazı özel fonksiyonları da vardır. Bunlardan birisi yine diğer dillerden biliyorsanız eğer `append` fonksiyonudur. Bu fonksiyonla dizilere ekleme yapabiliriz. Bir örnekle anlamaya çalışalım:

~~~~{.go}
sl1 := []int{1,2,3}
sl2 := append(sl1, 4,5)
~~~~

Yukarıdaki örnek kodda **sl1** dilimi `1,2,3` değerlerini almıştır. Daha sonra bu dilimlere ekleme yapmak istersek yukarıdaki `append(sl1, 4,5)` fonksiyonunu kullanmalıyız. Bu fonksiyonun ilk değeri var olan bir diziyi temsil eder. İkincil ve sonraki gelen değerler ise soldaki yani var olan diziye eklenecek değerleri temsil eder. Dizilerde ayrıca kopyalama işlemi de yapabiliriz.

Örnekle açıklamak gerekirse şöyle yapalım:

~~~~{.go}
sl1 := []int{1,2,3}
sl2 := make([]int, 2)

copy(sl2, sl1)

fmt.Println(sl2, sl1) // [1 2] [1 2 3]
~~~~

Yukarıdaki kodun ilk 2 satırını bildiğimizi varsayarak diğer satırdan yani `copy()` işlevinin olduğu satırdan bahsedelim.

Bu işlev iki değeri alır. İlk değer içerisine değer kopyalanacak olan değerdir. İkinci değer ise ilgili verinin adresini belirtir. Yukarıdaki örnekte **sl2** dilimi **sl1** diliminin ilk 2 değerini almıştır. Çünkü boyutu 2 olarak belirlendi. Yani sl1'in ilk 2 değerini kopyalamış olduk.

#Maps

Haritalar key-value çiftin, unordered collection yani sırasız bir koleksiyonudur. Bu Türkçe format için özür dilerim. key-valu ve unordered collection için aklıma uygun bir şey gelmedi. Haritaların basit şekilde tanımlanışı şöyledir:

~~~~{.go}
var x map[string]int
~~~~

İlk defa bu makalede gördüğümüz `map` keywordünü kullandığımızı fark etmişsinizdir. Tanımlanışları olarak ise **map[ANAHTAR_TIPI]DEGER_TURU**
şeklinde tanımlanırlar.

Hadi daha somut bir örnek kullanarak bunu açıklayalım:

~~~~{.go}
var x map[string]int

x["key"] = 10

fmt.Println(x)
~~~~

Yukarıdaki kod anahtar tipi `"key"` ve değer türü int olan bir map'i göstermektedir. Yine o kodda `fmt.Println` ile o değeri yazdırmaya çalışıyoruz. Ancak bu şekilde kullanımda `panic: runtime error: assignment to entry in nil map` hatasını alırız. Bu hata harita için bir atama olması gerektiğini belirtir. Bunu şöyle düzeltelim:

~~~~{.go}
fmt.Println(x["key"])
~~~~

Buradan da anlayacağımız gibi `x["key"], x["val"], x["falan"]` gibi farklı atamalar yapabiliriz. Bu atamaları çağırırken anahtar değerleri ile birlikte çağırmamız gerekiyor. Bir harita içerisindeki anahtar değer ve içeriğini `delete()` işlevi ile sileriz.

Haritaların kullanımıyla ilgili şu küçük uygulamayı örnek alabiliriz:

~~~~{.go}
package main

import "fmt"

func main() {
    lang := make(map[string]string)

    lang["PHP"] = "PHP: Hypertext Preprocessor"
    lang["HTML"] = "HyperText Markup Language"
    lang["ASP"] = "Active Server Pages"
    lang["JS"] = "Javascript"
    lang["CPP"] = "C++: C plus plus"

    fmt.Println(lang["PHP"])
}
~~~~

Bu programı geliştirip çalıştırın ve ne sonuç verdiğini görün. Yukarıdaki programın map bölümünü baştan şöle kodlayabiliriz:

~~~~{.go}
lang := map[string]string {
    "PHP": "PHP: Hypertext Preprocessor",
    "HTML": "HyperText Markup Language",
    "ASP": "Active Server Pages",
    "JS": "Javascript"
    "CPP": "C++: C plus plus",
}
~~~~

Bu kodun yukarıdaki koddan bir farkı yoktur. Yukarıdaki kodda `make` gömülü fonksiyonunu kullandık sadece. Son olarak haritalarla ilgili şu örneği vermek istiyorum.

Elementler ve bu elementlerin hallerini gösteren bir uygulama olduğunu varsayalım:

~~~~{.go}
elements := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen", 
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium", 
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium", 
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium", 
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }
~~~~

Haritalar da genellikle bilgi depolamak için kullanılırlar. Bunu biliyoruz. Yukarıdaki örnekte elementleri ve bu elementlerin hallerini sakladığımızı görmekteyiz. Haritanın gösterim şeklinin `map[string]string` olmadığını artık yeni halinin `map[string]map[string]string` olduğunu görmekteyiz.

Bu kullanım bize haritalar içinde de haritalar olabileceğini göstermektedir. Yukarıdaki haritaya bakarak Helium'un **state** değerini yani hal değerini yazdıralım:

~~~~{.go}
fmt.Println(elements["He"]["state"]) // gas
~~~~

Eğer tür belirtmeden bir state değeri yazdırmak isteseydik

~~~~{.go}
fmt.Println(elements["name"]["gas"]) // nil
~~~~

Bize nil bir değer döndürecekti.

Sanırım bu yazının sonuna gelmiş bulunmaktayız. İyi çalışmalar.
