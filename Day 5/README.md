
# HTTP File Downloader with Real-Time Progress Bar

This project is a simple HTTP file downloader written in Go. It allows users to download files from any URL and displays a real-time progress bar. The progress bar shows the percentage of the file downloaded, the total file size, the download speed, and the estimated time remaining. The project demonstrates how Go can handle HTTP requests, file downloads, and display a dynamic terminal UI for real-time progress.

## How It Works

The program works as follows:

1. **Download the File**: The program makes an HTTP GET request to fetch the file from the specified URL.
2. **Track Progress**: The program tracks the download progress by wrapping the HTTP response body in a custom `ProgressReader`. The progress bar updates dynamically as more of the file is downloaded.
3. **Display Real-Time Progress**: The progress bar shows the percentage of the file downloaded, the total size in megabytes, the current download speed in MB/s, and the estimated time remaining (ETA).
4. **Handle Missing File Size**: If the server does not provide the `Content-Length` header, the program will still display the amount downloaded without the percentage and ETA.

## Project Structure

```
Day 5/
├── go.mod               # Go module file (already initialized for you)
├── main.go              # Main file containing logic to handle user input and start the download
├── download/            # Folder containing download logic
│   └── download.go      # File with the DownloadFile function
├── progress/            # Folder containing progress bar logic
│   └── progress.go      # File with the ProgressReader to display progress
```

## How to Run the Program

### 1. Clone the Repository

Clone this repository and navigate to the `Day 5` directory:

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 5
```

### 2. Run the Program

The project is already initialized with `go.mod`, so there's no need to initialize it again. Just run the following command to start the program:

```bash
go run main.go <url> <filename>
```

For example, to download a 1GB test file:

```bash
go run main.go http://ipv4.download.thinkbroadband.com/1GB.zip 1GB_test.zip
```

### Example Output

As the file downloads, you will see a dynamic progress bar like this:

```
[======================              ] 50.00% 500.00/1000.00 MB, 5.00 MB/s, ETA: 100 seconds
```

Once the download completes, you'll see:

```
Download completed successfully!
```

## File Breakdown

### `main.go`

- **`main()`**: The entry point of the program. It takes user input for the file URL and output filename, then calls the `download.DownloadFile()` function.

### `download.go` (in the `download` package)

- **`DownloadFile(url string, filename string)`**: This function handles the HTTP GET request to download the file. It uses a `ProgressReader` from the `progress` package to track and display download progress.

### `progress.go` (in the `progress` package)

- **`ProgressReader`**: A custom reader that wraps around the HTTP response body to track and display the progress of the download.
- **`displayProgress()`**: Displays the progress bar, including the percentage of the file downloaded, file size in MB, download speed in MB/s, and estimated time remaining (ETA). It handles cases where the file size is not provided by the server.

## Features

- **Real-Time Progress Bar**: Dynamically updates as the file downloads, showing the percentage, file size, download speed, and ETA.
- **Handles Missing File Size**: If the `Content-Length` header is missing, it shows only the downloaded file size without percentage and ETA.
- **Easy to Use**: Takes in the file URL and desired output filename as command-line arguments.
