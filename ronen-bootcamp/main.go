package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"playground/ronen-bootcamp/ent"
	"playground/ronen-bootcamp/ent/audit"
	"playground/ronen-bootcamp/ent/hook"
	"playground/ronen-bootcamp/ent/user"
	"playground/ronen-bootcamp/rule"
	"strconv"
	"strings"
	"time"

	_ "playground/ronen-bootcamp/ent/runtime"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	bearerKey      = "bearerKey"
	jsonTimeFormat = "2006-01-02T15:04:05Z07:00"
)

var createUserMut = func(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
		tx, err := m.Tx()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		v, err := next.Mutate(ctx, m)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, _ := m.ID()
		usr := m.Client().User.GetX(ctx, id)
		identity := fmt.Sprintf("%v", ctx.Value(bearerKey))
		desc := fmt.Sprintf("user %s created successfully by %s", usr.Name, identity)
		if _, err = tx.Audit.Create().
			SetIdentity(identity).
			SetTimestamp(time.Now()).
			SetDescription(desc).
			SetBalance(usr.Balance).
			SetUser(usr).
			Save(ctx); err != nil {
			tx.Rollback()
		}
		return v, nil
	})
}
var updateUserMut = func(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
		tx, err := m.Tx()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		v, err := next.Mutate(ctx, m)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, _ := m.ID()
		usr := m.Client().User.GetX(ctx, id)
		identity := fmt.Sprintf("%v", ctx.Value(bearerKey))
		desc := fmt.Sprintf("balance updated successfully for user %s by %s, balance now: %v", usr.Name, identity, usr.Balance)
		if _, err = tx.Audit.Create().
			SetIdentity(identity).
			SetTimestamp(time.Now()).
			SetDescription(desc).
			SetBalance(usr.Balance).
			SetUser(usr).
			Save(ctx); err != nil {
			tx.Rollback()
		}
		return v, nil
	})
}

func main() {
	// Create a new ent client with in-memory database
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client.User.Use(
		hook.On(createUserMut, ent.OpCreate),
		hook.On(updateUserMut, ent.OpUpdateOne),
	)

	if err != nil {
		log.Panicf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}
	// Start the server
	if err := newServer(client).Run(); err != nil {
		log.Panicf("failed: %v: ", err)
	}

}

type server struct {
	client *ent.Client
	*gin.Engine
}

// newServer creates a new ent-bank server.
func newServer(client *ent.Client) *server {
	r := gin.Default()
	s := &server{client: client, Engine: r}
	r.Use(extractBearer)
	r.POST("/v1/user", s.createUser)
	r.PATCH("/v1/user/:id/balance", s.updateBalance)
	r.POST("/v1/user/:id/travel", s.TimeTravel)
	return s
}

// createUser creates a new user.
func (s *server) createUser(c *gin.Context) {
	var payload ent.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx, err := s.client.Tx(c)
	usr, err := tx.User.Create().
		SetName(payload.Name).
		SetEmail(payload.Email).
		SetBalance(payload.Balance).
		Save(c)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	log.Printf("user created successfully by %s", c.Value(bearerKey))
	c.JSON(http.StatusOK, gin.H{"id": usr.ID})
}

// updateBalance updates the balance of a User.
func (s *server) updateBalance(c *gin.Context) {
	var payload struct {
		Delta float64 `json:"delta"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx, err := s.client.Tx(c)
	usr, err := tx.User.UpdateOneID(id).
		AddBalance(payload.Delta).
		Save(c)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"balance": usr.Balance})
}

// extractBearer extracts the authorization token from the `Authorization` header and
// places it on the *gin.Context. .
func extractBearer(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Set(bearerKey, token)
	// set privacy rules according to the authorization token
	c = rule.NewContext(c, token)
	c.Next()
}

// TimeTravel retrieve the state of user record in given timestamp
func (s *server) TimeTravel(c *gin.Context) {
	var payload struct {
		Timestamp time.Time `json:"timestamp"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auditHistory, err := s.client.User.
		Query().
		Where(user.ID(id)).
		QueryAudits().
		Where(audit.TimestampGTE(payload.Timestamp)).
		Order(ent.Asc(audit.FieldTimestamp)).
		First(c) // return the first audit record connected to user id with timestamp greater or equal to given time
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balanceInTimePoint": auditHistory.Balance})
}
