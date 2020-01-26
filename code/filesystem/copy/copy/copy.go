package copy

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

func CopyFile(src string, dest string) error {

	srcStat, _ := os.Stat(src)
	if srcStat == nil {
		panic("src not found!")
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
			// エラー
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
			newPathName := path.Join(dest, destStat.Name())
			os.Mkdir(newPathName, 0777)
		} else {
			// エラー
		}
	}
	// 元
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// 先
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func getTimestamp() string {

	now := time.Now()
	return now.String()[0:23]
}

func initializeTempDir() string {

	dirpath := os.TempDir()
	timestamp := getTimestamp()
	dir := path.Join(dirpath, "tmp-"+timestamp)
	fmt.Printf("[TARCE] 作業ディレクトリ ... [%v]\n", dir)
	return dir
}
