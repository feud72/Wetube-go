package forms

type UserURI struct {
	ID string `uri:"id" binding:"required"`
}

type UserSignup struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
