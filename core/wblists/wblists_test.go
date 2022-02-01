package wblists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var Lists = NetLists{
	WhiteList: map[string]interface{}{"192.180.0.0/21":"", "192.180.0.0/22":"", "192.180.0.0/23":"", "192.180.0.0/20":""},
	BlackList: map[string]interface{}{"192.160.0.0/21":"", "192.160.0.0/22":"", "192.160.0.0/23":"", "192.160.0.0/20":""},
}

func TestCheckIpInList(t *testing.T) {

	t.Run("found in lists", func(t *testing.T) {

		ipAddr := "192.180.15.254"
		result, err := Lists.CheckIpInList(ipAddr)
		require.True(t, result)
		require.Equal(t, nil, err)

		ipAddr = "192.160.15.254"
		result, err = Lists.CheckIpInList(ipAddr)
		require.False(t, result)
		require.Equal(t, nil, err)
	})

	t.Run("NOT found in lists", func(t *testing.T) {
		ipAddr := "192.180.16.254"

		result, err := Lists.CheckIpInList(ipAddr)

		require.False(t, result)
		require.Equal(t, ErrNotFound, err)
	})

	t.Run("incorrect IP", func(t *testing.T) {
		ipAddr := "192.16.254"

		result, err := Lists.CheckIpInList(ipAddr)

		require.False(t, result)
		require.Equal(t, ErrWrongIP, err)
	})
}

func TestAddSubnetToList(t *testing.T) {

	t.Run("incorrect IP", func(t *testing.T) {
		subnet := "192.16.254.a/24"

		err := Lists.AddSubnetToList(subnet, "white")

		require.Error(t, err)
	})

	t.Run("successful add", func(t *testing.T) {
		subnet := "192.180.10.205/18"
		ipAddr := "192.180.10.205"

		err := Lists.AddSubnetToList(subnet, "white")
		require.Nil(t, err)

		result, err := Lists.CheckIpInList(ipAddr)
		require.True(t, result)
		require.Nil(t, err)
	})

	t.Run("error add existing subnet", func(t *testing.T) {
		subnet := "192.180.0.0/21"

		err := Lists.AddSubnetToList(subnet, "white")
		require.Error(t, err)

	})

	t.Run("error add existing white subnet to black", func(t *testing.T) {
		subnet := "192.180.0.0/21"

		err := Lists.AddSubnetToList(subnet, "black")
		require.Error(t, err)

	})

	
}

func TestRemoveSubnetFromList(t *testing.T) {

	t.Run("incorrect IP", func(t *testing.T) {
		subnet := "192.16.254.a/24"

		err := Lists.RemoveSubnetFromList(subnet, "white")

		require.Error(t, err)
	})

	t.Run("successful delete", func(t *testing.T) {
		subnet := "192.180.0.0/21"

		err := Lists.RemoveSubnetFromList(subnet, "white")
		require.Nil(t, err)
	})

	t.Run("error try to delete subnet not in list", func(t *testing.T) {
		subnet := "192.180.0.0/15"

		err := Lists.RemoveSubnetFromList(subnet, "white")
		require.Error(t, err)
	})
}
