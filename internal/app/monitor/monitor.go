package monitor

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Monitor struct {
	websiteList []string
}

func NewMonitor() *Monitor {
	return &Monitor{
		websiteList: make([]string, 0),
	}
}

// TODO: в последствии можно сделать выгрузку списка из webiste.Websites
// а способ загрузки списка указывать при конфигурации
func (mntr *Monitor) LoadListFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		website := scanner.Text()
		mntr.websiteList = append(mntr.websiteList, website)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (mntr *Monitor) StartMonitoring() {
	go func() {
		for _, wsite := range mntr.websiteList {
			_, err := AccessTime(wsite)
			if err != nil {
				fmt.Errorf("Error accessing site %s: %w\n", wsite, err)
			}
		}
		time.Sleep(1 * time.Minute)
	}()
}

func AccessTime(url string) (time.Duration, error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	return elapsed, nil
}
