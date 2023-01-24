package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"playground/ronen-bootcamp/ent"
	"playground/ronen-bootcamp/ent/enttest"
	"playground/ronen-bootcamp/ent/hook"
	"playground/ronen-bootcamp/rule"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var authorizedCtx = context.WithValue(context.Background(), rule.CtxKey, rule.UserViewer{Role: rule.Admin})

func TestCreateUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	srv := newServer(client)
	client.User.Use(hook.On(createUserMut, ent.OpCreate))

	w := httptest.NewRecorder()
	body := `{"name": "rotemtam", "email": "rotem@entgo.io", "balance": 100}`
	req, _ := http.NewRequest("POST", "/v1/user", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	// test audit
	audit, err := client.Audit.Get(authorizedCtx, 1)
	require.NoError(t, err)
	require.Equal(t, "1234", audit.Identity)
	require.Equal(t, "user rotemtam created successfully by 1234", audit.Description)

	// bad path: user already exist in DB
	resp := httptest.NewRecorder()
	srv.ServeHTTP(resp, req)
	require.Equal(t, 500, resp.Code)
	audit, err = client.Audit.Get(authorizedCtx, 2)
	require.Error(t, err)

	// bad path: unauthorized context
	audit, err = client.Audit.Get(context.Background(), 1)
	require.Error(t, err)
}

func TestUpdateBalance(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	user := createTestUser(client)

	srv := newServer(client)
	client.User.Use(hook.On(updateUserMut, ent.OpUpdateOne))

	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/user/%d/balance", user.ID)
	body := `{"delta": -300}`
	req, _ := http.NewRequest("PATCH", url, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	var payload struct {
		Balance float64 `json:"balance"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &payload)
	require.NoError(t, err)
	require.EqualValues(t, -200, payload.Balance)

	// test audit
	audit, err := client.Audit.Get(authorizedCtx, 1)
	require.NoError(t, err)
	require.Equal(t, "1234", audit.Identity)
	require.Equal(t, "balance updated successfully for user rotem by 1234, balance now: -200", audit.Description)

}

func TestTimeTravel(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	user := createTestUser(client)

	srv := newServer(client)
	client.User.Use(hook.On(createUserMut, ent.OpCreate),
		hook.On(updateUserMut, ent.OpUpdateOne))

	// set balance to -100 (100 - 200) in time t1
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/user/%d/balance", user.ID)
	body := `{"delta": -200}`
	req, _ := http.NewRequest("PATCH", url, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	t1 := time.Now()

	time.Sleep(time.Second)
	// set balance to 400 (-100 + 500) in time t2
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/user/%d/balance", user.ID)
	body = `{"delta": 500}`
	req, _ = http.NewRequest("PATCH", url, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	t2 := time.Now()

	// test time t1
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/user/%d/travel", user.ID)
	body = fmt.Sprintf("{\"timestamp\": \"%s\"}", t1.Format(jsonTimeFormat))
	req, _ = http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	var payload struct {
		BalanceInTimePoint float64 `json:"balanceInTimePoint"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &payload)
	require.NoError(t, err)
	require.EqualValues(t, -100, payload.BalanceInTimePoint)

	// test time t2
	w = httptest.NewRecorder()
	body = fmt.Sprintf("{\"timestamp\": \"%s\"}", t2.Format(jsonTimeFormat))
	req, _ = http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer 1234")
	srv.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &payload)
	require.NoError(t, err)
	require.EqualValues(t, 400, payload.BalanceInTimePoint)

	// test bad path: unauthorized context
	w = httptest.NewRecorder()
	body = fmt.Sprintf("{\"timestamp\": \"%s\"}", t2.Format(jsonTimeFormat))
	req, _ = http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Authorization", "unauthorized")
	srv.ServeHTTP(w, req)
	require.Equal(t, 500, w.Code)

}

func createTestUser(client *ent.Client) *ent.User {
	user := client.User.Create().
		SetName("rotem").
		SetEmail("rotem@entgo.io").
		SetBalance(100).
		SaveX(authorizedCtx)
	return user
}
