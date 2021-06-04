package main
import (
	"fmt"
	"math"
	"time"
	"sort"
)

func main() {
	//test2()
	//test3()
	fmt.Println("Hello World")
}

//class battery
type Battery struct {
	ID int
	Status string
	ColumnsList []Column
	FloorRequestButtonsList []FloorRequestButton
	//_amountOfColumns int
	//_amountOfFloors int
	//_amountOfBasements int
	//_amountOfElevatorPerColumn int
}

//class column
type Column struct {
	ColID byte
	Status string
	ServedFloors []int
	IsBasement bool
	ElevatorsList []Elevator
	CallButtonsList []CallButton
	_amountOfElevators int
}

//class elevator
type Elevator struct {
	ID int
	Status string
	CurrentFloor int
	Direction string
	Door Door
	FloorRequestsList []int
	CompletedRequestsList []int
}

//class callbutton
type CallButton struct {
	ID int
	Status string
	Floor int
	Direction string
}

//class floorrequestbutton
type FloorRequestButton struct {
	ID int
	Status string
	Floor int
	Direction string
}

//class door
type Door struct {
	ID int
	Status string
}

///////////////////////////////////////////////////BATTERY SECTION///////////////////////////////////////////////////
//battery ctor
func(b *Battery) Init(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) {
      b.ID = _id
      b.Status = "online"
      b.ColumnsList = b.ColumnsList
      b.FloorRequestButtonsList = b.FloorRequestButtonsList
//DOUBLE-CHECK//////////////////////////////
      var colArr = [4]byte{ 'A', 'B', 'C', 'D'}
		for i := 0; i < len(colArr); i++ {
			column := new(Column)
			column.ColID = byte(i)
			b.ColumnsList = append(b.ColumnsList, *column)  ///check
			//ColumnsList.append(column);
		}
	//DOUBLE-CHECK//////////////////////////////
		for i := 0; i <= _amountOfFloors; i++ {
			floor := -6
			floorRequestID := -6
			if (floor < 0) {
				floorb := FloorRequestButton{}
				// floorb.FloorbID = i;
				b.FloorRequestButtonsList = append(b.FloorRequestButtonsList, floorb)
				floor--
				floorRequestID--
			} else if (floor >= 0) {
				floorb := FloorRequestButton{}
				// floorb.FloorbID = i;
				b.FloorRequestButtonsList = append(b.FloorRequestButtonsList, floorb)
				floor++
				floorRequestID++
			}
		}
}

//method to assign elevator
func(b Battery) AssignElevator( _requestedFloor int, _direction string) Elevator {
	bestColumn := b.FindBestColumn(_requestedFloor)
 	chosenElevator := bestColumn.FindBestElevator(_requestedFloor, _direction)
	return chosenElevator
}

