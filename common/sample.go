package common

type Person struct {
	No   int
	Name string
	Age  int
}

func MakeData() map[int]*Person {
	return map[int]*Person{
		8:  {No: 8, Name: "Jack", Age: 12},
		4:  {No: 4, Name: "Dan", Age: 18},
		12: {No: 12, Name: "Lisa", Age: 24},
		2:  {No: 2, Name: "Kim", Age: 9},
		1:  {No: 1, Name: "Lee", Age: 14},
		3:  {No: 3, Name: "Park", Age: 16},
		5:  {No: 5, Name: "Moana", Age: 23},
		6:  {No: 6, Name: "Raven", Age: 35},
		7:  {No: 7, Name: "Liu", Age: 28},
		9:  {No: 9, Name: "Jinny", Age: 27},
		23: {No: 23, Name: "Jane", Age: 22},
		14: {No: 14, Name: "Jim", Age: 30},
		27: {No: 27, Name: "Ken", Age: 29},
	}
}

const (
	BANANA     = "banana"
	MISSISSIPI = "mississipi"
)
