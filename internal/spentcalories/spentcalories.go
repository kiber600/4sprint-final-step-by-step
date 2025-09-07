package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	//SplitData переменная хранит в себе отдельные части входных параметров
	//разделенных запятой.
	var splitData []string
	//Разделяем на части входные параметры в месте ",".
	splitData = strings.Split(data, ",")
	//Activity- переменная хранящая в себе вид активности(бег или ходьба).

	if len(splitData) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка: недостаточное количество данных")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка: Неправильно задано количество шагов")
	}
	times, err := time.ParseDuration(splitData[2])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка: Неправильно задано время")
	}
	activity := splitData[1]
	return steps, activity, times, nil

}

func distance(steps int, height float64) float64 {
	//StepLenght - переменная хранящая в себе длину шага.
	stepLenght := float64(stepLengthCoefficient * height)
	//Возвращаем значение пройденной дистанции в км.
	return (stepLenght * float64(steps)) / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	hourss := duration.Hours()
	//Возвращаем скорость в км в час.
	return distance(steps, height) / float64(hourss)

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	walkCalories, err := WalkingSpentCalories(steps, weight, height, duration)
	runCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return "", err
	}

	switch activity {
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч.\nСожгли каллорий: %.2f\n", activity, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), walkCalories), err

	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч.\nСожгли каллорий: %.2f\n", activity, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), runCalories), err

	default:
		return fmt.Sprintln("Неизвестный тип тренировки"), fmt.Errorf("Ошибка, неправильно задан тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверные данные по времени тренировки")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Нерпавильно задан вес")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверно задан рост")
	}
	if steps <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверные данные по количеству шагов")
	}
	//AveSpeed - получает данные по средней скорости из функции meanSpeed.
	speed := meanSpeed(steps, height, duration)
	//Minuts - хранит в себе время в минутах путем первода времни из переменной duration.
	minuts := duration.Minutes()
	//Callories - вычисляет кол-во затраченных каллорий.
	callories := (weight * speed * float64(minuts)) / float64(minInH)
	return callories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверные данные по времени тренировки")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Нерпавильно задан вес")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверно задан рост")
	}
	if steps <= 0 {
		return 0.0, fmt.Errorf("Ошибка: Неверные данные по количеству шагов")
	}
	//Speed - получает данные по средней скорости из функции meanSpeed.
	speed := meanSpeed(steps, height, duration)
	//Minuts - хранит в себе время в минутах путем первода времни из переменной duration.
	minuts := duration.Minutes()
	//Callories - вычисляет кол-во затраченных каллорий.
	callories := ((weight * speed * float64(minuts)) / float64(minInH)) * walkingCaloriesCoefficient
	return callories, nil
}
