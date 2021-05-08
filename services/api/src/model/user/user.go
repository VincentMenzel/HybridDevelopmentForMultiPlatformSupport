package user

type User struct {
	ID       int64  `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	Password string `json:"-" xml:"-"`
}
