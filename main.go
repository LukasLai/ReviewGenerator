package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Split Window Example")
	w.Resize(fyne.NewSize(300, 400))

	// 上半部分的按鈕
	button1 := createButton("1")
	button2 := createButton("2")
	button3 := createButton("3")
	button4 := createButton("4")
	button5 := createButton("5")

	// 下半部分的文本顯示
	outputLabel := widget.NewLabel("")

	// Reset 按鈕
	resetButton := widget.NewButton("Reset", func() {
		outputLabel.SetText("")
	})

	// 按鈕的點擊事件處理函數
	onClick := func(value string) func() {
		return func() {
			currentText := outputLabel.Text
			outputLabel.SetText(currentText + value)
		}
	}

	// 設置按鈕的點擊事件
	button1.OnTapped = onClick("1")
	button2.OnTapped = onClick("2")
	button3.OnTapped = onClick("3")
	button4.OnTapped = onClick("4")
	button5.OnTapped = onClick("5")

	// 上左按鈕
	LbuttonsContainer := container.NewVBox(
		button1,
		button2,
		button3,
	)
	//上右按鈕
	RbuttonsContainer := container.NewVBox(
		button4,
		button5,
		resetButton,
	)

	topContainer := container.NewHSplit(
		LbuttonsContainer,
		RbuttonsContainer,
	)

	// 設定視窗的內容為上下分割的容器
	w.SetContent(
		container.NewVSplit(
			topContainer,
			outputLabel,
		),
	)

	w.ShowAndRun()
}

// 創建按鈕函數
func createButton(label string) *widget.Button {
	button := widget.NewButton(label, nil)
	button.Importance = widget.HighImportance
	return button
}
