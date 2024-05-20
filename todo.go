package models

import "time"

// ToDoList struct defines the structure of a TO-DO list
type ToDoList struct {
    ID                int       `json:"id"`
    CreationDate      time.Time `json:"creation_date"`
    ModificationDate  time.Time `json:"modification_date"`
    DeletionDate      *time.Time `json:"deletion_date,omitempty"`
    CompletionPercent int       `json:"completion_percent"`
}

// ToDoMessage struct defines the structure of a TO-DO message
type ToDoMessage struct {
    ID               int       `json:"id"`
    ToDoListID       int       `json:"todo_list_id"`
    CreationDate     time.Time `json:"creation_date"`
    ModificationDate time.Time `json:"modification_date"`
    DeletionDate     *time.Time `json:"deletion_date,omitempty"`
    Content          string    `json:"content"`
    IsCompleted      bool      `json:"is_completed"`
}
