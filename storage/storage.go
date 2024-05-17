package storage

import (
	"fmt"
	"syscall"

	"github.com/ZainJavedDev/catch-up-check/discord_logger"
)

func CheckStorage() error {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/mnt/volume_sgp1_02", &fs)
	if err != nil {
		fmt.Println(err)
		return err
	}

	available := fs.Bavail * uint64(fs.Bsize)
	availableGB := float64(available) / float64(1<<30)

	storageMessage := fmt.Sprintf("Available storage in attached volume: %.2f GB\n", availableGB)
	fmt.Println(storageMessage)
	discord_logger.SendDiscordMessage(storageMessage)
	return nil
}
