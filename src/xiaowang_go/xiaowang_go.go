package main

import (
	"flag"
	"fmt"
	"os"
	"xiaowang_go/runner"
)

func main() {
	srcfileP := flag.String("s", "", "src file")
	destfileP := flag.String("d", "", "dest file")
	conffileP := flag.String("c", "", "config file")
	//seedP := flag.Int("s", 1, "override seed")
	flag.Parse()

	fmt.Printf("xiaowang gogo! \nconf:%s\nsrc:%s\ndest:%s\n", *conffileP, *srcfileP, *destfileP)

	for _, f := range []string{*conffileP, *srcfileP, *destfileP} {
		if f == "" {
			panic("not all args are provided, check usage with ./xiaowang_go -h")
		}
	}

	for _, f := range []string{*conffileP, *srcfileP} {
		panic_if_no_exists(f)
	}

	runner.Run(*conffileP, *srcfileP, *destfileP)
}

func panic_if_no_exists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		fmt.Printf("file does not exist", f)
		panic("file")
	}
	return true
}
