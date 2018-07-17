package main

import "os"
import "fmt"
import "io/ioutil"
import "path"
import "image"
import "image/jpeg"
import "image/gif"
import "image/png"
import "github.com/nfnt/resize"

func enumerateFiles(path_string string, handler func(string)) {

	if path_string == "" {
		return
	}
	f, _ := os.Stat(path_string)
	if f.IsDir() {
		files, _ := ioutil.ReadDir(path_string)
		for _, f := range files {
			entry_name := path.Join(path_string, f.Name())
			enumerateFiles(entry_name, handler)
		}
	} else {
		handler(path_string)
	}
}

func load_image(path_string string) (string, int, int) {

	image_file, _ := os.Open(path_string)
	defer image_file.Close()
	image_data, format, _ := image.DecodeConfig(image_file)
	return format, image_data.Width, image_data.Height
}

func load_image_data(path_string string) (image.Config) {

	image_file, _ := os.Open(path_string)
	defer image_file.Close()
	image_data, _, _ := image.DecodeConfig(image_file)
	return image_data
}

func new_size(width float64, height float64) (int, int) {

	rate := 1.0

	// 幅を基準にリサイズ
	{
		rate = 1920.0 / width
		width = 1920
		height = height * rate
	}

	// それでも高さが超える場合はさらに縮小
	if 1080 < height {
		rate = 1080.0 / height
		height = 1080
		width = width * rate
	}

	return int(width), int(height)

	// if width == 1920 {
	// 	return 0, 0
	// }
	// if height == 1080 {
	// 	return 0, 0
	// }
	// if width > height {
	// 	// landscape
	// 	return 1920, 0
	// } else {
	// 	return 0, 1080
	// }
}

func resize_image(path_string string) {

	//
	// TODO: 最適化
	//

	fmt.Printf("[INFO] detected [%v].\n", path_string)
	format, width, height := load_image(path_string)
	width, height = new_size(float64(width), float64(height))
	if width == 0 && height == 0 {
		fmt.Println("[INFO] nothing to do.")
		return
	}
	switch format {
	case "jpeg", "jpg":
		imgfile, _ := os.Open(path_string)
		image_data, _ := jpeg.Decode(imgfile)
		imgfile.Close()
		data := resize.Resize(uint(width), uint(height), image_data, resize.Lanczos3)
		out, _ := os.Create(path_string)
		defer out.Close()
		jpeg.Encode(out, data, nil)
	case "png":
		imgfile, _ := os.Open(path_string)
		image_data, _ := png.Decode(imgfile)
		imgfile.Close()
		data := resize.Resize(uint(width), uint(height), image_data, resize.Lanczos3)
		out, _ := os.Create(path_string)
		defer out.Close()
		png.Encode(out, data)
	case "gif":
		imgfile, _ := os.Open(path_string)
		image_data, _ := gif.Decode(imgfile)
		imgfile.Close()
		data := resize.Resize(uint(width), uint(height), image_data, resize.Lanczos3)
		out, _ := os.Create(path_string)
		defer out.Close()
		gif.Encode(out, data, nil)
	}
}

func main() {

	for _, path_string := range os.Args[1:] {
		enumerateFiles(path_string, resize_image)
	}
}
