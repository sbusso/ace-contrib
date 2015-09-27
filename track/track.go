package track

import (
  "github.com/plimble/ace"
  "code.google.com/p/go-uuid/uuid"
)

const (
    Tracker     = "__fat"
  )
// Middleware sets Tracking cookie
func Track() ace.HandlerFunc {

  return func(c *ace.C) {
    session := c.Sessions(Tracker)
    vid := session.GetString("__vid","")
    if vid == "" {
      vid = uuid.New()
      session.Set("__vid", vid)
    }

  }
}
