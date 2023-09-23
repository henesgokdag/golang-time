package main

import (
	"fmt"
	"time"
)

func main() {
	addOneMonth()
	calculateWithoutSeconds()
	parseTimeWithSeconds()
}

//Bir zaman anı için time.Time kullan
//Zaman anları ile uğraşırken ve karşılatırma, ekleme, çıkarma işlemleri için time.Time metodlarını kullanın.

// yanlış
func isActiveWrong(now, start, stop int64) bool {
	isActive := start <= now && now < stop
	println(isActive)
	return isActive
}

// doğru
func isActiveCorrect(now, start, stop time.Time) bool {
	isActive := (start.Before(now) || start.Equal(now)) && now.Before(stop)
	println(isActive)
	return isActive
}

// Süreler ile uğraşırken time.Duration kullanın.
// int şeklinde aldığımızda method duration'ı saniye mi millisaniye mi alıyor diye düşünüp içine girip okumanız gerekebilir.
// Ayrıca direkt duration almak kodun okunabilirliğini de arttıracaktır.

// yanlış
func AddDelayWrong(delay int) {
	for {
		// ...
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// doğru
func AddDelayCorrect(delay time.Duration) {
	for {
		// ...
		time.Sleep(delay)
	}
}

func addOneMonth() {

	//eğer zamanın bir ay sonrasını almak istiyorsak
	start := time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC)

	//bir sonraki ayın takvim gününde aynı zamanını istiyorsak Time.AddDate kullanmalıyız.
	newDay := start.AddDate(0 /* years */, 1 /* months */, 0 /* days */)

	//Ancak bir anlık zamanın 30 gün sonraki zamanının  olmasını garantilemek istiyorsak Time.Add kullanmalıyız.
	maybeNewDay := start.Add(30 * 24 * time.Hour)

	fmt.Println(newDay)
	fmt.Println(maybeNewDay)
	fmt.Println(newDay.Equal(maybeNewDay))
}

/*
Dış sistemler için time.Time ve time.Duration kullan
Mümkün olduğunda, dış sistemler ile etkileşim için time.Duration ve time.Time kullanın. Örnek olarak:

Komut-satırı flag'leri: flag, time.ParseDuration aracılığıyla time.Duration'ı destekler.
JSON: encoding/json, bir RFC 3339 string'i olarak UnmarshalJSON metodu ile time.Time'ı dönüştürmeyi destekler.
SQL: database/sql, DATETIME veya TIMESTAMP sütunlarını time.Time'a dönüştürmeyi ve temel sürücü destekliyorsa geri döndürmeyi destekler
YAML: gopkg.in/yaml.v2, RFC 3339 string'i olarak time.Time'ı ve time.ParseDuration aracılığıyla time.Duration'ı destekler.
Bu etkileşimlerde time.Duration'ı kullanmak mümkün olmadığında, int veya float64 kullanın ve alan adına birimi ekleyin.

Örneğin, encoding/json, time.Duration öğesini desteklemediğinden, birim alan adına dahil edilir.
*/

// {"interval": 2}
type ConfigWrong struct {
	Interval int `json:"interval"`
}

// {"intervalMillis": 2000}
type ConfigCorrect struct {
	IntervalMillis int `json:"intervalMillis"`
}

/*
Eğer time.Time kullanamadığınız durumlar söz konusuysa string ve datetime'ı RFC 3339 olarak formatlayabilirsiniz.
Bu format Time.UnmarshalText de varsayılan olarak kullanılır ve time.RFC3339 aracılığıyla Time.Format ve time.Parse içinde de kullanılabilir.
*/

// Time Paketi hesaplamlarda artık saniyeleri hesaba katmaz.
func calculateWithoutSeconds() {
	format := "2006-01-02 15:04:05 -0700 MST"
	t1, _ := time.Parse(format, "2015-06-30 23:59:58 +0000 UTC")
	t2, _ := time.Parse(format, "2015-06-30 23:59:59 +0000 UTC")
	// 2015-06-30 23:59:60 tarihinde artık saniye vardır
	t4, _ := time.Parse(format, "2015-07-01 00:00:00 +0000 UTC")

	fmt.Println(t2.Sub(t1))
	fmt.Println(t4.Sub(t2))
}

// time.Parse methodu artık saniyeleri hesaba katmaz.
func parseTimeWithSeconds() {
	// 2005-12-31T23:59:60Z tarihinde artık saniye vardır.
	// http://datacenter.iers.org/eop/-/somos/5Rgv/getTX/16/bulletinc-030.txt

	t, err := time.Parse(time.RFC3339, "2005-12-31T23:59:60Z")
	if err != nil {
		fmt.Println("parse error:", err.Error())
	} else {
		fmt.Println("parsed as:", t.Format(time.RFC3339))
	}
}