//method to find best column
func(b *Battery) FindBestColumn(_requestedFloor int) Column {
	var chosenColumn Column
	for _, currentColumn := range b.ColumnsList {
	if currentColumn.ServedFloors[0] == _requestedFloor {
		chosenColumn = currentColumn
		break
	}
}
	return chosenColumn
}
///////////////////////////////////////////////////COLUMN SECTION///////////////////////////////////////////////////
//column ctor
func(c *Column) Init(_id byte, _amountOfElevators int, _servedFloors int, _isBasement bool) {
	c.ColID = _id
	c.Status = "online"
	c.ServedFloors = c.ServedFloors
	c.IsBasement = false
	c.ElevatorsList = c.ElevatorsList
	c.CallButtonsList = c.CallButtonsList

	//push to elevatorslist
	for i := 1; i < c._amountOfElevators; i++ {
		elevator := Elevator()
		c.ElevatorsList = append(c.ElevatorsList, elevator)
	}

	//for i := 0; i < _amountOfElevators; i++ {
	//	elevator := new(Elevator[i])
	//	c.ElevatorsList = append(c.ElevatorsList, *elevator)
	//}

	//for i := 0; i < _amountOfElevators; i++ {
	//	c.ElevatorsList = append(c.ElevatorsList, Elevator{})
	//}

	////////////////////double-check
	///call-buttons
	for j := 0; j < _servedFloors; j++ {
		if (j != 1) {
			//button := new(CallButton(j, j, "down"))
			button := CallButton{}
			c.CallButtonsList = append(c.CallButtonsList, button)
		} else if (j == 1 || j == -6 || j == -5 || j == -4 || j == -3 || j == -2 || j == -1) {
			//button := new(CallButton(j, j, "up"))
			button := CallButton{}
			c.CallButtonsList = append(c.CallButtonsList, button)
		}
	}
}
 func(c *Column) PushServedFloors(_amountOfColumns int) {
 	//for item := 0; item < _amountOfColumns; item++ {
 		if (c.ColID == 'A') {
 			for i := 0; i > -6; i-- {
 				c.ServedFloors = append(c.ServedFloors, i)
			}
		} else if (c.ColID == 'B') {
			for i := 2; i < 21; i++ {
				c.ServedFloors = append(c.ServedFloors, i)
			}
		} else if (c.ColID == 'C') {
			for i := 21; i < 41; i++ {
				c.ServedFloors = append(c.ServedFloors, i)
			}
		} else if (c.ColID == 'D') {
			for i := 41; i < 61; i++ {
				c.ServedFloors = append(c.ServedFloors, i)
			}
		} else {
			fmt.Println("Fail")
		}
	//}
 }

 //func CreateCallButtons(_amountOfFloors int) {
 //	buttonFloor := 1
 //	callButtonID := 0
 //	for i := 0; i < _amountOfFloors; i++ {
 //		if (buttonFloor < _amountOfFloors) {
 //			callButton := new(CallButton(callButtonID, buttonFloor, "up"))
 //			c.CallButtonsList.append(callButton)
 //			callButtonID++
	//	}
	//	if (buttonFloor > 1) {
	//		callButton := new(CallButton(callButtonID, buttonFloor, "down"))
	//		c.CallButtonsList.append(callButton)
	//		callButtonID++
	//	} else {
	//		buttonFloor++
	//	}
	//}
 //}

 //func(c *Column) CreateElevators(_amountOfFloors, _amountOfElevators int) {
 //	for i := 0; i <= _amountOfElevators; i++ {
 //		elevator := new(Elevator(i, _amountOfFloors))
 //		c.ElevatorsList = append(c.ElevatorsList, elevator)
	//}
 //}

 func (c *Column) RequestElevator(_requestedFloor int, _direction string) Elevator {
 	elevator := c.FindBestElevator(_requestedFloor, _direction)
 	elevator.FloorRequestsList = append(elevator.FloorRequestsList, _requestedFloor)
 	elevator.SortFloorList()
 	e := Elevator{} ///double-check
 	e.Go()
 	fmt.Println("<OPENING DOORS")
 	return elevator
 }

 var bestElevatorInfo = map[string]int {
"rank":           5,
"distance":       1000,
}

var bestElevatorInfo1 = map[string]Elevator {
	"chosenElevator": Elevator{},
}

 func(c *Column) FindBestElevator(_requestedFloor int, _direction string) {
	 for _, elevator := range c.ElevatorsList {
		 if (_requestedFloor == elevator.CurrentFloor && _direction == elevator.Direction) {
			 isBestElevator(elevator, 1)
		 } else if (_requestedFloor < elevator.CurrentFloor && elevator.Direction == "down" && _direction == elevator.Direction) {
			 isBestElevator(elevator, 2)
		 } else if (_requestedFloor > elevator.CurrentFloor && elevator.Direction == "up" && _direction == elevator.Direction) {
			 isBestElevator(elevator, 2)
		 } else if (elevator.Status == "idle") {
			 isBestElevator(elevator, 3)
		 } else {
			 isBestElevator(elevator, 4)
		 }
	 }
 }

func isBestElevator(elevator Elevator, elevatorRank int) Elevator {
	//bestElevatorInfo := make(map[string]int)
	if (elevatorRank < bestElevatorInfo["rank"]) {
		bestElevatorInfo1["chosenElevator"] = elevator
		bestElevatorInfo["rank"] = elevatorRank
		bestElevatorInfo["distance"] = int(math.Abs(elevator.CurrentFloor - _requestedFloor))
	} else if (elevatorRank == bestElevatorInfo["rank"]) {
		var elevatorDistance = int(math.Abs(elevator.CurrentFloor - _requestedFloor))
		if (elevatorDistance < bestElevatorInfo["distance"]) {
			bestElevatorInfo1["chosenElevator"] = elevator
			bestElevatorInfo["distance"] = elevatorDistance
		}
	}
	return bestElevatorInfo1["chosenElevator"]
}



 //	var chosenElevator int
 //	if (_requestedFloor == 1) {
 //		for _, elevator := range c.ElevatorsList {
 //			if (elevator.CurrentFloor == 1 && elevator.Status == "stopped") {
 //				chosenElevator = elevator
	//		}else if (elevator.CurrentFloor ==1 && elevator.Status == "idle") {
	//			chosenElevator = elevator
	//		} else if (_requestedFloor != 1) {
	//			chosenElevator = elevator;
	//		}
	//	}
	//}
	//return chosenElevator
 //}

