package utils

import (
	"fmt"
	"time"
)

func GetFileNameAndDirectory(slug string, postedAt time.Time) (string, string, string) {
	// Extract components
	year := postedAt.Year()
	month := postedAt.Month().String()
	day := postedAt.Day()
	hour := postedAt.Hour()
	minute := postedAt.Minute()

	// Format the components as per your requirement
	monthFormatted := fmt.Sprintf("%02d", postedAt.Month()) // E.g., 04 for April
	dayFormatted := fmt.Sprintf("%02d", day)                // E.g., 23 for 23rd
	hourFormatted := fmt.Sprintf("%02d", hour)              // E.g., 11 for 11 AM
	minuteFormatted := fmt.Sprintf("%02d", minute)          // E.g., 39 for 39 minutes
	yearStr := fmt.Sprintf("%d", year)

	// Generate the directory path and file name
	directoryPath := fmt.Sprintf("%d/%s", year, month)
	fileName := fmt.Sprintf("%s%s%d@%s_%s_%s_%sh_%s.md", yearStr, month, day, yearStr, monthFormatted, dayFormatted, hourFormatted+minuteFormatted, slug)
	fullPath := fmt.Sprintf("%s/%s", directoryPath, fileName)

	return fileName, directoryPath, fullPath
}
