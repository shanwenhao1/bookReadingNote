package distributedId

import (
	"fmt"
	"github.com/sony/sonyflake"
	"os"
	"time"
)

/*
	自定义函数, 获取MachineID， 如不设置, 默认本机IP的低16位作为machine Id
*/
func getMachineID() (uint16, error) {
	/*
		var machineID uint16
		var err error
		machineID = readMachineIDFromLocalFile()
		if machineID == 0 {
			machineID, err = generateMachineID()
			if err != nil {
				return 0, err
			}
		}

		return machineID, nil
	*/
	return 1, nil
}

/*
	自定义函数, 用于检查MachineID是否冲突
*/
func checkMachineID(machineID uint16) bool {
	/*
		sAddResult, err := saddMachineIDToRedisSet()
		if err != nil || sAddResult == 0 {
			return true
		}

		err = saveMachineIDToLocalFile(machineID)
		if err != nil {
			return true
		}

		return false
	*/
	return true
}

func UseSonyFlake() {
	t, _ := time.Parse("2006-01-02", "2020-04-11")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	var sf *sonyflake.Sonyflake
	sf = sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)
}
