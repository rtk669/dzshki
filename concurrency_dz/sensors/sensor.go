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
	return Data{
		X:     rand.Intn(100),
		Y:     rand.Intn(100),
		Value: rand.Intn(1000),
	}
}

func NewMeasure(sensorId int, sensorType string, data Data) Measure {
	// для каждого sensorType придумать свою обработку value
	return Measure{
		SensorId:   sensorId,
		SensorType: sensorType,
		Data:       data,
	}
}
