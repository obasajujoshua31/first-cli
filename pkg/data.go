package pkg

type User struct {
	Email    string
	Password string
}

var Users = []User{
	{
		Email:    "example@example.com",
		Password: "example",
	}, {
		Email:    "test@test.com",
		Password: "test",
	},
}
