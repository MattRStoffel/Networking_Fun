package internal

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func ReverseWords(buf []byte) []byte {
	tmp := strings.Split(string(buf), " ")
	slices.Reverse(tmp)
	return []byte(strings.Join(tmp, " "))
}

func GetInput() string {
	fmt.Println("input a lowwercase sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(input, "\n")[0]
}
