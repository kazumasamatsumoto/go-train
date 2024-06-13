package main

import "fmt"

func main() {
	fmt.Println("=====基本型======")
	fmt.Println("=====ゼロ値======")
	{
		var x int
		var y float32
		var z string

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		// Go言語では変数が初期化されていない場合、自動的にその型に対応するゼロ値が割り当てられます。
	}

	fmt.Println("=========リテラル==========")
	fmt.Println("=========整数リテラル==========")
	fmt.Println(1_234_453_32)

	fmt.Println("=========浮動小数点数リテラル==========")
	fmt.Println(1_234.345_32)

	fmt.Println("rune リテラル")
	{
		fmt.Println("1->\x41")
		fmt.Println("2->\u0041")
		fmt.Println("3->\U00000041")

		/**
		このコードは、Go言語でのruneリテラルを示しています。runeリテラルはUnicodeのコードポイントを表現するために使われます。具体的には、文字「A」のUnicode表現を異なる方法で表示しています。以下にそれぞれの方法について説明します：
		\x41
		\u0041
		\U00000041
		それぞれの形式は以下のように解釈されます：
		\x41 は、16進数のエスケープシーケンスであり、1バイト（2桁の16進数）で表現される文字を示します。ここでは、16進数の41が示す文字、つまり「A」です。
		\u0041 は、16進数のUnicodeエスケープシーケンスであり、4桁の16進数で表現されるUnicodeコードポイントを示します。ここでも、16進数の0041が示す文字、つまり「A」です。
		\U00000041 は、32ビットのUnicodeエスケープシーケンスであり、8桁の16進数で表現されるUnicodeコードポイントを示します。ここでも、16進数の00000041が示す文字、つまり「A」です。
		*
		*/
		fmt.Println("1->\u3042")
		fmt.Println("2->\U00003042")

		fmt.Println("1->\U0001F600")
		fmt.Println("\n2.1.2.4文字列リテラル")
		x := `バッククォートを使って”ロー文字列リテラル"を書くことで改行や二重引用符（ダブルクォート）を文字列中に含めることが可能です。`
		fmt.Println(x)
	}

	fmt.Println("\n2.1.2.5リテラルと型")
	{
		var x float32 = 2 * 0.23 * 3
		fmt.Println(x)
		var a float64 = 3.14
		var b float64 = 10 / a
		fmt.Println(b)

		// var c int = "abc" エラーになる
		// var s string = 12 エラーになる
		var d int
		// d = 12.3 // コンパイル時のエラー （桁落ちする）
		d = 12.0             // エラーにならない（桁落ちしない）
		fmt.Println("d:", d) // d: 12
	}

	fmt.Println("===== 2.1.3 論理型 =====")
	{
		var flag bool
		var isAwesome = true
		fmt.Println(flag)
		fmt.Println(isAwesome)
	}

	fmt.Println("===== あふれる =====")
	{
		var x byte = 100
		// ./ex0200.go:78:16: cannot use 1000 (untyped int constant) as byte value in variable declaration (overflows)
		fmt.Println(x)
	}

	fmt.Println("===== 2.1.4.4整数関連の演算子 =====")
	{
		var x int = 10
		x *= 2
		fmt.Println(x) // 20
	}
	fmt.Println("===== 2.1.4.5　浮動小数点数型 =====")
	{
		var x float64 = 0
		fmt.Println(x) // 0
		x += 1.333456
		fmt.Println(x) // 1.333456

		x = 1.5
		var y float64 = 2.5
		x /= y
		fmt.Println(x) // 0.6
		// 		x = x % y    // エラー。 floatにたいして「%」は使えない

		fmt.Println("x/0=", x/0)   // +Inf
		fmt.Println("-x/0=", -x/0) // -Inf
	}

	fmt.Println("===== 2.1.5　文字列とrune =====")
	{
		var a, b string
		a = "イロハ"
		b = "あいうえお"
		fmt.Println(a == b)                  // false
		fmt.Println(a != b)                  // true
		fmt.Println(`"あ" < "い":`, "あ" < "い") // true
		fmt.Println(`"ア" < "あ":`, "ア" < "あ") // false
		fmt.Println("a+b:", a+b)             // イロハあいうえお
	}
	fmt.Println("===== 2.2 変数の宣言 =====")
	{
		var x, y int = 10, 20
		fmt.Println(x, y) // 10 20
	}

	{
		var x, y int
		fmt.Println(x, y) // 0 0
	}

	{
		var x, y = 10, "hello"
		fmt.Println(x, y) // 10 hello
	}

	{
		var (
			x    int
			y        = 20
			z    int = 30
			d, e     = 40, "hello"
			f, g string
		)
		fmt.Println(x, y, z, d, e) // 0 20 30 40 hello

		fmt.Println("f=|", f, "| g=|", g, "|") // f=|  | g=|  |
	}

	{
		var x = 10
		fmt.Println(x) // 10
	}

	{
		x := 10
		fmt.Println(x) // 10
	}

	{
		var x = 10
		// var y int32 = x  // エラー
		var y int = x
		fmt.Println(y) // 10
	}
	{
		x := 10
		var y = 10
		var z int64 = 10

		fmt.Println(x == y)
		//	fmt.Println(x==z)  // エラー  x と z は型が異なる
		fmt.Println(int64(x) == z) // true
	}

	fmt.Println("===== 2.4型付きの定数と型のない定数 =====")
	{
		const x = 10
		var y int = x
		var z float64 = x
		var d byte = x

		fmt.Println("x, y, z, d:", x, y, z, d)

		const typedX int = 10
		fmt.Println("typedX:", typedX)
		fmt.Println(x == typedX) // true

		//		 z = typedX  // エラー
		z = float64(typedX)
		fmt.Println("z:", z)
	}

	{ // 2.5　使われない変数
		x := 10
		x = 20
		fmt.Println(x)
		x = 30
	}

}
