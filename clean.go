package baja

import (
	"fmt"
	"os"
)

func Clean() {
	cleans := []string{"public"}
	for _, d := range cleans {
		fmt.Println("Clean", d)
		os.RemoveAll(fmt.Sprintf("./%s", d))
	}
}