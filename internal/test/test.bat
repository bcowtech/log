@ECHO OFF
FOR /F %%F IN ('DIR /B *_test.go') DO (
  go test -v %%F
)
