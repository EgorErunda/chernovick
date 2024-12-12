package main

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type  string
	Speed float64
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	travelTimes := make(map[string]float64)

	for _, vehicle := range vehicles {
		travelTime := vehicle.CalculateTravelTime(distance)
		switch v := vehicle.(type) {
		case Car:
			travelTimes[v.Type] = travelTime
		case Motorcycle:
			travelTimes[v.Type] = travelTime
		default:
			travelTimes["Unknown"] = travelTime
		}
	}

	return travelTimes
}
