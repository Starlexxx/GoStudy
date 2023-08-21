package users

type Customer struct {
	age int
}

func (c Customer) Age() int {
	return c.age
}

type Employee struct {
	age int
}

func (e Employee) Age() int {
	return e.age
}

type Age interface {
	Age() int
}

func MaxAge(users []Age) int {
	max := 0
	for _, user := range users {
		if user.Age() > max {
			max = user.Age()
		}
	}

	return max
}
