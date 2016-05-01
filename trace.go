package log

import (
	"runtime"
	"time"
)

// Traceable interface for a traceable object
type Traceable interface {
	End()
}

// TraceEntry is an object used in creating a trace log entry
type TraceEntry struct {
	start time.Time
	end   time.Time
	entry *Entry
}

// End completes the trace and logs the entry
func (t *TraceEntry) End() {
	t.end = time.Now().UTC()

	t.entry.Fields = append(t.entry.Fields,
		F("duration", Logger.durationFunc(t.end.Sub(t.start))),
		F("start", t.start.Format(Logger.timeFormat)),
		F("end", t.end.Format(Logger.timeFormat)),
	)

	if Logger.logCallerInfo {
		_, t.entry.File, t.entry.Line, _ = runtime.Caller(t.entry.calldepth)
	}

	Logger.handleEntry(t.entry)
	Logger.tracePool.Put(t)
}
