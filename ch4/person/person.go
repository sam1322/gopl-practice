package person

type Person struct {
	Name string // Exported field
	Age  int    // Exported field
	ssn  string // Unexported field
}

func NewPerson(name string, age int, ssn string) Person {
	return Person{name, age, ssn}
	//return Person{Name: name, Age: age, ssn: ssn}
}

func (p Person) GetSSN() string {
	return p.ssn
}
