package kubeflake

import (
	"crypto/md5"
	"encoding/binary"
	"os"

	"github.com/sony/sonyflake"
)

var sf = initSonyflake()

func initSonyflake() *sonyflake.Sonyflake {
	st := sonyflake.Settings{
		MachineID: machineID,
	}
	return sonyflake.NewSonyflake(st)
}

func machineID() (uint16, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return 0, err
	}
	sum := md5.Sum([]byte(hostname))
	machineID := binary.BigEndian.Uint16(sum[len(sum)-2:]) // extract last 2 bytes
	return uint16(machineID), nil
}

func New() (uint64, error) {
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Must(id uint64, _ error) uint64 {
	return id
}
