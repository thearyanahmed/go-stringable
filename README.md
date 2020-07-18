# go-stringable

*go-stringable* is a package that gives us an interface similar to laravel's Stringable class. 

**API** 

**initilization** 
```
import (
	"github.com/thearyanahmed/go-stringable/stringable"
)

myString := stringable.ToStringable("Hello world.")

myString.Get() // "Hello world." string
```

**Sluginfy**
```
myString.Slugify().Get() // "hello-world"
```

**ToLower**
```
myString.ToLower().Get() // "hello world"
```

**SnakeCase**
```
myString.SnakeCase().Get() // "hello_world"
```

**WordCount**
```
myString.WordCount().Get() // hello world -> 2 
```

**LetterCount**
```
myString.LetterCount().Get() // hello world -> 11 
```

**UCWords**
```
myString.UCWords().Get() // Hello World 
```

**After**
```
myString := stringable.ToStringable("The quick brown fox jumps over the lazy dog.")
myString.After("jumps").Get() // over the lazy dog.
```

**Before**
```
myString := stringable.ToStringable("The quick brown fox jumps over the lazy dog.")
myString.Before("jumps").Get() // The quick brown fox 
```

**Between**
```
myString := stringable.ToStringable("The quick brown fox jumps over the lazy dog.")
myString.Between("quick","over").Get() // brown fox jumps 
```

**ReplaceAll**
```
myString := stringable.ToStringable("one and one is two.")
myString.ReplaceAll("one","1").Get() // 1 and 1 is two
```

**ParseWith**
```
myString := stringable.ToStringable("The quick brown fox jumps over the lazy dog.")

parserMap := map[string]string{
	"quick": "fast",
	"over" : "above",
}

myString.ParseWith(parserMap).Get() // The fast brown fox jumps above the lazy dog.
```


**Slugify**
```
myString := stringable.ToStringable("one and one is two.")
myString.Slugify().Get() // one-and-one-is-two
```






















