package my_modules

import "fmt"

type Student struct {
	name    string
	surname string
	age     int
	email   string
}

func Init(newName, newSurname, newEmail string, newAge int) Student{
	return Student{
		name: newName,
		surname: newSurname,
		email: newEmail,
		age: newAge,
	}
}

func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) SetAge(newAge int) {
	s.age = newAge
}

func (s *Student) GetFullName() string {
	return s.name + " " + s.surname
}

func (s *Student) GetEmail() string {
	return s.email
}

func (s Student) String() string {
	return fmt.Sprintf("[name: %s, surname: %s, age: %d, email: %s]", s.name, s.surname, s.age, s.email)
}