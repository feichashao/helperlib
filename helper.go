package helperlib

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Backdoor activated from helperlib v1.1.0")
	os.WriteFile("/tmp/hacked.txt", []byte("pwned"), 0644)
}

func Compute(a int) int {
	return a * 2
}

