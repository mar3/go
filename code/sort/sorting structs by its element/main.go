package main

import "sort"
import "fmt"

type Idol struct {
	Name string
}

type Idols []Idol

func (this Idol) String() string {
	return fmt.Sprintf("%s", this.Name)
}

func (this Idols) Len() int {
	return len(this)
}

func (this Idols) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this Idols) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func enum_idols() ([]Idol) {
	result := []Idol{}
	result = append(result, Idol{Name: "PASSPO☆"})
	result = append(result, Idol{Name: "じゅじゅ"})
	result = append(result, Idol{Name: "tipToe."})
	result = append(result, Idol{Name: "There There Theres"})
	result = append(result, Idol{Name: "MIGMA SHELTER"})
	result = append(result, Idol{Name: "RHYMEBERRY"})
	result = append(result, Idol{Name: "sora tob sakana"})
	result = append(result, Idol{Name: "・・・・・・・・・"})
	result = append(result, Idol{Name: "トキヲイキル"})
	result = append(result, Idol{Name: "劇場版ゴキゲン帝国"})
	result = append(result, Idol{Name: "ヤなことそっとミュート"})
	return result
}

func main() {
	idols := enum_idols()
	sort.Sort(Idols(idols))
	for _, e := range idols {
		fmt.Println(e)
	}
}
