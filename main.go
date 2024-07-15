package main

import (
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("PostPerson")

    urlEntry := widget.NewEntry()
    urlEntry.SetPlaceHolder("Enter URL here")
    methodSelect := widget.NewSelect([]string{"GET", "POST"}, func(value string) {})

    responseLabel := widget.NewLabel("Response will be shown here")

    sendButton := widget.NewButton("Send", func() {
        url := urlEntry.Text
        method := methodSelect.Selected

        resp, err := sendRequest(method, url)
        if err != nil {
            responseLabel.SetText(fmt.Sprintf("Error: %v", err))
            return
        }

        responseLabel.SetText(resp)
    })

    myWindow.SetContent(container.NewVBox(
        widget.NewLabel("GoPostman"),
        urlEntry,
        methodSelect,
        sendButton,
        responseLabel,
    ))

    myWindow.ShowAndRun()
}

func sendRequest(method, url string) (string, error) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return "", err
    }

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
