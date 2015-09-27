package track

import (
  "github.com/plimble/ace"
  "code.google.com/p/go-uuid/uuid"
  "github.com/influxdb/influxdb/client"
  // "time"
)

const (
  Tracker  = "_t_"
  V_ID    = "_v_id"
  Measurement = "tracking"
)

// Middleware sets Tracking cookie
func Track(i *client.Client) ace.HandlerFunc {
  return func(c *ace.C) {
    session := c.Sessions(Tracker)
    _vid := session.GetString(V_ID,"")
    if _vid == "" {
      _vid = uuid.New()
      session.Set(V_ID, _vid)
    }

    var pts = make([]client.Point, 1)

    pts[0] = client.Point{
      Measurement: Measurement,
      Tags: map[string]string{
        "event": "view",
        "vid": _vid,
      },
      Fields: map[string]interface{}{
        "value": 1,
      },
      // Time: time.Now(),
      Precision: "s",
    }

    bps := client.BatchPoints{
      Points:          pts,
      Database:      "mixevents",
      RetentionPolicy: "default",
  	}
    go i.Write(bps)
    c.Next()
  }
}
