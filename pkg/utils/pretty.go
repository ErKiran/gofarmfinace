package utils

import (
	"encoding/json"
	"fmt"
)

func Pretty(title string, value interface{}) {
	js, _ := json.MarshalIndent(value, "", " ")
	fmt.Println(title, string(js))
}
