package helpers

import (
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/wikipedia/wiki_web_crawler/constants"
	"bufio"
	"log"
	"os"
)

func LoadVisitedFromLog(visited map[string]bool) (map[string]bool, error) {
	files, err := global_helpers.GetAllFiles(constants.VisitedLogDir)
	//if no log files exist
	if err != nil {
		return visited, err
	}
	lastLog, err := getLastFile(files)
	if err != nil {
		return visited, err
	}

	//open log file
	lastLogFile, err := os.Open(lastLog)

	if err != nil {
		return visited, err
	}
	logScanner := bufio.NewScanner(lastLogFile)
	logScanner.Split(bufio.ScanLines)

	for logScanner.Scan() {
		visited[logScanner.Text()] = true
	}
	err = lastLogFile.Close()
	if err != nil {
		return visited, err
	}
	log.Println("visited initialized from log")
	return visited, nil
}
