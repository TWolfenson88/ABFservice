package core

import "github.com/TWolfenson88/ABFservice/core/wblists"

type ABFConfig struct {
	BucketConfig BucketConfig
	BucketStorage BucketStorage
	NetLists  wblists.NetLists
}