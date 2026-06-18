# ASCII Art Web

## Description
A web-based ASCII art generator that converts text into graphical
representations using ASCII characters. Supports three banner styles:
standard, shadow, and thinkertoy.

## Authors
- Agbaji Alexander

## Usage
1. Make sure the banner files are in the project root:
   standard.txt, shadow.txt, thinkertoy.txt

2. Run the server:
   go run main.go

3. Open your browser at:
   http://localhost:8080

4. Type your text, choose a banner, click Generate.

## Implementation Details
The server uses Go's net/http package to handle HTTP requests.
HTML templates are rendered using html/template.

### Algorithm
1. User submits text and banner choice via HTML form (POST)
2. Server reads the chosen banner file
3. Each character is looked up in the banner's 95-character table
4. Characters are printed row by row (8 rows per character)
5. Result is sent back to the browser inside an HTML template

### HTTP Endpoints
- GET  /           → serves the main page
- POST /ascii-art  → processes text and returns ASCII art

### HTTP Status Codes
- 200 OK                → successful generation
- 400 Bad Request       → empty text or invalid banner
- 404 Not Found         → template or banner file missing
- 500 Internal Server   → unexpected server error