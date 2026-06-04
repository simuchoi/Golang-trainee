 package user

 type User struct {
	name string // приватное поле (доступно только внутри пакета user)
	Age int // публичное поле (доступно везде и всем)
 }


 func NewUser(name string, age int) User { // публичный конструктор (начинается с большой буквы)
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	return User {
		name: name,
		Age: age,
	}
 }


 func (u *User) SetNewName (name string) { // публичный сеттер для имени
	if name != "" {
		u.name = name
	}
 }


 func (u *User) SetNewAge (age int) { // публичный сеттер для возраста
	if age > 0 && age <= 150 {
		u.Age = age
	}
 }


 func (u *User) GetName() string { // публичный геттер для имени
	return u.name
 }
 
 
 func (u *User) GetAge() int { // публичный геттер для возраста
	return u.Age
 }