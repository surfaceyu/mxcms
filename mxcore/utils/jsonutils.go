package utils

import (
	"encoding/json"
)

func Json2StringSlice(str string) []string{
	var res []string
	json.Unmarshal([]byte(str), &res)
	return res
}
