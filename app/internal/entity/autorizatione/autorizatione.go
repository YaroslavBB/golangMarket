package autorizatione

type User struct {
	// UserID   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
