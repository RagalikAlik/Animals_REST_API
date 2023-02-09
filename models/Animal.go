package models

import "time"

type Animal struct {
	Id                 int64     `json:"id"`
	AnimalTypes        []int64   `json:"animalTypes"`
	Weight             float32   `json:"weight"`
	Length             float32   `json:"length"`
	Height             float32   `json:"height"`
	Gender             string    `json:"gender"`
	LifeStatus         string    `json:"lifeStatus"`
	ChippingDateTime   time.Time `json:"chippingDateTime"`
	ChipperId          int       `json:"chipperId"`
	ChippingLocationId int64     `json:"chippingLocationId"`
	VisitedLocations   []int64   `json:"visitedLocations"`
	DeathDateTime      time.Time `json:"deathDateTime"`
}
