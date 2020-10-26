# gophergl
An abstraction over OpenGL and WebGL designed to allow you to write it once and run it both the browser and desktop


### OpenGL builidng

```
go run example.go
```

### Javascript building

```
gopherjs build example.go -o main.js
```

Then
```
firefox index.html
```