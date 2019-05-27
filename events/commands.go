package events

// NoOp command
type NoOp struct{}

// ImportMedia command
type ImportMedia struct {
	Path string
}

// GenerateThumbShot command
type GenerateThumbShot struct {
	ID   string
	Path string
}

// GenerateBigShot command
type GenerateBigShot struct {
	ID   string
	Path string
}

// ApproveMedia command
type ApproveMedia struct {
	ID string
}
