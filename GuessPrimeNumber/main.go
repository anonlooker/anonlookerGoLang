package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Rand a number between min and max
func Rand(min, max int) int {
	if max < min {
		min, max = max, min
	}
	return min + int(math.Floor((float64(max-min+1))*randFloat64()))
}

var minVal = 2
var maxVal = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randFloat64() float64 {
	return rand.Float64()
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func generateNewNumber(lbl *widget.Label) {
	n := Rand(minVal, maxVal)
	lbl.SetText(strconv.Itoa(n))
}

func main() {
	guessPrimeNumberApp := app.New()
	mainWindow := guessPrimeNumberApp.NewWindow("Guess the Prime Number")

	instructionLabel := container.NewHBox(layout.NewSpacer(), widget.NewLabel("Is the number prime?"), layout.NewSpacer())
	numberLabel := widget.NewLabel("")

	resultLabel := widget.NewLabel("")

	yesButton := widget.NewButton("Yes", func() {
		number := atoi(numberLabel.Text)
		if IsPrime(number) {
			resultLabel.SetText("Correct! " + numberLabel.Text + " is prime.")
		} else {
			resultLabel.SetText("Wrong! " + numberLabel.Text + " is not prime.")
		}
		generateNewNumber(numberLabel)
	})
	noButton := widget.NewButton("No", func() {
		number := atoi(numberLabel.Text)
		if !IsPrime(number) {
			resultLabel.SetText("Correct! " + numberLabel.Text + " is not prime.")
		} else {
			resultLabel.SetText("Wrong! " + numberLabel.Text + " is prime.")
		}
		generateNewNumber(numberLabel)
	})

	// Settings button at top-right (icon only)
	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		settingsWindow := guessPrimeNumberApp.NewWindow("Settings")
		minEntry := widget.NewEntry()
		minEntry.SetText(strconv.Itoa(minVal))
		maxEntry := widget.NewEntry()
		maxEntry.SetText(strconv.Itoa(maxVal))

		saveBtn := widget.NewButton("Save", func() {
			minParsed := atoi(minEntry.Text)
			maxParsed := atoi(maxEntry.Text)
			if minParsed <= 0 {
				minParsed = 2
			}
			if maxParsed < minParsed {
				// swap or set to min+1
				maxParsed = minParsed + 1
			}
			minVal = minParsed
			maxVal = maxParsed
			generateNewNumber(numberLabel)
			settingsWindow.Close()
		})
		cancelBtn := widget.NewButton("Cancel", func() { settingsWindow.Close() })

		settingsContent := container.NewVBox(
			widget.NewLabel("Set min and max values"),
			widget.NewForm(
				&widget.FormItem{Text: "Min", Widget: minEntry},
				&widget.FormItem{Text: "Max", Widget: maxEntry},
			),
			container.NewHBox(saveBtn, cancelBtn),
		)
		settingsWindow.SetContent(settingsContent)
		settingsWindow.Resize(fyne.NewSize(300, 180))
		settingsWindow.Show()
	})

	// center the number label horizontally
	numberLabelContainer := container.NewHBox(layout.NewSpacer(), numberLabel, layout.NewSpacer())

	// center the Yes/No buttons horizontally and add space between them
	buttons := container.NewHBox(layout.NewSpacer(), yesButton, layout.NewSpacer(), noButton, layout.NewSpacer())
	mainContent := container.NewVBox(instructionLabel, numberLabelContainer, buttons, resultLabel)
	topBar := container.NewHBox(layout.NewSpacer(), settingsButton)
	content := container.NewBorder(topBar, nil, nil, nil, mainContent)
	mainWindow.SetContent(content)

	generateNewNumber(numberLabel)
	// ensure the window is large enough to display the full title
	mainWindow.Resize(fyne.NewSize(420, 180))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
