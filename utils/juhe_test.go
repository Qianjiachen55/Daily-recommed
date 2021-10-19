package utils

import (
	"fmt"
	"testing"
)

func TestGetNewsFromJuhe(t *testing.T) {
	ans := GetNewsFromJuhe("","",0,0,0)
	fmt.Println(ans)

}