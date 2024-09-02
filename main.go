package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("数当てゲームへようこそ！1から100までの数を当ててください。")

	for {
		fmt.Print("あなたの予想: ")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("無効な入力です。整数値を入力してください。")
			// 標準入力のバッファをクリア
			fmt.Scanf("%*s")
			continue
		}

		if guess < target {
			fmt.Println("もっと大きい数です。")
		} else if guess > target {
			fmt.Println("もっと小さい数です。")
		} else {
			fmt.Println("おめでとうございます！正解です！")
			break
		}
	}
}
