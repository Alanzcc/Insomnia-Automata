// Connects to MongoDB and sets a Stable API version
package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Replace the placeholder with your Atlas connection string

type Automaton struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Creation    time.Time          `bson:"creation,omitempty"`
	Workflow    string             `bson:"workflow,omitempty"`
	Status      string
	Logs        []string
}
