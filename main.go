package main

import (
	"fmt"
	"github.com/TWolfenson88/ABFservice/core"
	"github.com/TWolfenson88/ABFservice/core/wblists"
)

type ABFConfig struct {
	BucketConfig core.BucketConfig
	BucketStorage core.BucketStorage
	NetLists  wblists.NetLists
}

func InitABF() *ABFConfig {
	return &ABFConfig{
		BucketConfig:  core.BucketConfig{},
		BucketStorage: core.BucketStorage{},
		NetLists:      wblists.NetLists{},
	}
}

func main() {
	fmt.Println("Hello test")

	//заинитить необходимые структуры
	
	ABFcfg := InitABF()


	//постучаться в бд
	//получить конфиги на лимитер
	//получить чернобелые листы
	//закинуть данные в структуры
	//сесть слушать порт

	
}
