//Package wblists contain methods for control black/white lists of subnets
package wblists

import (
	"errors"
	"net"
)

var ErrWrongIP = errors.New("wrong IP address")
var ErrNotFound = errors.New("not found in lists")
var ErrExistInWhiteList = errors.New("subnet already in white list")
var ErrExistInBlackList = errors.New("subnet already in black list")
var ErrDeleteNonExistNet = errors.New("subnet not exist in list")

type NetLists struct {
	WhiteList map[string]interface{}
	BlackList map[string]interface{}
}

//CheckIpInList checking given ipStr in white and black list.
//If IP founded in white list, CheckIpInList return true, nil,
//if IP founded in black list, CheckIpInList return false, nil,
//in other cases return false, error.
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

//AddSubnetToList add given subnet address into "white" or "black" list.
//If subnet already exist in one of lists, return error, else nil.
func(nl *NetLists) AddSubnetToList(netStr string, list string) error {

	_, _, err := net.ParseCIDR(netStr)
	if err != nil {
		return err
	}

	_, wOk := nl.WhiteList[netStr]
	_, bOk := nl.BlackList[netStr]

	if wOk {
		return ErrExistInWhiteList
	}
	if bOk {
		return ErrExistInBlackList
	}

	switch list {
	case "white":
		nl.WhiteList[netStr] = ""
	case "black":
		nl.BlackList[netStr] = ""
	}

	return nil
}

//RemoveSubnetFromList remove given subnet address from "white" or "black" list.
//If subnet don't exist in selected list, return error, else nil.
func(nl *NetLists) RemoveSubnetFromList(netStr string, list string) error {
	_, _, err := net.ParseCIDR(netStr)
	if err != nil {
		return err
	}

	switch list {
	case "white":
		if _, ok := nl.WhiteList[netStr]; ok {
			delete(nl.WhiteList, netStr)
		}else {
			return ErrDeleteNonExistNet
		}
	case "black":
		if _, ok := nl.BlackList[netStr]; ok {
			delete(nl.BlackList, netStr)
		}else {
			return ErrDeleteNonExistNet
		}
	}

	return nil
}
