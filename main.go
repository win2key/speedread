package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Display words</title>
    <style>
        #display {
            font-family: Arial, sans-serif;
            height: 100vh;
            margin: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
            // justify-content: center;
        }
    </style>
    <script>
    let isPaused = false;  // Variable to check if the display is paused

    function togglePause() {
        isPaused = !isPaused;  // Toggle the paused state

        if (!isPaused) {
            showWord();  // If it's not paused anymore, continue displaying words
        }
    }

    function displayWords() {
        const inputText = document.getElementById('inputText').value;
        const words = inputText.split(/\s+/);
        const displayElement = document.getElementById('display');
        const speed = 60000 / document.getElementById('wpm').value; // converting wpm to milliseconds
        let index = 0;

        function showWord() {
            if (index < words.length && !isPaused) {
                displayElement.innerText = words[index];
                let delay = speed;
                if (words[index].endsWith('.') || words[index].endsWith('?') || words[index].endsWith(',') || words[index].endsWith('!') ||
                    /[\.,\?!][”"»]$/.test(words[index])) {
                    delay = 2 * speed;  // Double the delay for the specified conditions
                }
                index++;
                setTimeout(showWord, delay);
            } else if (isPaused) {
                setTimeout(showWord, 500);  // If it's paused, check again after 500ms
            }
        }
        
        showWord();
    }
    </script>
</head>
<body>
    <div>
        <textarea id="inputText" rows="4" cols="50"></textarea><br>
        Words per minute: <input type="number" id="wpm" value="350"><br>
        <button onclick="displayWords()">Submit</button>
        <button onclick="togglePause()">Pause</button>
    </div>
	<div id="display" style="font-size: 48px; margin-top: 20px; " ></div>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8004")
	http.ListenAndServe(":8004", nil)
}
