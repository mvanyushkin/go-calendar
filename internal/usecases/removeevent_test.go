package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWhenRemovingEventExists(t *testing.T) {
	event := entities.Event{
		Id:          111,
		Title:       "",
		Description: "",
		Time:        time.Time{},
	}
	s := store.NewInMemoryEventStore().LoadFromSlice([]entities.Event{
		event,
	})

	u := RemoveEventUseCase{
		UseCase: UseCase{store: s},
		Event:   &event,
	}

	err := u.Do()
	assert.Nil(t, err)
}

func TestWhenRemovingEventDoesntExist(t *testing.T) {
	s := store.NewInMemoryEventStore()
	event := entities.Event{
		Id:          111,
		Title:       "",
		Description: "",
		Time:        time.Time{},
	}
	u := RemoveEventUseCase{
		UseCase: UseCase{store: s},
		Event:   &event,
	}

	err := u.Do()
	assert.NotNil(t, err)
}
