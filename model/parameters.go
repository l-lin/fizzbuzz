package model

// Parameters used in fizz-buzz
type Parameters struct {
	Int1  int    `json:"int1" binding:"required"`
	Int2  int    `json:"int2" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

// Equal checks if the given p1 has the same values
func (p Parameters) Equal(p1 Parameters) bool {
	return p.Int1 == p1.Int1 &&
		p.Int2 == p1.Int2 &&
		p.Limit == p1.Limit &&
		p.Str1 == p1.Str1 &&
		p.Str2 == p1.Str2
}
