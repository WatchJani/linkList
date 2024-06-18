package main

import timeExpire "root/time_expire"

func main() {
	expire := timeExpire.NewTimeExpire()

	for range 10000 {
		expire.AppendData(timeExpire.RandomData())
	}

	expire.TimeValue.PrintLeftRight()
}
