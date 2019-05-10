package redis

import "testing"

var (
	addr = "127.0.0.1:6389"
	pwd  = "foobared"
	db   = 0
)

func TestNewPool(t *testing.T) {
	rpool := newPool(addr, pwd, db)

	err := rpool.Get().Err()
	if err != nil {
		t.Error(" error new pool,err:", err)
	}
}
