package author

// Request model.
type Request struct {
	UserName    string `form:"name" json:"name" validate:"required"`
	Email       string `form:"email" json:"email" validate:"required,email"`
	PhoneNumber string `form:"phone-number" json:"phoneNumber" validate:"required,phone"`
	Password    string `form:"password" json:"password" validate:"required,gte=6,lte=20"`
	Bio         string `form:"biography" json:"biography"`
	Avatar      string `form:"avatar" json:"avatar"`
}

// Response model.
type Response struct {
	UserName    string `db:"username"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Password    string `db:"password"`
	Bio         string `db:"bio"`
	Avatar      string `db:"avatar"`
}

// ToPsqlDBStorage func.
func (r *Request) ToPsqlDBStorage() Response {
	return Response{
		UserName:    r.UserName,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		Password:    r.Password,
		Bio:         r.Bio,
		Avatar:      r.Avatar,
	}
}
