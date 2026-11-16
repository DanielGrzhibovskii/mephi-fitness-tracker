// Package actioninfo processes different activity types through a common interface.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser parses activity data and builds a human-readable report.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info prints valid activity reports and logs errors without stopping the batch.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования отчёта для %q: %v", data, err)
			continue
		}
		fmt.Println(info)
	}
}
