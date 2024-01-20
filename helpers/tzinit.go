package helpers

import (
  "time"
  "os"
)

func Location() *time.Location {
  location, _ := time.LoadLocation(os.Getenv("TIMEZONE"))
  return location
}