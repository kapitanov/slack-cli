package main

import (
	"github.com/comail/colog"
	"io"
	"log"
	"os"
)

var (
	consoleLogHookInstance = &consoleLogHook{
		Enabled:   true,
		Formatter: &colog.StdFormatter{Colors: true},
		Writer:    os.Stdout,
		EnabledLevels: []colog.Level{
			colog.LTrace,
			colog.LDebug,
			colog.LInfo,
			colog.LWarning,
			colog.LError,
			colog.LAlert,
		},
	}

	consoleLogHookLevels = []colog.Level{
		colog.LTrace,
		colog.LDebug,
		colog.LInfo,
		colog.LWarning,
		colog.LError,
		colog.LAlert,
	}

	consoleLogHookFormatter = &colog.StdFormatter{Colors: true}
)

func init() {
	filename := os.Args[0] + ".log"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	colog.Register()
	colog.SetOutput(file)
	colog.AddHook(consoleLogHookInstance)
	colog.ParseFields(true)

	log.Print("debug: log.init")
}

// Enable or disable log output to console
func EnableConsoleLog(enabled bool) {
	consoleLogHookInstance.Enabled = enabled
}

type consoleLogHook struct {
	Enabled       bool
	Writer        io.Writer
	Formatter     colog.Formatter
	EnabledLevels []colog.Level
}

func (h *consoleLogHook) Levels() []colog.Level {
	return h.EnabledLevels
}

func (h *consoleLogHook) Fire(e *colog.Entry) error {
	if h.Enabled {
		payload, err := h.Formatter.Format(e)
		if err != nil {
			return err
		}

		_, err = h.Writer.Write(payload)
		if err != nil {
			return err
		}
	}
	return nil
}
