package timeExpire

import (
	"math/rand"
	skipList "root/skip_list"
	"time"
)

type DataBlock struct {
	Data    *[]byte
	Pointer *skipList.Node[string, DataBlock]
}

type TimeExpire struct {
	RealValue map[string]*DataBlock
	TimeValue *skipList.SkipList[string, DataBlock]
}

func NewTimeExpire() *TimeExpire {
	return &TimeExpire{
		RealValue: make(map[string]*DataBlock),
		TimeValue: skipList.NewSkipList[string, DataBlock](32, 0.6),
	}
}

func (t *TimeExpire) AppendData(data []byte) {
	timeNow := time.Now().Add(time.Duration(rand.Intn(60)) * time.Minute).String()

	//create time structure
	pointer := t.TimeValue.Add(timeNow, DataBlock{
		Data: &data,
	})

	//add to real data structure
	t.RealValue[string(data[:5])] = &DataBlock{
		Data: &data,
	}

	//update Pointer
	pointer.Value.Pointer = pointer
}

func RandomData() []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, 120)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return bytes
}
