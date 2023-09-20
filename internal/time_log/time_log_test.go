package time_log_test

import (
	"testing"

	"github.com/ricomonster/revelations/internal/time_log"
)

func TestLog(t *testing.T) {
	svc := time_log.NewTimeLog()
	err := svc.Log(time_log.TimeLogData{
		Duration:  "02:43",
		StartTime: "09:30",
	})
	if err != nil {
		t.Errorf("TestLog: Error %v", err)
	}
}
