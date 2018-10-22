package material

// Interface for a GPU texture
type IMaterial interface {
	Bind()
	Width() int
	Height() int
}
