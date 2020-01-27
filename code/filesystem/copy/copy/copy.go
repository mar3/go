package copy

import (
	"copy/util"
	"fmt"
	"io"
	"os"
	"path"
)

// ファイルまたはディレクトリを複製します。
func CopyFile(src string, dest string) error {

	srcStat, _ := os.Stat(src)
	if srcStat == nil {
		fmt.Print("[ERROR] src not found!")
		return nil
	}

	destStat, err := os.Stat(dest)
	if err != nil {
		fmt.Print("[ERROR] Unknown device. reason: ")
		fmt.Print(err)
		fmt.Println()
	} else if destStat == nil {
		fmt.Print("[ERROR] Unknown device. (nil)")
		fmt.Println()
	} else if destStat.IsDir() {
		// ディレクトリが存在する場合
		if srcStat.IsDir() {
			// ディレクトリを作成します。
			newPathName := path.Join(dest, srcStat.Name())
			fmt.Printf("[TRACE] Creating a new directory ... [%s]\n", newPathName)
			os.Mkdir(newPathName, 0777)
		} else {
			// ファイルを作成します。
			newPathName := path.Join(dest, srcStat.Name())
			fmt.Printf("[TRACE] Creating a new file ... [%s]\n", newPathName)
			file, err := os.Create(newPathName)
			if err != nil {
				fmt.Print("[ERROR] Cannot create file. reason: ")
				fmt.Print(err)
				fmt.Println()
				return nil
			}
			rfile, err := os.Open(src)
			if err != nil {
				fmt.Print("[ERROR] Cannot read file. reason: ")
				fmt.Print(err)
				fmt.Println()
				return nil
			}
			defer rfile.Close()
			defer file.Close()
			io.Copy(file, rfile)
		}
	} else {
		// ディレクトリが存在する場合
		if destStat.IsDir() {
			// ディレクトリを作成します。
			newPathName := path.Join(dest, srcStat.Name())
			fmt.Printf("[TRACE] Creating a new directory ... [%s]\n", newPathName)
			os.Mkdir(newPathName, 0777)
		} else {
			// エラー
		}
	}
	return nil
}

// 一時ディレクトリを初期化します。
func initializeTempDir() string {

	dirpath := os.TempDir()
	timestamp := util.GetTimestamp()
	dir := path.Join(dirpath, "tmp-"+timestamp)
	fmt.Printf("[TARCE] 作業ディレクトリ ... [%v]\n", dir)
	return dir
}
