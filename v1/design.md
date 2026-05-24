## Error handling
### 1. Simple Failure
In this case, caller doesn't need to switch following logic. Just check if error occured.
```go
func ReturnSimple() (string, error){
    return "", fmt.Errorf("This is simple error just value")
}

simple, err := ReturnSimple()
if err!=nil{
    //fmt.Println(err) or
    //os.Exit()
}
```

### 2. Simple Custom Error
Define error value to check by errors.Is(). Export error struct as official or unify error in a pacakge.
```go
var ErrNotFound = errors.New("not found")
func readFile(path string)(File, error){
    //There is no file to read
    return File{}, ErrNotFound
}

func EditFile(path string){
    file, err := readFile("./text.txt")
    if errors.Is(err, ErrNotFound){
        //do something, like create a new file
    }
}
```

### 3. Complex Custome Error
Analyze error detail and handle it with field of error struct.
```go
type ConnectionError struct {
    WaitTime int
}

func (e *ConnectionError) Error() string {
    return fmt.Sprintf("connection failed, wait %d", e.WaitTime)
}

func GetOriginalResponse()(Response, error){
    //something fails
    return Response{}, fmt.Errorf("%w", &ConnectionError{WaitTime:100})
}

res, err := GetOriginalResponse()

var ce *ConnectionError
if errors.As(err, &ce) {
    if ce.WaitTime <= 50{
        //do something
    }else if ce.WaitTime <= 100 {
        //do something
    }else{
        //...
    }
}

```

### Reference
https://leapcell.io/blog/the-subtle-art-of-error-creation-understanding-errors-new-and-fmt-errorf-in-go#:~:text=for%20future%20refactors.-,General%20Guideline%3A,-Lean%20towards%20fmt