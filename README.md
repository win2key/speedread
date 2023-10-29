# speedread

This simple web application allows users to input text and then displays the text one word at a time at a specified words-per-minute (WPM) rate. The application uses basic HTML, CSS, and JavaScript on the frontend and Go on the backend.

### Features

1. **Word-by-Word Display**: Enter any text and view it one word at a time at your desired reading speed.
2. **Punctuation Aware**: The display delays are doubled for words ending with punctuation marks (e.g., `.`, `?`, `!`, `,`) or quotations like `”"»`.

### Prerequisites

Make sure you have Go installed on your system. You can download and install Go from [the official website](https://golang.org/).

### Setup and Run

1. **Clone the Repository**:
    ```bash
    git clone git@github.com:win2key/speedread.git
    cd speedread
    ```

2. **Run the Application**:
    ```bash
    go run main.go
    ```

3. **Access the Web App**: Open a browser and navigate to `http://localhost:8004`. 

### Usage

1. Enter the desired text in the textarea provided.
2. Specify the words-per-minute (WPM) rate at which you want the words to be displayed.
3. Click the "Submit" button.
4. The words will be displayed one-by-one in the large font below the textarea.

### Customization

You can easily customize the appearance and functionality by modifying the `htmlTemplate` constant in the `main.go` file.

### Contributing

Feel free to fork the repository and submit pull requests for any enhancements, bug fixes, or other contributions.

### License

This project is open-source. Check the LICENSE file for more details.

