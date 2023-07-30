package asciicast

import "time"

type Stream struct {
	Frames        []Frame
	elapsedTime   time.Duration
	lastWriteTime time.Time
	maxWait       time.Duration
}

func NewStream(maxWait float64) *Stream {
	return &Stream{
		lastWriteTime: time.Now(),
		maxWait:       time.Duration(maxWait*1000000) * time.Microsecond,
	}
}

func (s *Stream) Write(p []byte) (int, error) {
	frame := Frame{}
	frame.EventType = "o"
	frame.Time = s.incrementElapsedTime().Seconds()
	frame.EventData = make([]byte, len(p))
	copy(frame.EventData, p)
	s.Frames = append(s.Frames, frame)

	return len(p), nil
}

func (s *Stream) Close() {
	s.incrementElapsedTime()

	if string(s.Frames[len(s.Frames)-1].EventData) == "exit\r\n" {
		s.Frames = s.Frames[:len(s.Frames)-1]
	}
}

func (s *Stream) Duration() time.Duration {
	return s.elapsedTime
}

func (s *Stream) incrementElapsedTime() time.Duration {
	now := time.Now()
	d := now.Sub(s.lastWriteTime)

	if s.maxWait > 0 && d > s.maxWait {
		d = s.maxWait
	}

	s.elapsedTime += d
	s.lastWriteTime = now

	return s.elapsedTime
}
