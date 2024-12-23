package context

import (
    "context"
    "time"
)

// contextWithCleanup explicitly implements the context.Context interface
type contextWithCleanup struct {
    ctx     context.Context
    cleanup func()
    done    chan struct{} // Used to ensure cleanup is only performed once
}

// Deadline returns the deadline from the embedded context
func (c *contextWithCleanup) Deadline() (time.Time, bool) {
    return c.ctx.Deadline()
}

// Done returns the Done channel from the embedded context
func (c *contextWithCleanup) Done() <-chan struct{} {
    return c.done
}

// Err returns the error from the embedded context
func (c *contextWithCleanup) Err() error {
    select {
    case <-c.done:
        return context.Canceled
    default:
        return nil
    }
}

// Value retrieves a value from the embedded context
func (c *contextWithCleanup) Value(key any) any {
    return c.ctx.Value(key)
}

// watch monitors the parent context and triggers cleanup when it is canceled
func (c *contextWithCleanup) watch() {
    defer close(c.done)
    <-c.ctx.Done() // Wait for the parent context to be canceled or timed out
    if c.cleanup != nil {
        c.cleanup()
    }
}

func NewContextWithCleanup(ctx context.Context, cleanup func()) *contextWithCleanup {
    c := &contextWithCleanup{
        ctx:     ctx,
        cleanup: cleanup,
        done:    make(chan struct{}),
    }
    go c.watch() // Start monitoring the parent context in a goroutine
    return c
}
