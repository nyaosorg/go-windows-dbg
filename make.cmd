call :"%1"
exit /b

:""
:"all"
    go fmt
    go build
    exit /b

:"test"
    tasklist | findstr /I dbgview || start dbgview
    go run example.go
    exit /b

:"test-ndebug"
    tasklist | findstr /I dbgview || start dbgview
    go run -tags=ndebug example.go
    exit /b
