package logging

import "time"

type PingLogMessage struct {
	Message string    `bson:"Message"`
	Date    time.Time `bson:"Date"`
}
