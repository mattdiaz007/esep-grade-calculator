# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

Grades are averaged and weighted. The final numeric weighted score is converted to an A-F letter grade. 
Add grades per category with AddGrade(name, score, type).
Compute a final letter grade via GetFinalGrade().
Weighted category averaging.
