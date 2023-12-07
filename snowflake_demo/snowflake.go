package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {
	//basic
	n, err := snowflake.NewNode(1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id:", id)
		fmt.Println("node:", id.Node(), "step:", id.Step(), "time:", id.Time(), "\n")
	}
	//level
	// t, _ := time.Parse("2006-01-02", "2018-01-01")
	// settings := snowflake.Settings{
	// 	StartTime:      t,
	// 	MachineID:      getMachineID,
	// 	CheckMachineID: checkMachineID,
	// }
	// sf := snowflake.NewSonyflake(settings)
	// id, err := sf.NextID()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(id)

}

// func getMachineID() (uint16, error) {
// 	var machineID uint16
// 	var err error
// 	machineID = readMachineIDFromLocalFile()
// 	if machineID == 0 {
// 		machineID, err := generateMachineID()
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// }
// func checkMachineID(machineID uint16) bool {
// 	saddResult, err := saddMachineIDToRedisSet()
// 	if err != nil || saddResult == 0 {
// 		return true
// 	}

// 	err := saveMachineIDToLocalFile(machineID)
// 	if err != nil {
// 		return true
// 	}

// 	return false
// }
