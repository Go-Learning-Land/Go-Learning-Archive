# Golang Testing


### [Go Testing Fundamentals](https://github.com/quii/learn-go-with-tests) 

[Credits](https://github.com/quii/learn-go-with-tests)

1. [Install Go](https://github.com/quii/learn-go-with-tests/blob/main/install-go.md) - Set up environment for productivity.
2. [Hello, world](https://github.com/quii/learn-go-with-tests/blob/main/hello-world.md) - Declaring variables, constants, if/else statements, switch, write your first go program and write your first test. Sub-test syntax and closures.
3. [Integers](https://github.com/quii/learn-go-with-tests/blob/main/integers.md) - Further Explore function declaration syntax and learn new ways to improve the documentation of your code.
4. [Iteration](https://github.com/quii/learn-go-with-tests/blob/main/iteration.md) - Learn about `for` and benchmarking.
5. [Arrays and slices](https://github.com/quii/learn-go-with-tests/blob/main/arrays-and-slices.md) - Learn about arrays, slices, `len`, varargs, `range` and test coverage.
6. [Structs, methods & interfaces](https://github.com/quii/learn-go-with-tests/blob/main/structs-methods-and-interfaces.md) - Learn about `struct`, methods, `interface` and table driven tests.
7. [Pointers & errors](https://github.com/quii/learn-go-with-tests/blob/main/pointers-and-errors.md) - Learn about pointers and errors.
8. [Maps](https://github.com/quii/learn-go-with-tests/blob/main/maps.md) - Learn about storing values in the map data structure.
9. [Dependency Injection](https://github.com/quii/learn-go-with-tests/blob/main/dependency-injection.md) - Learn about dependency injection, how it relates to using interfaces and a primer on io.
10. [Mocking](https://github.com/quii/learn-go-with-tests/blob/main/mocking.md) - Take some existing untested code and use DI with mocking to test it.
11. [Concurrency](https://github.com/quii/learn-go-with-tests/blob/main/concurrency.md) - Learn how to write concurrent code to make your software faster.
12. [Select](https://github.com/quii/learn-go-with-tests/blob/main/select.md) - Learn how to synchronise asynchronous processes elegantly.
13. [Reflection](https://github.com/quii/learn-go-with-tests/blob/main/reflection.md) - Learn about reflection
14. [Sync](https://github.com/quii/learn-go-with-tests/blob/main/sync.md) - Learn some functionality from the sync package including `WaitGroup` and `Mutex`
15. [Context](https://github.com/quii/learn-go-with-tests/blob/main/context.md) - Use the context package to manage and cancel long-running processes
16. [Intro to property based tests](https://github.com/quii/learn-go-with-tests/blob/main/roman-numerals.md) - Practice some TDD with the Roman Numerals kata and get a brief intro to property based tests
17. [Maths](https://github.com/quii/learn-go-with-tests/blob/main/math.md) - Use the `math` package to draw an SVG clock
18. [Reading files](https://github.com/quii/learn-go-with-tests/blob/main/reading-files.md) - Read files and process them
19. [Templating](https://github.com/quii/learn-go-with-tests/blob/main/html-templates.md) - Use Go's html/template package to render html from data, and also learn about approval testing
20. [Generics](https://github.com/quii/learn-go-with-tests/blob/main/generics.md) - Learn how to write functions that take generic arguments and make your own generic data-structure
21. [Revisiting arrays and slices with generics](https://github.com/quii/learn-go-with-tests/blob/main/revisiting-arrays-and-slices-with-generics.md) - Generics are very useful when working with collections. Learn how to write your own `Reduce` function and tidy up some common patterns.


### Build an application with TDD

Now that you have hopefully digested the _Go Fundamentals_ section you have a solid grounding of a majority of Go's language features and how to do TDD.

This next section will involve building an application.

Each chapter will iterate on the previous one, expanding the application's functionality as our product owner dictates.

New concepts will be introduced to help facilitate writing great code but most of the new material will be learning what can be accomplished from Go's standard library.

By the end of this, you should have a strong grasp as to how to iteratively write an application in Go, backed by tests.

* [HTTP server](https://github.com/quii/learn-go-with-tests/blob/main/http-server.md) - We will create an application which listens to HTTP requests and responds to them.
* [JSON, routing and embedding](https://github.com/quii/learn-go-with-tests/blob/main/json.md) - We will make our endpoints return JSON and explore how to do routing.
* [IO and sorting](https://github.com/quii/learn-go-with-tests/blob/main/io.md) - We will persist and read our data from disk and we'll cover sorting data.
* [Command line & project structure](https://github.com/quii/learn-go-with-tests/blob/main/command-line.md) - Support multiple applications from one code base and read input from command line.
* [Time](https://github.com/quii/learn-go-with-tests/blob/main/time.md) - using the `time` package to schedule activities.
* [WebSockets](https://github.com/quii/learn-go-with-tests/blob/main/websockets.md) - learn how to write and test a server that uses WebSockets.