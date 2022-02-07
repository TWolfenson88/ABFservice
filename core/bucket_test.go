package core

import (
	"github.com/TWolfenson88/ABFservice/core/wblists"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

var BC = BucketConfig{
	LogLimit:  10,
	PassLimit: 100,
	IPLimit:   1000,
}

var BS = BucketStorage{
	Logins: make(map[string]*Login),
	Passes: make(map[string]*Password),
	IPs:    make(map[string]*IP),
	mu:     sync.Mutex{},
}

func TestBucketConfig_LoginBucket(t *testing.T) {
	t.Run("test log bucket", func(t *testing.T) {
		login := "KEK"

		for i := 0; i < 10; i++ {
			result := BS.LoginBucket(login, BC.LogLimit)

			require.True(t, result)
		}
	})

	t.Run("test reject login", func(t *testing.T) {
		login := "KEK"

		result := BS.LoginBucket(login, BC.LogLimit)

		require.False(t, result)

	})
}

func TestBucketConfig_PassBucket(t *testing.T) {
	t.Run("test pass bucket", func(t *testing.T) {
		pass := "supersecret"

		for i := 0; i < 100; i++ {
			result := BS.PassBucket(pass, BC.PassLimit)

			require.True(t, result)
		}
	})
	t.Run("test reject pass", func(t *testing.T) {
		pass := "supersecret"

		result := BS.PassBucket(pass, BC.PassLimit)

		require.False(t, result)

	})
}

func TestBucketConfig_IPBucket(t *testing.T) {
	t.Run("test IP bucket", func(t *testing.T) {
		pass := "191.111.0.12"

		for i := 0; i < 1000; i++ {
			result := BS.IPBucket(pass, BC.IPLimit)

			require.True(t, result)
		}
	})
	t.Run("test reject IP", func(t *testing.T) {
		pass := "191.111.0.12"

		result := BS.IPBucket(pass, BC.IPLimit)

		require.False(t, result)

	})
}

func TestLimit(t *testing.T) {

	prms := RequestParams{
		Login:    "rere",
		Password: "eeee",
		IPaddr:   "192.168.10.205",
	}
	
	lists := wblists.NetLists{
		WhiteList: map[string]interface{}{"192.180.0.0/20":nil},
		BlackList: map[string]interface{}{"192.170.0.0/20":nil},
	}

	t.Run("test pass limit", func(t *testing.T) {

		var result bool

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC, lists)
		}

		require.True(t, result)
	})

	t.Run("test reject limit", func(t *testing.T) {

		var result bool

		result, _ = BS.Limit(prms, BC, lists)

		require.False(t, result)
	})

	t.Run("allow white IP", func(t *testing.T) {

		var result bool

		prms.IPaddr = "192.180.14.200"

		for i := 0; i < 2000; i++ {
			result, _ = BS.Limit(prms, BC, lists)
		}

		require.True(t, result)
	})

	t.Run("reject black IP", func(t *testing.T) {

		var result bool

		prms.IPaddr = "192.170.14.200"

		for i := 0; i < 2000; i++ {
			result, _ = BS.Limit(prms, BC, lists)
		}

		require.False(t, result)
	})
}

func TestResetBuckets(t *testing.T) {

	lists := wblists.NetLists{
		WhiteList: nil,
		BlackList: nil,
	}

	t.Run("test bucket reseting", func(t *testing.T) {

		var result bool

		prms := RequestParams{
			Login:    "rere",
			Password: "eeee",
			IPaddr:   "192.168.10.205",
		}

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC, lists)
		}

		require.False(t, result)

		BS.ResetBuckets(prms)

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC, lists)
		}
		require.True(t, result)
	})
}
