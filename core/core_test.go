package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var Lists = NetLists{
	WhiteList: map[string]interface{}{"192.180.0.0/21":"", "192.180.0.0/22":"", "192.180.0.0/23":"", "192.180.0.0/20":""},
	BlackList: map[string]interface{}{"192.160.0.0/21":"", "192.160.0.0/22":"", "192.160.0.0/23":"", "192.160.0.0/20":""},
}

func TestCheckIpInList(t *testing.T) {

	//WhiteList := []string{"192.180.0.0/21", "192.180.0.0/22", "192.180.0.0/23", "192.180.0.0/20"}
	//BlackList := []string{"192.160.0.0/21", "192.160.0.0/22", "192.160.0.0/23", "192.160.0.0/20"}

	t.Run("found in lists", func(t *testing.T) {
		ipAddr := "192.180.15.254"

		result, err := Lists.CheckIpInList(ipAddr)

		require.True(t, result)
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
		ipAddr := "192.16.254.a/24"

		err := Lists.AddSubnetToList(ipAddr)

		require.Error(t, err)
	})

	//WhiteList := []string{"192.180.0.0/21", "192.180.0.0/22", "192.180.0.0/23", "192.180.0.0/20"}
	//BlackList := []string{"192.160.0.0/21", "192.160.0.0/22", "192.160.0.0/23", "192.160.0.0/20"}
	
	//t.Run("successful added", func(t *testing.T) {
	//
	//	subnet := "192.168.10.205/24"
	//
	//	require.Nil(t, AddSubnetToList(subnet, WhiteList, BlackList))
	//})

	//t.Run("exist in another list", func(t *testing.T) {
	//
	//	subnet := "192.160.0.0/21"
	//
	//	require.Error(t, ErrFoundInAnotherList, AddSubnetToList(subnet, WhiteList, BlackList))
	//})

	//t.Run("exist in current list", func(t *testing.T) {
	//
	//	subnet := "192.180.0.0/21"
	//
	//	require.Error(t, ErrAleradyInList, AddSubnetToList(subnet, WhiteList, BlackList))
	//})
	
}
