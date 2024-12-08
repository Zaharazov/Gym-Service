package database

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GenerateGraph(graphPath string, userCount, adminCount, coachCount int) error {
	// Создаём новый график
	p := plot.New()

	p.Title.Text = "Количество пользователей"
	p.X.Label.Text = "Роли"
	p.Y.Label.Text = "Количество"

	// Данные для графика
	values := plotter.Values{float64(userCount), float64(adminCount), float64(coachCount)}

	// Создаём столбчатую диаграмму
	bar, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		return err
	}

	// Подписываем оси
	p.NominalX("Пользователи", "Админы", "Тренеры")
	bar.Color = color.RGBA{R: 100, G: 150, B: 255, A: 255}
	p.Add(bar)

	// Сохраняем график как изображение
	if err := p.Save(6*vg.Inch, 4*vg.Inch, graphPath); err != nil {
		return err
	}

	return nil
}
