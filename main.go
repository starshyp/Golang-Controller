package main
import "fmt"

func main() {
	fmt.Println("Hello World")
}

//class battery
type Battery struct {
	ID									int
	Status                           string
	ColumnsList, FloorRequestButtonsList int
}

//battery ctor
  func(b *Battery) Init(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements int) {
      b.ID = id
      b.Status = "online"
      b.ColumnsList = []Column
      b.FloorRequestButtonsList = []int
  }

  //for Init(i = 0; i <= _amountOfColumns; ID++ int)
  //{
  //    Column column = new(Column(ID, 5, 20, false))
  //    column.ID = ID
  //    this.ColumnsList.Add(column)
  //}

//method to assign elevator
func AssignElevator( _requestedFloor int, _direction string) int {
	bestColumn = FindBestColumn(_requestedFloor)
 	chosenElevator = FindBestElevator(_requestedFloor, _direction)
	return bestColumn, chosenElevator
	elevator.Go()
}

//method to find best column
func FindBestColumn(_requestedFloor int) int {
	chosenColumn = 0;
	for _, chosenColumn := range b.ColumnsList {
	if b.ColumnsList.servedFloors[0] == _requestedFloor {
		chosenColumn = currentColumn;
	}
}
	return chosenColumn
})


//class column
type Column struct {
	ID int
	Status string
	ServedFloors []int
	IsBasement bool
	ElevatorsList []int
	CallButtonsList []int
}

//column ctor
func(c *Column) Init(_id int, _amountOfElevators int, _servedFloors int, _isBasement bool) {
	c.ID = id
	c.Status = "online"
	c.ServedFloors = []int
	c.IsBasement = false
	c.ElevatorsList = []int
	c.CallButtonsList = []int

	for i := 0; i < _amountOfElevators; i++ {
		var elevator int
		c.ElevatorsList.append(elevator)
	}

	for j := 0; j < ServedFloors; j++ {
		var floorButton := j;
		if (j != 1)
		{
			button = new CallButton(j, j, "down");
			c.CallButtonsList.append(button)
		}
		else if (j != _servedFloors)
		{
			button = new CallButton(j, j, "up");
			c.CallButtonsList.append(button)
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
			callButton := new (CallButton(callButtonID, buttonFloor, "down"))
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
 	elevator.FloorRequestsList.append(_floor)
 	elevator.Go();
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

