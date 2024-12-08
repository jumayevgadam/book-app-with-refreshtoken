package author

// Request model.
type Request struct {
	UserName string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,gte=6,lte=20"`
	Bio      string `form:"biography" json:"biography"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Response model.
type Response struct {
	UserName string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Bio      string `db:"bio"`
	Avatar   string `db:"avatar"`
}

// ToPsqlDBStorage func.
func (r *Request) ToPsqlDBStorage() Response {
	return Response{
		UserName: r.UserName,
		Email:    r.Email,
		Password: r.Password,
		Bio:      r.Bio,
		Avatar:   r.Avatar,
	}
}

// AllUser model.
type Author struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"biography"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

// AllUserData model.
type AuthorData struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Bio      string `db:"bio"`
	Avatar   string `db:"avatar"`
	Password string `db:"password"`
}

// ToServer sends response to server.
func (a *AuthorData) ToServer() *Author {
	return &Author{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
		Bio:      a.Bio,
		Avatar:   a.Avatar,
		Password: a.Password,
	}
}
