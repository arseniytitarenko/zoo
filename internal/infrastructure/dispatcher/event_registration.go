package dispatcher

import (
	"fmt"
	"time"
	"zoo/internal/domain"
)

func RegisterEventHandlers(d *EventDispatcher) {
	d.Register(domain.AnimalMoved, func(e domain.Event) {
		event := e.(domain.AnimalMovedEvent)
		fmt.Printf("%v: 📦 Animal %s перемещено из %s в %s\n",
			event.OccurredAt.Format(time.RFC822),
			event.AnimalID,
			event.FromEnclosureID,
			event.ToEnclosureID)
	})
	d.Register(domain.FeedingTime, func(e domain.Event) {
		event := e.(domain.FeedingTimeEvent)
		fmt.Printf("%v: 🍽 Животное %s покормлено (%s) — запланировано на %v\n",
			event.OccurredAt.Format(time.RFC822),
			event.AnimalID,
			event.FoodType,
			event.ScheduledAt.Format(time.RFC822))
	})
}
