package main

import "fmt"

// バッチ処理基底クラス
type IBatch interface {
	Invoke()
}

// Batch1: XXX CSV 登録処理
type Batch1 struct {
}

func (this *Batch1) Invoke() {
	fmt.Println("[TARCE] <Batch1.Invoke()> !")
}

// コンフィギュレーションクラス
type Configuration struct {
}

func (this *Configuration) LoadBatchInstance(id string) IBatch {
	switch id {
	case "Batch1":
		batch := Batch1{}
		return &batch
	}
	return nil
}

// タスク管理
type TaskManager struct {
}

func (this *TaskManager) Execute(batch IBatch) {

	// バッチ処理オブジェクトを初期化します。
	if batch == nil {
		fmt.Println("[WARN] batch is nil.")
		return
	}
	// バッチ処理を呼び出します。
	batch.Invoke()
}

// アプリケーションクラス
type Application struct {
}

func (this *Application) Run() {

	fmt.Println("### START ###")
	conf := Configuration{}
	batch := conf.LoadBatchInstance("Batch1")
	runner := TaskManager{}
	runner.Execute(batch)
	fmt.Println("Ok.")
	fmt.Println("--- END ---")
}

// エントリーポイント
func main() {

	app := Application{}
	app.Run()
}
