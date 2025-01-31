package response

type LoginResponse struct {
	User           UserResponse   `json:"user"`
	AccessResponse AccessResponse `json:"accessResponse"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
