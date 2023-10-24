package common

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(label string, data interface{}) {
	marhsalledData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s %s\n", label, string(marhsalledData))
}
