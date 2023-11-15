package render

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var window fyne.Window

func MazeApp() {
	myApp := app.New()
	window = myApp.NewWindow("GoMaze")
	window.Resize(fyne.NewSize(470, 310))
	window.SetFixedSize(true)

	rect := canvas.NewRectangle(color.White)
	rect.Resize(fyne.NewSize(20, 20))

	image := canvas.NewImageFromResource(theme.FyneLogo())
	image.Move(fyne.NewPos(160, 0))
	image.Resize(fyne.NewSize(300, 300))

	content := container.NewWithoutLayout(initInterface(), image)

	window.SetContent(content)
	window.ShowAndRun()
}

func initInterface() *fyne.Container {
	inputHeight := widget.NewEntry()
	inputHeight.Text = "20"

	inputHeight.SetPlaceHolder("Введите высоту")

	inputWidth := widget.NewEntry()
	inputWidth.Text = "20"
	inputWidth.SetPlaceHolder("Введите длину")

	ButtonRender := widget.NewButton("Создать лабиринт", func() {
		intHeight, err := strconv.Atoi(inputHeight.Text)
		if err != nil {
			dialog.ShowInformation("Ошибка", "Введите корректную высоту", window)
			return
		}

		if (intHeight < 10) || (intHeight > 200) {
			dialog.ShowInformation("Ошибка", "Высота должна быть не меньше 10 и не больше 200", window)
			return
		}

		intWidth, err := strconv.Atoi(inputWidth.Text)
		if err != nil {
			dialog.ShowInformation("Ошибка", "Введите корректную длину", window)
			return
		}

		if (intWidth < 10) || (intWidth > 200) {
			dialog.ShowInformation("Ошибка", "Длина должна быть не меньше 10 и не больше 200", window)
			return
		}

		fmt.Println(intHeight, intWidth)
	})

	ButtonGoMaze := widget.NewButton("Пройти лабиринт", func() {

	})

	combo := widget.NewSelect([]string{"Широкий", "Узкий"}, func(value string) {
		log.Println("Select set to", value)
	})

	controllerContainer := container.NewGridWithRows(5, inputHeight, inputWidth, ButtonRender, ButtonGoMaze, combo)

	controllerContainer.Move(fyne.NewPos(0, 0))
	controllerContainer.Resize(fyne.NewSize(150, 300))

	return controllerContainer
}
