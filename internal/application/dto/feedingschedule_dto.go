package dto

import (
	"github.com/google/uuid"
)

type NewFeedingScheduleRequest struct {
	AnimalID    string `json:"animal_id"`
	FoodType    string `json:"food_type"`
	ScheduledAt string `json:"scheduled_at"`
}

type FeedingScheduleResponse struct {
	ID          uuid.UUID `json:"id"`
	AnimalID    uuid.UUID `json:"animal_id"`
	FoodType    string    `json:"food_type"`
	ScheduledAt string    `json:"scheduled_at"`
	OccurredAt  string    `json:"occurred_at"`
}
