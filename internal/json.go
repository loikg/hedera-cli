package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

type M map[string]interface{}

func (m M) String() string {
	bytes, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal internal.M: %v", err))
	}
	return string(bytes)
}

func ConsolePrint(w io.Writer, data any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
