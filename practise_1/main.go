package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if ok := scanner.Scan(); !ok {
			fmt.Println("beda")
			return
		}

		fields := strings.Fields(scanner.Text())

		for i := range fields {
			fmt.Println(fields[i])
		}

	}

}
