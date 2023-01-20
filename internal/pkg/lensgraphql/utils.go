package lensgraphql

import (
	"encoding/json"
	"fmt"
	"strings"
)

func PrintJson(obj interface{}) {
	jsonVal, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonVal))
}

func isValidHex(s string) bool {
	s = strings.TrimPrefix(s, "0x")
	if len(s)%2 != 0 {
		fmt.Println("Provided Hex value is odd-length")
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func isValidProfileId(s string) bool {
	if strings.HasPrefix(s, "0x") && isValidHex(s) {
		return true
	}
	fmt.Println("Please provide a valid Lens profileId (e.g. 0x05)")
	return false
}

func getFirstArg(args []string) string {
	if len(args) < 1 {
		fmt.Println("Please provide an argument")
		return ""
	}
	return args[0]
}
