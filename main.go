package main
import "fmt"

func main() {
	fmt.Println("Hello World")
}

type Battery struct {
	ID, Status                           string
	ColumnsList, FloorRequestButtonsList int
}

  func(b *Battery) Init(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements int) {
      b.ID = id
      b.Status = "online"
      b.ColumnsList = []
      b.FloorRequestButtonsList = []
  }

  for (int i = 0; i <= _amountOfColumns; ID++)
  {
      Column column = new Column(ID, 5, 20, false);
      column.ID = ID
      this.ColumnsList.Add(column);
  }

//method to assign elevator
func AssignElevator(int _requestedFloor, string _direction) {
	bestColumn = FindBestColumn(_requestedFloor);
 chosenElevator = FindBestElevator(_requestedFloor, _direction);
return (bestColumn, chosenElevator);
elevator.Go();
}

//method to find best column
func FindBestColumn(_requestedFloor int) {
	for _, chosenColumn := range columnslist
	if currentColumn.servedFloors[0] == _requestedFloor {
		chosenColumn = currentColumn;
	}
})
return chosenColumn