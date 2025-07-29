package utils

import (
	"database/sql"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

// Closer is a universal Close() func
//
// It closes the given resource if it is a *os.File or *sql.Conn.
func Closer(v interface{}) {
	switch t := v.(type) {
	case *os.File:
		if err := t.Close(); err != nil {
			log.Printf("warning: failed to close file: %v", err)
		}
	case *sql.Conn:
		if err := t.Close(); err != nil {
			log.Printf("warning: failed to close database connection: %v", err)
		}
	case io.ReadCloser:
		if err := t.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}
}

// GetFunctionName returns the name of the function at the given stack frame depth
func GetFunctionName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		fullName := runtime.FuncForPC(pc).Name()
		funcName := path.Ext(fullName)[1:]
		return funcName
	} else {
		return "unknown"
	}
}
