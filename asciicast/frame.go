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
	xx := x.([]interface{})
	if len(xx) == 3 {
		f.Time = xx[0].(float64)
		f.EventType = xx[1].(string)
		s := []byte(xx[2].(string))
		b := make([]byte, len(s))
		copy(b, s)
		f.EventData = b
	} else if len(xx) == 2 {
		f.Time = xx[0].(float64)
		s := []byte(xx[1].(string))
		b := make([]byte, len(s))
		copy(b, s)
		f.EventData = b
	}
	return nil
}
