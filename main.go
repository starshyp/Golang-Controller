package main
import "fmt"

func main() {
	fmt.Println("Hello World")
}

//class battery
type Battery struct {
	ID int
	Status string
	ColumnsList []Column
	FloorRequestButtonsList []FloorRequestButton
	_amountOfColumns int
	_amountOfFloors int
	_amountOfBasements int
	_amountOfElevatorPerColumn int
}

//battery ctor
func(b *Battery) Init(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) {
      b.ID = _id
      b.Status = "online"
      b.ColumnsList = []Column
      b.FloorRequestButtonsList = []FloorRequestButton

	  //for i := 0; i < _amountOfColumns; i++ {
	  //	var column int
	  //	b.ColumnsList.append(column)
	  //}
  }
  //for Init(i := 0; i <= _amountOfColumns; ID++ int) {
  //   Column column = new(Column(ID, 5, 20, false))
  //   column.ID = ID
  //   this.ColumnsList.Add(column)
  //}

//method to assign elevator
func AssignElevator( _requestedFloor int, _direction string) (int, int) {
	bestColumn := FindBestColumn(_requestedFloor)
 	chosenElevator := FindBestElevator(_requestedFloor, _direction)
	return bestColumn, chosenElevator
	Go()
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

//class column
type Column struct {
	ID int
	Status string
	ServedFloors []int
	IsBasement bool
	ElevatorsList []Elevator
	CallButtonsList []CallButton
	_amountOfElevators int
}

//column ctor
func(c *Column) Init(_id int, _amountOfElevators int, _servedFloors int, _isBasement bool) {
	c.ID = _id
	c.Status = "online"
	c.ServedFloors = []int
	c.IsBasement = false
	c.ElevatorsList = []int
	c.CallButtonsList = []int

	for i := 0; i < _amountOfElevators; i++ {
		var elevator int
		append(c.ElevatorsList(elevator))
	}

	for j := 0; j < ServedFloors; j++ {
		if (j != 1) {
			button = new(CallButton(j, j, "down"))
			append(c.CallButtonsList(button))
		} else if (j != _servedFloors) {
			button = new(CallButton(j, j, "up"))
			append(c.CallButtonsList(button))
		}
	}
}
 func PushServedFloors(_amountOfColumns int) {
 	for item := 0; item < _amountOfColumns; item++ {
 		if (c.ID == 0) {
 			for i := 0; i > -6; i-- {
 				c.ServedFloors.append(i)
 				fmt.Println(i)
			}
		} else if (c.ID == 1) {
			for i := 0; i < 21; i++ {
				c.ServedFloors.append(i)
			}
		} else if (c.ID == 2) {
			c.ServedFloors.append(1)
			for i := 21; i < 41; i++ {
				c.ServedFloors.append(i)
			}
		} else if (c.ID == 3) {
			c.ServedFloors.append(1)
			for i := 41; i < 61; i++ {
				c.ServedFloors.append(i)
			}
		} for_, i := range c.ServedFloors) {
			fmt.Println(i)
		}
	}
 }

 func CreateCallButtons(_amountOfFloors int) {
 	buttonFloor := 1
 	callButtonID := 0
 	for i := 0; i < _amountOfFloors; i++ {
 		if (buttonFloor < _amountOfFloors) {
 			callButton := new(CallButton(callButtonID, buttonFloor, "up"))
 			c.CallButtonsList.append(callButton)
 			callButtonID++
		}
		if (buttonFloor > 1) {
			callButton := new(CallButton(callButtonID, buttonFloor, "down"))
			c.CallButtonsList.append(callButton)
			callButtonID++
		} else {
			buttonFloor++
		}
	}
 }

 func CreateElevators(_amountOfFloors, _amountOfElevators int) {
 	for i := 0; i < _amountOfElevators; i++ {
 		elevator := new(Elevator(elevatorID, _amountOfFloors))
 		ElevatorsList.append(elevator)
 		elevatorID++
	}
 }

 func RequestElevator(_requestedFloor int, _direction string) {
 	elevator := new(FindBestElevator(_floor, _direction))
 	FloorRequestsList.append(_floor)
 	Go();
 	fmt.Println("<OPENING DOORS")
 	return elevator
 }

 func FindBestElevator(_requestedFloor int, _direction string) int {
 	var chosenElevator int
 	if (_requestedFloor == 1) {
 		for _, elevator := range c.ElevatorsList {
 			if (elevator.CurrentFloor == 1 && elevator.Status == "stopped") {
 				chosenElevator = elevator
			}else if (elevator.CurrentFloor ==1 && elevator.Status == "idle") {
				chosenElevator = elevator
			} else if (_requestedFloor != 1) {
				chosenElevator = elevator;
			}
		}
	}
	return chosenElevator
 }

//class elevator
type Elevator struct {
	ID int
	Status string
	CurrentFloor []int
	Direction string
	Door Door
	FloorRequestsList []int
	CompletedRequestsList []int
}

//elevator ctor
func(e *Elevator) Init(_id int) {
	e.ID = _id
	e.Status = "idle"
	e.CurrentFloor = 1
	e.Direction = null
	e.Door = new(Door)
	e.FloorRequestsList = []int
	e.CompletedRequestsList = []int
}

func CreateFloorRequestButtons(_amountOfFloors int) {
	buttonFloor := 1
	var direction string
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton = new(FloorRequestButton(ID, buttonFloor, direction))
		FloorRequestsList.append(i)
		buttonFloor++
		ID++
	}
}

func RequestFloor(floor int) {
	e.FloorRequestsList.append(floor)
	Go()
}

func Go() {
	for true len(e.FloorRequestsList != 0) {
		destination := e.FloorRequestsList[0]
		e.Status = "moving"
		if (e.CurrentFloor < destination) {
			e.Direction = "up"
			for (e.CurrentFloor < destination) {
				fmt.Println("Elevator #" + e.ID + " is now at floor " + e.CurrentFloor)
				e.CurrentFloor++
			}
		} else if (e.CurrentFloor > destination) {
			e.Direction = "down"
			for (e.CurrentFloor > destination) {
				e.CurrentFloor--
			}
		}
		e.Status = "stopped"
		//e.FloorRequestsList.
	}
	e.Status = "idle"
}

//class callbutton
type CallButton struct {
	ID int
	Status string
	Floor int
	Direction string
}

//callbutton ctor
func(cb *CallButton) Init(_id int, _floor int, _direction string) {
	cb.ID = _id
	cb.Status = "off"
	cb.Floor = _floor
	cb.Direction = _direction
}

//class floorrequestbutton
type FloorRequestButton struct {
	ID int
	Status string
	Floor int
	Direction string
}

//floorrequestbutton ctor
func(fl *FloorRequestButton) Init(_id int, _floor int, _direction string) {
	fl.ID = _id
	fl.Status = "off"
	fl.Floor = _floor
	fl.Direction = _direction
}

//class door
type Door struct {
	ID int
	Status string
}

//door ctor
func(d *Door) Init(_id int) {
	d.ID = _id
	d.Status = "off"
}

//class global
type Global struct {
	CurrentIDBattery := 0
	CurrentIDColumn := 0
	CurrentIDElevator := 0
	CurrentIDCallButton := 0
	CurrentIDFloorRequestButton := 0
	CurrentIDDoor := 0
}

//old code
//for i:= 0; i < len(b.ColumnsList); i++ {