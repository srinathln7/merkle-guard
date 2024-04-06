# Package util 

The `util.go` file provides utility functions that are commonly used throughout the application. Let's break down each function:

1. **ServerLog(msg string)**:
   - Logs a message with the prefix "[grpc-server]" in blue color to distinguish server-related logs.
   - Parameters:
     - `msg`: The message to be logged.

2. **ClientLog(msg string)**:
   - Logs a message with the prefix "[grpc-client]" in yellow color to distinguish client-related logs.
   - Parameters:
     - `msg`: The message to be logged.

3. **ErrLog(msg string)**:
   - Logs an error message in red color to indicate an error.
   - Parameters:
     - `msg`: The error message to be logged.

4. **ReadFilesFromDir(dir string) ([][]byte, error)**:
   - Reads the contents of files from a specified directory.
   - Returns a slice of byte slices containing the content of each file and any encountered error.
   - Parameters:
     - `dir`: The directory path from which to read files.

5. **WriteFile(directory, fileName, content string) error**:
   - Writes content to a file in a specified directory.
   - Creates the directory if it doesn't exist and writes the content to the specified file.
   - Parameters:
     - `directory`: The directory path where the file will be written.
     - `fileName`: The name of the file to be written.
     - `content`: The content to be written to the file.
   - Returns any encountered error during file writing.

These utility functions encapsulate common operations such as logging, file reading, and file writing, providing a convenient and consistent way to perform these tasks throughout the application.