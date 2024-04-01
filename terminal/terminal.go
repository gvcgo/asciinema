package terminal

import "io"

type Terminal interface {
	Size() (int, int, error)
	Record(command string, writer io.Writer, envs ...string) error
	Write([]byte) error
}
