package time_log

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type timeLog struct{}

type TimeLogData struct {
	Duration    string
	StartTime   string
	EndTime     string
	Mode        string
	Service     string
	Description string
}

func NewTimeLog() *timeLog {
	return &timeLog{}
}

func (tl *timeLog) Log(timeLogData TimeLogData) error {
	duration, err := convertToSeconds(timeLogData.Duration)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse("15:04", timeLogData.StartTime)
	if err != nil {
		return err
	}

	today := time.Now()
	dateTime := time.Date(today.Year(), today.Month(), today.Day(), timeParsed.Hour(), timeParsed.Minute(), 0, 0, today.Location())
	log.Printf("duration %d", duration)
	log.Println("Parsed time:", dateTime)
	return nil
}

func convertToSeconds(value string) (int, error) {
	// Convert the given duration to seconds
	splitTime := strings.Split(value, ":")
	if len(splitTime) < 1 {
		return 0, fmt.Errorf("invalid time passed")
	}

	multiplierIndex := 0
	multipliers := []int{1, 60, 60 * 60, 60 * 60 * 24}

	timeInSeconds := 0
	for i := len(splitTime) - 1; i >= 0; i-- {
		c, err := strconv.Atoi(splitTime[i])
		if err != nil {
			return 0, err
		}

		timeInSeconds += c * multipliers[multiplierIndex]

		// Move to next index
		multiplierIndex += 1
	}

	return timeInSeconds, nil
}
