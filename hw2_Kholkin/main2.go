package calculator

type Squarer interface {
    Square() float64
}

func TotalArea(arr []Squarer) float64 {
        var total float64
        for key := range arr {
                total += arr[key].Square()
        }
        return total
}
