package main

import (
        "fmt"
        "time"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculateMeetingDuration() returns the number of minutes that has passed or an error if there is one
func calculateMeetingDuration(start, end string) (int, error) {
        // code goes here
        layout := "15:04" // The correct layout for "HH:MM" (24-hour format)

        startTime, err := time.Parse(layout, start)
        if err != nil {
                return 0, fmt.Errorf("error parsing start time: %w", err)
        }

        endTime, err := time.Parse(layout, end)
        if err != nil {
                return 0, fmt.Errorf("error parsing end time: %w", err)
        }

        duration := endTime.Sub(startTime)

        return int(duration.Minutes()), nil
}

func main() {
        var result string
        meetingTimes := []struct {
                start, end string
        }{
                {"09:00", "10:30"},
                {"14:15", "16:00"},
                {"23:00", "23:30"},
        }
        for _, tc := range meetingTimes {
                duration, err := calculateMeetingDuration(tc.start, tc.end)
                if err != nil {
                        result += fmt.Sprintf("Error: %v\n", err)
                } else {
                        result += fmt.Sprintf("Meeting lasted %d minutes\n", duration)
                }
        }
        fmt.Print(result)
}