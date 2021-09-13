package user

type User struct {
	Id        int    `json:"id"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Gender    string `json:"gender"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	OptIn     bool   `json:"opt_in"`
	IsDeleted bool   `json:"is_deleted"`
}
