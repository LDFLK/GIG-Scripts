package helpers

import (
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/wikipedia/wiki_web_crawler/constants"
	"bufio"
	"log"
	"os"
)

func LoadQueueFromLog(queue chan string) error {
	files, err := global_helpers.getAllLogs(constants.QueueLogDir)
	//if no log files exist
	if err != nil {
		return err
	}
	lastLog, err := getLastFile(files)
	if err != nil {
		return err
	}

	//open log file
	lastLogFile, err := os.Open(lastLog)

	if err != nil {
		return err
	}
	logScanner := bufio.NewScanner(lastLogFile)
	logScanner.Split(bufio.ScanLines)

	for logScanner.Scan() {
		go func(url string) { queue <- url }(logScanner.Text())
	}
	err = lastLogFile.Close()
	if err != nil {
		return err
	}
	log.Println("queue initialized from log")
	return nil
}
