
# Word Count Utility

This project is a command-line utility written in Go that counts the number of words or lines from standard input. This is similar to the `wc` command in Unix-like systems, providing functionality to count words by default and lines with an optional flag.

## Features

- **Word Counting**: Counts the total number of words by default.
- **Line Counting**: With the `-l` flag, counts the total number of lines instead of words.

## Installation

Ensure that Go (version 1.23.2 or later) is installed on your system.

1. Clone this repository:
   ```bash
   git clone https://github.com/abrishk26/wc.git
   cd wc
   ```

2. Build the project:
   ```bash
   go build -o wc
   ```

## Usage

Run the compiled program and pipe in text via standard input. Use the `-l` flag to count lines instead of words.

### Examples

1. **Counting Words**:
   ```bash
   echo "Hello World" | ./wc
   ```

2. **Counting Lines**:
   ```bash
   echo -e "Line 1\nLine 2" | ./wc -l
   ```

## Code Structure

- `main.go`: Contains the main functionality, including:
  - Parsing flags for word or line counting.
  - Counting words or lines using the `count` function, which scans input and increments a counter based on the mode specified.

## License

This project is open-source and available under the [MIT License](LICENSE).
