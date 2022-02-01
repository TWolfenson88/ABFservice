package core

import (
	"errors"
	"net"
)

var ErrWrongIP = errors.New("wrong IP address")
var ErrNotFound = errors.New("not found in lists")
var ErrFoundInAnotherList = errors.New("found in another list")
var ErrAleradyInList = errors.New("subnet already in list")

type NetLists struct {
	WhiteList map[string]interface{}
	BlackList map[string]interface{}
}

func(nl *NetLists) CheckIpInList(ipStr string) (bool, error)  {
	var IPaddr net.IP
	var found = false

	if IPaddr = net.ParseIP(ipStr); IPaddr == nil {
		return false, ErrWrongIP
	}


		for s, _ := range nl.WhiteList {
			if _, ipNet, err := net.ParseCIDR(s); err != nil {
				return false, err
			}else if ipNet.Contains(IPaddr){
				return true, nil
			}
		}

		for s, _ := range nl.BlackList {
			if _, ipNet, err := net.ParseCIDR(s); err != nil {
				return false, err
			}else if ipNet.Contains(IPaddr){
				found = true
				break
			}
		}
	if found {
		return false, nil
	}
	return false, ErrNotFound
}

func(nl *NetLists) AddSubnetToList(netStr string) error {

	_, _, err := net.ParseCIDR(netStr)
	if err != nil {
		return err
	}

	//for _, s := range WhiteList {
	//	if netStr == s {
	//		return ErrAleradyInList
	//	}
	//}
	//
	//for _, s := range BlackList {
	//	if netStr == s {
	//		return ErrFoundInAnotherList
	//	}
	//}

	return nil
}
