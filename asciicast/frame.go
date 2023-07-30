package asciicast

import (
	"encoding/json"
	"fmt"
)

type Frame struct {
	Time      float64 // Delay
	EventType string
	EventData []byte //Data
}

func (f *Frame) MarshalJSON() ([]byte, error) {
	s, _ := json.Marshal(string(f.EventData))
	json := fmt.Sprintf(`[%.6f, "o", %s]`, f.Time, s)
	return []byte(json), nil
}

func (f *Frame) UnmarshalJSON(data []byte) error {
	var x interface{}

	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}

	f.Time = x.([]interface{})[0].(float64)

	s := []byte(x.([]interface{})[1].(string))
	b := make([]byte, len(s))
	copy(b, s)
	f.EventData = b

	return nil
}
