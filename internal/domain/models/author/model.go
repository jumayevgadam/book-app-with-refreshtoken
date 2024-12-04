package author

// Request model.
type Request struct {
	Name string `form:"name" json:"name" validate:"required"`
}
