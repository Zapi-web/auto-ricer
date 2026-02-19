package executor

import (
	"log"
	"os/exec"
	"time"
)

func UpdateTheme(ch chan string) {
	debDur := 200 * time.Millisecond
	var timer *time.Timer

	for event := range ch {
		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(debDur,  func() {
			cmd := exec.Command("/bin/bash", "./scripts/update_theme.sh", event)
			
			err := cmd.Run()

			if err != nil {
				log.Println(err)
			}
		})

		log.Println("new event:",event)
	}
}
