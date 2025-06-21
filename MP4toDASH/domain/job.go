package domain

type Job struct {
	ID              string
	OutputBackePath string
	Status          string
	Video           *Video
}
