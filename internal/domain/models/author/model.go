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
