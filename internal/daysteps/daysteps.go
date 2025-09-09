package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
		return 0, time.Duration(0), fmt.Errorf("недостаточно данных")
	}
	//В переменную steps записываются кол-во шагов, преобразованные из строки data.
	steps, err := strconv.Atoi(splitStrings[0])
	if err != nil {
		return 0, time.Duration(0), err
	}
	if steps <= 0 {
		return 0, time.Duration(0), fmt.Errorf("кол-во шагов равно 0")
	}

	//Times переменная, полученная преобразованием входной части переменной data
	// в минуты.
	times, err := time.ParseDuration(splitStrings[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("неправильно задано время")
	}
	if times <= 0 {
		return 0, time.Duration(0), fmt.Errorf("время меньше нуля")
	}

	return steps, times, nil

}

func DayActionInfo(data string, weight, height float64) string {
	steps, times, err := parsePackage(data)
	if steps <= 0 {

		log.Println(err)
		return ""
	}
	if times <= 0 {
		log.Println(err)

		return ""
	}
	//distanceKm высчитывает, проеденное расстояние в киллометрах
	distanceKm := (float64(steps) * stepLength) / float64(mInKm)
	walkCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, times)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, walkCalories)

}
