package lensgraphql

import (
	"encoding/json"
	"fmt"
)

func printJson(obj interface{}) {
	jsonVal, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonVal))
}
