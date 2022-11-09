package todo


type User struct {
	Id int	`json:"-" db:"id"`
	Name int	`json:"name" binding:"required"`
	Username string	`json:"username" binding:"required"`
	Password string	`json:"password" binding:"required"`
}