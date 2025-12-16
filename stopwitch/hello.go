package main

import (
	"fmt"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var running bool
var startTime time.Time
var elapsedTime time.Duration
var ticker *time.Ticker // 声明一个全局或可访问的 Ticker
var mu sync.Mutex

// Format time function
func formatTime(d time.Duration) string {
	// Convert duration to hours, minutes, seconds, milliseconds
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000 / 10 // show two-digit milliseconds

	return fmt.Sprintf("%02d:%02d:%02d.%02d", h, m, s, ms)
}

// Start/resume logic
func startStopwatch(label *widget.Label) {
	mu.Lock()
	if running {
		mu.Unlock()
		//加入一个弹窗提示已运行
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Info",
			Content: "Stopwatch is already running!",
		})
		return // 已经在运行
	}
	running = true
	startTime = time.Now() // 记录本次启动时间
	ticker = time.NewTicker(time.Millisecond * 10)
	mu.Unlock()

	// start a goroutine to handle timing
	go func(t *time.Ticker) {
		for range t.C {
			mu.Lock()
			if !running {
				mu.Unlock()
				return // if stopped, exit goroutine
			}
			currentDuration := elapsedTime + time.Since(startTime)
			mu.Unlock()

			// update UI on main thread
			fyne.Do(func() {
				label.SetText(formatTime(currentDuration))
			})
		}
	}(ticker)
}

func stopStopwatch() {
	mu.Lock()
	if !running {
		mu.Unlock()
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Info",
			Content: "Stopwatch is already Stop!",
		})
		return
	}
	running = false
	if ticker != nil {
		ticker.Stop() // stop the ticker
		ticker = nil
	}
	// add the duration of this run to total elapsed time
	elapsedTime += time.Since(startTime)
	mu.Unlock()
}

func resetStopwatch(label *widget.Label) {
	mu.Lock()
	if elapsedTime == 0 && !running {
		mu.Unlock()
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Info",
			Content: "Stopwatch is already Reset!",
		})
		return
	}
	running = false
	if ticker != nil {
		ticker.Stop() // stop the ticker
		ticker = nil
	}
	elapsedTime = 0
	mu.Unlock()
	fyne.Do(func() {
		label.SetText("00:00:00.00") // reset display
	})
}

func main() {
	a := app.New()
	w := a.NewWindow("Stopwatch")
	timeLabel := widget.NewLabel("00:00:00.00") // initial display time
	timeLabel.TextStyle.Monospace = true        // use monospace font for aligned digits

	startBtn := widget.NewButton("Start", func() { startStopwatch(timeLabel) })
	stopBtn := widget.NewButton("Stop", func() { stopStopwatch() })
	resetBtn := widget.NewButton("Reset", func() { resetStopwatch(timeLabel) })

	// Layout: HBox for buttons, VBox for time and button row
	buttonRow := container.NewHBox(startBtn, stopBtn, resetBtn)
	content := container.NewVBox(timeLabel, buttonRow)
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 150))
	w.ShowAndRun()
}
