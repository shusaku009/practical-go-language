package sub

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	Men      Portion
	Aburaage bool
	Ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		Men:      p,
		Aburaage: aburaage,
		Ebiten:   ebiten,
	}
}

var tempuraUdon = NewUdon(Large, false, 2)

// 別名の関数によるオプション引数
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		Men:      p,
		Aburaage: false,
		Ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		Men:      p,
		Aburaage: true,
		Ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		Men:      p,
		Aburaage: false,
		Ebiten:   3,
	}
}

var kakeUdon = NewKakeUdon(Large)

type Option struct {
	Men      Portion
	Aburaage bool
	Ebiten   uint
}

// 構造体を利用したオプション引数
func NewUdonFn(opt Option) *Udon {
	if opt.Ebiten == 0 && time.Now().Hour() < 10 {
		opt.Ebiten = 1
	}
	return &Udon{
		Men:      opt.Men,
		Aburaage: opt.Aburaage,
		Ebiten:   opt.Ebiten,
	}
}

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdonOp(p Portion) *fluentOpt {
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

// ビルダーを利用したオプション引数
func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		Men:      o.men,
		Aburaage: o.aburaage,
		Ebiten:   o.ebiten,
	}
}

// Functional Optionパターンを使ったオプション引数
type OptFunc func(r *Udon)

func NewUdonWithOpts(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.Men = p }
}

func OptAburaage() OptFunc {
	return func(r *Udon) { r.Aburaage = true }
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) { r.Ebiten = n }
}

var (
	FlagStr = flag.String("string", "default", "文字列フラグ")
	FlagInt = flag.Int("int", -1, "数値フラグ")
)

func Strings() string {
	src := []string{"Back", "To", "The", "Future", "Part", "III"}
	var title string
	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Println(title)
	return title
}

func BuilderStrings() string {
	src := []string{"Back", "To", "The", "Future", "Part", "III"}

	var builder strings.Builder
	builder.Grow(100)
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
	return builder.String()
}

func TimeNow() string {
	now := time.Now()
	tz, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("エラー：タイムゾーンの読み込みに失敗しました:", err)
		return now.String()
	}
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)
	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))
	return now.String()
}
