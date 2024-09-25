package create

import (
	"fmt"
	"log"
	"os"
	process_monitor "saruman/cmd/system"
	"strings"
	"time"
)

func Create(processName string, message string) {
	stats, _ := process_monitor.CollectSystemStats(processName)

	filename := processName + ".json"

	currentTime := time.Now()

	line := fmt.Sprintf(
		"{'cpu_usage': '%.2f%%', 'memory_usage': '%.2f GiB', 'total_system_memory': '%.2f GiB',"+
			" 'available_system_memory': '%.2f GiB', 'number_of_CPUs': %d, 'message': '%s', 'datetime': '%s'}\n",
		stats.ProcessCpu,
		stats.ProcessMemory,
		stats.TotalSystemMemory,
		stats.AvailableSystemMemory,
		stats.NumCpu,
		message,
		currentTime,
	)

	log.Printf(line)
	line = strings.ReplaceAll(line, "'", "\"")

	createFile(line, filename)
}

func createFile(content string, filename string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		log.Printf("Erro ao escrever no arquivo: %v\n", err)
		return false, err
	}

	return true, nil
}
