package application

import "fmt"
import "copy/copy"

type Application struct {
}

func (this *Application) Copy(left string, right string) {

	fmt.Printf("[%s] >> [%s]\n", left, right)

	copy.CopyFile(left, right)
}
