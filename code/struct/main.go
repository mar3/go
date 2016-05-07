package main

import "fmt"

type Player struct {
	name string
}

//
// 構造体を参照で受け取る関数
//
func test1(u *Player) {

	// WARNING: ここでの変更は呼び出し側に影響を及ぼす
	//                 C++ や Java の参照と同じ
	u.name = "Albert Collins"
}

//
// 構造体をコピーで受け取る関数
//
func test2(u Player) {

	// SAFE: ここでの変更は呼び出し側に影響を及ぼさない
	//            C++ のコピーと同じ
	u.name = "Yngwie Malmsteen"
}

func main() {

	user := Player{
		name: "Wynton Kelly",
	}

	test1(&user)

	test2(user)

	fmt.Println(user)
}
