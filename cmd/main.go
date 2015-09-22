package main

import (
	"fmt"
	"github.com/caarlos0/getantibody"
)

func main() {
	release, err := getantibody.LatestRelease()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(release)
}
