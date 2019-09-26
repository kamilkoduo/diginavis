package playground

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

func CalculateEntropy(data string) (entropy float64, err error) {
	entropy = 0
	l := len(data)
	if l == 0 {
		err = errors.New("Expected non-empty string")
		return entropy, err
	}

	for i := 0; i < 256; i++ {
		b := string(byte(i))
		fi := strings.Count(data, b)
		pi := float64(fi) / float64(l)
		if pi > 0 {
			entropy += -pi * math.Log2(pi)
		}
	}
	return entropy, err
}
func PrintNonEmptyStr(data string) (err error) {
	if len(data) == 0 {
		err = errors.New("Cannot print empty str")
	} else{
		fmt.Print(data)
	}
	return err
}

func ReturnStringAfter2s(data string) (string, error)  {
	time.Sleep(time.Duration(2*1e9))
	return data, nil
}


func main() {
	fmt.Println("Go testing. Entropy")
	ent, _ := CalculateEntropy("12345")
	fmt.Printf("%.3f", ent)
}
