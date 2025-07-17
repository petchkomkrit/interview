package main

import (
	"fmt"

	"github.com/go-beginner-training/quiz1"
	"github.com/go-beginner-training/quiz2"
)

func main() {
	fmt.Printf("Quiz 1 : %+v", quiz1.IsValidPassword("P@assw0rd"))
	fmt.Printf("Quiz 1 : %+v", quiz2.FirstNonRepeatingChar("Success"))
}
