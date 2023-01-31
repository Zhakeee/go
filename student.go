package zhake

type Student struct {
	FirstName  string
	SecondName string
	Age        int
}

func (s Student) GetFirstName() string {
	return s.FirstName

}
func (s *Student) SetFirstName(firstname string) {
	s.FirstName = firstname
}
func (s Student) GetSecondName() string {
	return s.FirstName
}
func (s *Student) SetSecondName(secondname string) {
	s.FirstName = secondname
}
func (s Student) GetAge() int {
	return s.Age

}
func (s *Student) SetAge(age int) {
	s.Age = age
}
