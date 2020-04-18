# How to design & develop REST microservices in Golang (Go)

### Best parts

### Fixing mistakes we made

- Video 17: 
> how to marshall structs (the yellow mark)
- Video 18: 
- Video 19: 

### Modules and Sum

- go.mod allows golang to manage dependencies easily
- go.sum assures that dependencies on which our app depends on, doesn't change unexpectedly. Each dependency is attached to a hash that defines the dependency version

### External packages

- router: *gin-gogic* which is faster than *httprouter*
- logger: *zap* developed by UBER `https://github.com/uber-go/zap`
  > Video 19: fixing mistakes - no logging system
  
