package sensors

import "math/rand"

type Measure struct {
	SensorId   int
	SensorType string
	Data       Data
}

type Data struct {
	X     int
	Y     int
	Value int
}

func NewData(sensorType string) Data {

	switch sensorType {
	case "pressure":
		return Data{
			X: rand.Intn(100),
			Y: rand.Intn(100),
			//мм рт ст
			Value: rand.Intn(200) + 600,
		}
	case "humidity":
		return Data{
			X: rand.Intn(100),
			Y: rand.Intn(100),
			//влажность в процентах
			Value: rand.Intn(100),
		}
	case "activity":
		return Data{
			X: rand.Intn(100),
			Y: rand.Intn(100),
			//10-ти бальная шкала активности
			Value: rand.Intn(10),
		}
	default:
		//ну тут только ошибка хотя конечно по хорошему ее можно и возвращать тогда, но это какой-то зализанный рефакторинг случая которого не будет
		return Data{}
	}

}

func NewMeasure(sensorId int, sensorType string, data Data) Measure {
	return Measure{
		SensorId:   sensorId,
		SensorType: sensorType,
		Data:       data,
	}
}