///////////////////////////////////////////////////ELEVATOR SECTION///////////////////////////////////////////////////
//elevator ctor
func(e *Elevator) Init(_id int) {
	e.ID = _id
	e.Status = "idle"
	e.CurrentFloor = 1
	e.Direction = ""
	e.Door = Door{}
	e.FloorRequestsList = e.FloorRequestsList
	e.CompletedRequestsList = e.CompletedRequestsList
}
////double-check//////////////////////////////
//func(e *Elevator) PushToFloorRequestsList(_amountOfFloors int) {
//	buttonFloor := 1
//	var direction string
//	for i := 0; i < _amountOfFloors; i++ {
//		//var _ = new(FloorRequestButton{ID, buttonFloor, direction})
//		var floorRequestButton = FloorRequestButton{}
//		e.FloorRequestsList = append(e.FloorRequestsList, floorRequestButton)
//		buttonFloor++
//		ID++
//		break
//	}
//}

///double-check/////////
func(e *Elevator) RequestFloor(floor int) {
	e.FloorRequestsList = append(e.FloorRequestsList)
	e.Go()
	var door = new(Door) ///double-check
	door.SwingDoors()
}

func(d *Door) SwingDoors() {
	d.Status = "open"
	fmt.Println("<< Opening Doors >>")
	time.Sleep(1 * time.Second)
	d.Status = "closed"
	fmt.Println(">> Closing Doors <<")
}
//double-check////////////
func(e *Elevator) Go() {
	for len(e.FloorRequestsList) != 0 {
		destination := e.FloorRequestsList[0]
		e.Status = "active"
		if (e.CurrentFloor < destination) {
			e.Direction = "up"
			for (e.CurrentFloor < destination) {
				fmt.Println("Elevator # ", e.ID, " is now at floor ", e.CurrentFloor)
				e.CurrentFloor++
			}
		} else if (e.CurrentFloor > destination) {
			e.Direction = "down"
			for (e.CurrentFloor > destination) {
				fmt.Println("Elevator # ", e.ID, " is now at floor ", e.CurrentFloor)
				e.CurrentFloor--
			}
		}
		e.Status = "stopped"
		//e.FloorRequestsList[len(e.FloorRequestsList)-1]
		a := e.FloorRequestsList
		//i := 0

		//a[i] = a[len(a)-1]
		a[len(a)-1] = ""
		a = a[:len(a)-1]

		//e.FloorRequestsList[:len(e.FloorRequestsList)0]
		fmt.Println("Elevator # ", e.ID, " is now at floor ", e.CurrentFloor)

		var d = new(Door)
		d.SwingDoors() ////double-check
	}
	e.Status = "idle"
}

func(e *Elevator) SortFloorList() {
	if (e.Direction == "up") {
		sort.Ints(e.FloorRequestsList)
	} else {
		sort.Ints(e.FloorRequestsList)
		sort.Reverse(sort.IntSlice(e.FloorRequestsList))
	}
}

//callbutton ctor
func(cb *CallButton) Init(_id int, _floor int, _direction string) {
	cb.ID = _id
	cb.Status = "off"
	cb.Floor = _floor
	cb.Direction = _direction
}

//floorrequestbutton ctor
func(fl *FloorRequestButton) Init(_id int, _floor int, _direction string) {
	fl.ID = _id
	fl.Status = "online"
	fl.Floor = _floor
	fl.Direction = _direction
}

//door ctor
func(d *Door) Init(_id int) {
	d.ID = _id
	d.Status = "closed"
}

//class global
//type Global struct {
//	CurrentIDBattery := 0
//	CurrentIDColumn := 0
//	CurrentIDElevator := 0
//	CurrentIDCallButton := 0
//	CurrentIDFloorRequestButton := 0
//	CurrentIDDoor := 0
//}

//old code
//for i:= 0; i < len(b.ColumnsList); i++ {