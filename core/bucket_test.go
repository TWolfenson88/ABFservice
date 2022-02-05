package core

import (
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

		for i := 0; i < 5; i++ {
			result := BS.LoginBucket(login, BC.LogLimit)

			require.True(t, result)
		}
	})
}

func TestBucketConfig_PassBucket(t *testing.T) {
	t.Run("test pass bucket", func(t *testing.T) {
		pass := "supersecret"

		for i := 0; i < 50; i++ {
			result := BS.PassBucket(pass, BC.PassLimit)

			require.True(t, result)
		}
	})
}

func TestBucketConfig_IPBucket(t *testing.T) {
	t.Run("test pass bucket", func(t *testing.T) {
		pass := "191.111.0.12"

		for i := 0; i < 500; i++ {
			result := BS.IPBucket(pass, BC.IPLimit)

			require.True(t, result)
		}
	})
}

func TestLimit(t *testing.T) {
	t.Run("test limit", func(t *testing.T) {

		var result bool

		prms := RequestParams{
			Login:    "rere",
			Password: "eeee",
			IPaddr:   "keck",
		}

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC)
		}

		require.True(t, result)
	})
}

func TestResetBuckets(t *testing.T) {
	t.Run("test bucket reseting", func(t *testing.T) {

		var result bool

		prms := RequestParams{
			Login:    "rere",
			Password: "eeee",
			IPaddr:   "keck",
		}

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC)
		}

		require.False(t, result)
		//fmt.Println(result, logins)

		BS.ResetBuckets(prms)

		for i := 0; i < 10; i++ {
			result, _ = BS.Limit(prms, BC)
		}
		//fmt.Println(result, logins)
		require.True(t, result)
	})
}
