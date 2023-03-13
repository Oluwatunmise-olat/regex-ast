# regex-ast
Low budget "single line comment" stripper. 

Strip off `single line` comments from your code.

Supports files with extension:
 - `.ts`
 - `.js`
 - `.py`
 - `.go`
 
 Usage: 
  - `Step One:`Open your terminal in `pulled` project directory
  - `Step Two`: Run the command `go run main.go`. Note assumptions are made that you have go installed.
  - `Step Three`: Provide the `Absolute Path` to the file you want to strip comments off.
  - `Step Four`: View your file to check that the changes have been made accordingly.

Improvements:
- Switch to `goroutines` to handle multiple file stripping at once
- Use actual `ast` package rather than `regex`
