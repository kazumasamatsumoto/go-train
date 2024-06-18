package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// このエラーは、const変数に対して値を再割り当てしようとしたために発生しています。const変数は一度定義されたら値を変更することができない定数です。そのため、以下のコードでxとyに新しい値を割り当てようとするとエラーが発生します。

	// x = x + 1
	// y = "bye"

	// fmt.Println(x)
	// fmt.Println(y)
}
