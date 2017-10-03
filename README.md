# Go Problem Sheet 1
My name is Cian Hatton, I'm currently in my 3rd year studying Software Development at [GMIT](http://www.gmit.ie/).

I created this repository to hold my solutions to this [problem sheet](https://data-representation.github.io/problems/go-fundamentals.html).

This is problem sheet 1 for the [Data Representation and Querying](https://data-representation.github.io/) module.

All solutions are written in the [Go](https://golang.org/) programming language.

In order to run the files, first clone the repository.

```bash
> git clone https://github.com/CHatton/Go-Problem-Sheet-1.git
```

Make sure that you have the [Go compiler](https://golang.org/dl/)  installed.

Navigate to the folder of the file you want to run.

```bash
> cd problem-sheets
```

You can run the program by either first building it.

```bash
> go build <file-name>
```

And then running the executable.

```bash
> ./<file-name>
```

Or you can simply use the command

```bash
> go run <file-name>
```

There are some test cases included in this repository. To run them
navigate to the root of the project.

Then use the command

```bash
> go test test\util_test.go
```

This will run all of the tests in the test suite. For more information about testing in Go, see [The Docs]{https://golang.org/pkg/testing/}.

You should get an output that looks like this.

```bash
=== RUN   TestReverseString
--- PASS: TestReverseString (0.00s)
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestAsIntSlice
--- PASS: TestAsIntSlice (0.00s)
PASS
ok      command-line-arguments  0.020s
```



