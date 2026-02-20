package executor

import (
	"os/exec"
	"time"

	"github.com/Zapi-web/auto-ricer/internal/logger"
)

func UpdateTheme(ch chan string, path string) {
	debDur := 200 * time.Millisecond
	var timer *time.Timer

	for event := range ch {
		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(debDur,  func() {
			cmd := exec.Command("/bin/bash", path, event)

			err := cmd.Run()

			if err != nil {
				logger.Log.Error("Script failed", "event", event, "error", err)
				return
			}
			logger.Log.Info("Script successfully executed")
		})

		logger.Log.Debug("new event from channel", "event", event)
	}
}
