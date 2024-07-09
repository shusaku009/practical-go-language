package main

import (
	"fmt"
	"practical-go-language/sub"
)

func main() {
	kakeUdon := sub.NewKakeUdon(sub.Large)
	kitsuneUdon := sub.NewKitsuneUdon(sub.Regular)
	tempuraUdon := sub.NewTempuraUdon(sub.Small)

	// 構造体を利用したオプション引数
	optionUdon := sub.NewUdonFn(sub.Option{
		Men:      sub.Regular,
		Aburaage: true,
		Ebiten:   2,
	})

	// ビルダーを利用したオプション引数
	fluentUdon := sub.NewUdonOp(sub.Large).Aburaage().Ebiten(3).Order()

	// Functional Optionパターンを使ったオプション引数
	funcOptionUdon := sub.NewUdonWithOpts(
		sub.OptMen(sub.Large),
		sub.OptAburaage(),
		sub.OptEbiten(2),
	)

	// 作成したUdonの情報を表示
	fmt.Printf("Kake Udon: %+v\n", kakeUdon)
	fmt.Printf("Kitsune Udon: %+v\n", kitsuneUdon)
	fmt.Printf("Tempura Udon: %+v\n", tempuraUdon)
	fmt.Printf("Option Udon: %+v\n", optionUdon)
	fmt.Printf("Fluent Udon: %+v\n", fluentUdon)
	fmt.Printf("Func Option Udon: %+v\n", funcOptionUdon)

	result := sub.Strings()
	fmt.Println(result)

	bresult := sub.BuilderStrings()
	fmt.Println(bresult)

	timeresult := sub.TimeNow()
	fmt.Println(timeresult)

	timeduration := sub.TimeDuration()
	fmt.Println(timeduration)
}
