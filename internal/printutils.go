package internal

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type M map[string]interface{}

func PrettyPrintJSON(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "    ")
	cobra.CheckErr(err)

	fmt.Println(string(bytes))
}
