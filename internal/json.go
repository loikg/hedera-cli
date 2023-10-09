package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

// M is shothand alias for map[string]interface{}
type M map[string]interface{}

func (m M) String() string {
	bytes, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal internal.M: %v", err))
	}
	return string(bytes)
}

// ConsolePrint pretty print the given data as json to w io.Writer
func ConsolePrint(w io.Writer, data any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
