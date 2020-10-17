package exec

import (
	"bytes"

	"github.com/rs/zerolog/log"
)

type writer struct {
	prefix string
	remain []byte
}

func (w *writer) Write(p []byte) (n int, err error) {
	idx := bytes.IndexByte(p, '\n')
	if idx < 0 {
		w.remain = append(w.remain, p...)
	} else {
		log.Info().Msgf("[%s] %s%s", w.prefix, w.remain, p[:idx])
		w.remain = append(w.remain[:0], p[idx+1:]...)
	}
	return len(p), nil
}

func (w *writer) flush() {
	if len(w.remain) > 0 {
		log.Info().Msgf("[%s] %s", w.prefix, w.remain)
	}
}
