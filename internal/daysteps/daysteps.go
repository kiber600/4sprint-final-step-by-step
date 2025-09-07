package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	// SplitStrings переменная хранит в себе отдельные части входных параметров
	//разделенных запятой.
	var splitStrings []string
	//Разделяем на части входные параметры в месте ",".
	splitStrings = strings.Split(data, ",")

	if len(splitStrings) != 2 {
		return 0, time.Duration(0), fmt.Errorf("Ошибка: недостаточно данных")
	}
	//В переменную steps записываются кол-во шагов, преобразованные из строки data.
	steps, err := strconv.Atoi(splitStrings[0])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка: неправильно заданно кол-во шагов")
	}
	if steps <= 0 {
		return 0, time.Duration(0), fmt.Errorf("Ошибка: Кол-во шагов равно 0")
	}

	//Times переменная, полученная преобразованием входной части переменной data
	// в минуты.
	times, err := time.ParseDuration(splitStrings[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка: неправильно задано время")
	}

	return steps, times, nil

}

func DayActionInfo(data string, weight, height float64) string {
	steps, times, _ := parsePackage(data)
	if steps <= 0 {
		fmt.Println("Ошибка: кол-во шагов равно 0")
		return ""
	}
	if times <= 0 {
		fmt.Println("Ошибка: неправильно задано время. Время меньше или равно 0")
		return ""
	}
	//distanceKm высчитывает, проеденное расстояние в киллометрах
	distanceKm := (float64(steps) * stepLength) / float64(mInKm)
	walkCalories, _ := WalkingSpentCalories(steps, weight, height, times)
	return fmt.Sprintf("Количество шагов:%d\nДистанция составила:%.2f\nВы сожгли:%2.f", steps, distanceKm, walkCalories)

}
