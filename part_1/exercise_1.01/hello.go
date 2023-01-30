package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	for {
		currentTime := time.Now().UTC()
		hashOfCurrentTime := md5.Sum([]byte(currentTime.String()))
		fmt.Println(currentTime, ":", hex.EncodeToString(hashOfCurrentTime[:]))
		time.Sleep(5 * time.Second)
	}
}
