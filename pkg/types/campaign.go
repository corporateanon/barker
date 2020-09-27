package types

type Campaign struct {
	ID      int64
	Title   string `binding:"required"`
	Message string `binding:"required"`
	Active  bool
}
