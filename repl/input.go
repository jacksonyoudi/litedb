package repl

import "bytes"

type InputBuffer struct {
	buffer bytes.Buffer
}

func newInputBuffer() *InputBuffer {
	return &InputBuffer{
		buffer: bytes.Buffer{},
	}
}

func (in *InputBuffer) ReadString() string {
	return in.buffer.String()
}

func (in *InputBuffer) Reset() {
	in.buffer.Reset()
}

func (in *InputBuffer) Len() int {
	return in.buffer.Len()
}

func (in *InputBuffer) Write(s string) (int, error) {
	return in.buffer.WriteString(s)
}
