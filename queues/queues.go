package queues

type Queue = []interface{}

var Queues = map[int64]Queue{}

func Push(chatId int64, item interface{}) int {
	Queues[chatId] = append(Queues[chatId], item)
	return len(Queues[chatId])
}

func Pull(chatId int64) interface{} {
	if len(Queues[chatId]) == 0 {
		return nil
	}

	toReturn := Queues[chatId][0]
	Queues[chatId] = Queues[chatId][1:]
	return toReturn
}
