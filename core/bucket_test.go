package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestLoginBucket(t *testing.T) {
//	t.Run("test log bucket", func(t *testing.T) {
//		login := "KEK"
//
//		BC := BucketConfig{
//			LogLimit:  10,
//			PassLimit: 100,
//			IPLimit:   1000,
//		}
//
//		for i := 0; i < 10; i++ {
//			result := BC.LoginBucket(login)
//
//			require.True(t, result)
//		}
//	})
//}

var BC = BucketConfig{
	LogLimit:  10,
	PassLimit: 100,
	IPLimit:   1000,
}

func TestBucketConfig_LoginBucket(t *testing.T) {
	t.Run("test log bucket", func(t *testing.T) {
		login := "KEK"

		for i := 0; i < 10; i++ {
			result := BC.LoginBucket(login)

			require.True(t, result)
		}
	})
}

func TestBucketConfig_PassBucket(t *testing.T) {
	t.Run("test pass bucket", func(t *testing.T) {
		pass := "supersecret"

		for i := 0; i < 100; i++ {
			result := BC.PassBucket(pass)

			require.True(t, result)
		}
	})
}

func TestBucketConfig_IPBucket(t *testing.T) {
	t.Run("test pass bucket", func(t *testing.T) {
		pass := "191.111.0.12"

		for i := 0; i < 1000; i++ {
			result := BC.IPBucket(pass)

			require.True(t, result)
		}
	})
}
