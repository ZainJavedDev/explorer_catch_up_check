package database

import (
	"fmt"
	"time"

	"github.com/ZainJavedDev/catch-up-check/discord_logger"
)

func Check() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sql := "SELECT MAX(start_time) FROM ranked_matches"
	row := db.QueryRow(sql)

	var maxStartTime int64
	err = row.Scan(&maxStartTime)
	if err != nil {
		return err
	}

	maxStart := time.Unix(maxStartTime, 0)
	hoursAgo := time.Since(maxStart).Hours()

	message := fmt.Sprintf("The explorer is %.2f hours behind.\n", hoursAgo)

	fmt.Print(message)
	discord_logger.SendDiscordMessage(message)
	return nil
}
