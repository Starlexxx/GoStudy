package methodless_users

type Customer struct {
	age int
}

type Employee struct {
	age int
}

func Oldest(users []interface{}) (oldest interface{}) {
	maxAge := 0
	for _, user := range users {
		switch user := user.(type) {
		case Customer:
			if user.age > maxAge {
				oldest = user
				maxAge = user.age
			}
		case Employee:
			if user.age > maxAge {
				oldest = user
				maxAge = user.age
			}
		}
	}

	return oldest
}
