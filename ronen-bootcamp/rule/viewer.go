package rule

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	AdminToken      = "1234"
	CtxKey          = "viewer-context"
	_          Role = 1 << iota
	Admin
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool // If viewer is admin.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	Role Role // Attached roles.
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

// fromContext returns the Viewer stored in a context.
func fromContext(ctx context.Context) Viewer {
	v, ok := ctx.Value(CtxKey).(Viewer)
	if !ok {
		return nil
	}
	return v
}

// NewContext returns a new context with the given viewer attached.
func NewContext(ctx *gin.Context, token string) *gin.Context {
	if token == AdminToken {
		ctx.Set(CtxKey, UserViewer{Role: Admin})
	}
	return ctx
}
