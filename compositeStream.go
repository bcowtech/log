package log

import "io"

type compositeStream struct {
	writers []io.Writer
}

func (w *compositeStream) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		_, err = writer.Write(p)
		if err != nil {
			return -1, err
		}
	}
	return -1, err
}
