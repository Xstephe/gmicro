package code

//go:generate codeg -type=int -doc -output ./error_code_generated.md
const (
	// ErrInventoryNotFound - 404: Inventory not found.
	ErrInventoryNotFound int = iota + 100601

	// ErrInvSellDetailNotFound - 400: Inventory sell detail not found.
	ErrInvSellDetailNotFound

	// ErrInvNotEnough - 400: Inventory not enough.
	ErrInvNotEnough
)
