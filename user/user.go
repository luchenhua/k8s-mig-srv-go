package user

type User struct {
	ID        int    `form:"id,omitempty" json:"id,omitempty" gorm:"primaryKey"`
	UUID      string `form:"uuid,omitempty" json:"uuid,omitempty"`
	FirstName string `form:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string `form:"last_name,omitempty" json:"last_name,omitempty"`
	Gender    string `form:"gender,omitempty" json:"gender,omitempty"`
	Mobile    string `form:"mobile,omitempty" json:"mobile,omitempty"`
	Email     string `form:"email,omitempty" json:"email,omitempty"`
	OptIn     bool   `form:"opt_in,omitempty" json:"opt_in,omitempty"`
	IsDeleted bool   `form:"is_deleted,omitempty" json:"is_deleted,omitempty"`
}
