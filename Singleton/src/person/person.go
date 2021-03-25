package person

type Person struct {
	name string
	age int
}

func NewPerson() *Person {
	return &Person{}
}

func (p *Person) SetName(name string)  {
	p.name = name
}

func (p *Person) SetAge(age int)  {
	p.age = age
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}
