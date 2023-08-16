package writer

import "io"

func Write(w io.Writer, strings []interface{}) {
	for _, s := range strings {
		if str, ok := s.(string); ok {
			w.Write([]byte(str))
		}
	}
}
