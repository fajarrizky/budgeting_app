package logger

import (
	"encoding/json"
	"log"
)

func Render(v any) string {
	byteStr, err := json.Marshal(v)

	if err != nil {
		log.Printf("failed to log struct: %+v | error: %s", v, err.Error())
		return err.Error()
	}

	return string(byteStr)
}
