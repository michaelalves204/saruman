package find_process

import (
	"errors"

	"github.com/shirou/gopsutil/process"
)

func ByName(procName string) (int32, error) {
	processes, err := process.Processes()

	if err != nil {
		return 0, err
	}

	var pid int32

	for _, proc := range processes {
		name, err := proc.Name()
		if err == nil && name == procName {
			pid = proc.Pid
			return pid, nil
		}
	}

	return -1, errors.New("PID NOT FOUND")
}
