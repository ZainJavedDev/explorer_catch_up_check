package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/ZainJavedDev/catch-up-check/database"
	"github.com/ZainJavedDev/catch-up-check/discord_logger"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "SELECT MAX(start_time) FROM ranked_matches"
	row := db.QueryRow(sql)

	var maxStartTime int64
	err = row.Scan(&maxStartTime)
	if err != nil {
		log.Fatal(err)
	}

	maxStart := time.Unix(maxStartTime, 0)
	hoursAgo := time.Since(maxStart).Hours()

	message := fmt.Sprintf("The explorer is %.2f hours behind.\n", hoursAgo)

	fmt.Print(message)
	discord_logger.SendDiscordMessage(message)

	fs := syscall.Statfs_t{}
	err = syscall.Statfs("/mnt/volume_sgp1_01", &fs)
	if err != nil {
		fmt.Println(err)
		return
	}

	available := fs.Bavail * uint64(fs.Bsize)
	availableGB := float64(available) / float64(1<<30)

	storageMessage := fmt.Sprintf("Available storage in attached volume: %.2f GB\n", availableGB)
	fmt.Println(storageMessage)
	discord_logger.SendDiscordMessage(storageMessage)
}
