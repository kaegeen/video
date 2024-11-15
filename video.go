package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Function to log the video you watched
func logVideo(videoTitle string) error {
	// Open the log file (create if it doesn't exist)
	file, err := os.OpenFile("video_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the video title and current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err = fmt.Fprintf(file, "[%s] %s\n", timestamp, videoTitle)
	if err != nil {
		return err
	}

	return nil
}

// Function to count how many videos were watched today
func countVideosToday() (int, error) {
	file, err := os.Open("video_log.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)

	// Get today's date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")

	// Loop through each line and count videos watched today
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 10 && line[:10] == today { // Check if the line's timestamp matches today's date
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func main() {
	// Menu for user input
	fmt.Println("YouTube Video Tracker")
	fmt.Println("----------------------")
	fmt.Println("1. Log a new video")
	fmt.Println("2. View videos watched today")
	fmt.Println("3. Exit")
	fmt.Print("Please choose an option: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Prompt user for video title and log it
		fmt.Print("Enter the video title or URL you watched: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		videoTitle := scanner.Text()

		// Log the video watched
		err := logVideo(videoTitle)
		if err != nil {
			fmt.Println("Error logging video:", err)
		} else {
			fmt.Println("Video logged successfully!")
		}

	case 2:
		// Count and display videos watched today
		count, err := countVideosToday()
		if err != nil {
			fmt.Println("Error counting videos:", err)
		} else {
			fmt.Printf("You watched %d videos today.\n", count)
		}

	case 3:
		// Exit the program
		fmt.Println("Exiting program.")
		return

	default:
		fmt.Println("Invalid option. Please try again.")
	}
}
