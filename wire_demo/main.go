package main

// func main() {
// 	mess := Initialize()
// 	me := mess.GetMessage()
// 	fmt.Println(me)
// }

func main() {
	// monster := NewMonster()
	// player := NewPlayer("dj")
	// mission := NewMission(player, monster)
	mission := InitMission("dj")
	mission.Start()
}
