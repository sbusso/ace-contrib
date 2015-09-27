package track

import (
  "github.com/plimble/ace"
  "github.com/plimble/sessions/store/cookie"
  "code.google.com/p/go-uuid/uuid"
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
