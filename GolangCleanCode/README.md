# [Clean Go Code](https://github.com/Pungyeon/clean-go-article/blob/master/README.md)
[credits](https://github.com/Pungyeon/clean-go-article/blob/master/README.md)
[credits for style guide](https://github.com/uber-go/guide/blob/master/style.md)
## Preface: Why Write Clean Code?

This document is a reference for the Go community that aims to help developers write cleaner code. Whether you're working on a personal project or as part of a larger team, writing clean code is an important skill to have. Establishing good paradigms and consistent, accessible standards for writing clean code can help prevent developers from wasting many meaningless hours on trying to understand their own (or others') work.

> <em>We don’t read code, we <b>decode</b> it &ndash; Peter Seibel</em>

As developers, we're sometimes tempted to write code in a way that's convenient for the time being without regard for best practices; this makes code reviews and testing more difficult. In a sense, we're <em>encoding</em>&mdash;and, in doing so, making it more difficult for others to decode our work. But we want our code to be usable, readable, and maintainable. And that requires coding the <em>right</em> way, not the easy way.

This document begins with a simple and short introduction to the fundamentals of writing clean code. Later, we'll discuss concrete refactoring examples specific to Go.

##### A short word on `gofmt`
I'd like to take a few sentences to clarify my stance on `gofmt` because there are plenty of things I disagree with when it comes to this tool. I prefer snake case over camel case, and I quite like my constant variables to be uppercase. And, naturally, I also have many opinions on bracket placement. *That being said*, `gofmt` does allow us to have a common standard for writing Go code, and that's a great thing. As a developer myself, I can certainly appreciate that Go programmers may feel somewhat restricted by `gofmt`, especially if they disagree with some of its rules. But in my opinion, homogeneous code is more important than having complete expressive freedom.

## Table of Contents
- [Clean Go Code](#clean-go-code)
  - [Preface: Why Write Clean Code?](#preface-why-write-clean-code)
        - [A short word on `gofmt`](#a-short-word-on-gofmt)
  - [Table of Contents](#table-of-contents)
  - [Introduction to Clean Code](#introduction-to-clean-code)
    - [Test-Driven Development](#test-driven-development)
    - [Naming Conventions](#naming-conventions)
      - [Comments](#comments)
      - [Function Naming](#function-naming)
      - [Variable Naming](#variable-naming)
    - [Cleaning Functions](#cleaning-functions)
      - [Function Length](#function-length)
      - [Function Signatures](#function-signatures)
    - [Variable Scope](#variable-scope)
    - [Variable Declaration](#variable-declaration)
  - [Clean Go](#clean-go)
    - [Return Values](#return-values)
      - [Returning Defined Errors](#returning-defined-errors)
      - [Returning Dynamic Errors](#returning-dynamic-errors)
    - [Nil Values](#nil-values)
    - [Pointers in Go](#pointers-in-go)
      - [Pointer Mutability](#pointer-mutability)
    - [Closures Are Function Pointers](#closures-are-function-pointers)
    - [Interfaces in Go](#interfaces-in-go)
    - [The Empty `interface{}`](#the-empty-interface)
  - [Summary](#summary)
    
* [Clean Go](#Clean-Go)
    * [Return Values](#Return-Values) 
      * [Returning Defined Errors](#Returning-Defined-Errors)
      * [Returning Dynamic Errors](#Returning-Dynamic-Errors)
    * [Pointers in Go](#Pointers-in-Go)
    * [Closures Are Function Pointers](#Closures-are-Function-Pointers)
    * [Interfaces in Go](#Interfaces-in-Go)
    * [The Empty `interface{}`](#The-Empty-Interface)
* [Summary](#Summary)
* [Uber Style Guide](#Uber-Go-Style-Guide)
## Table of Contents Uber Style Guide
- [Introduction](#introduction)
- [Guidelines](#guidelines)
  - [Pointers to Interfaces](#pointers-to-interfaces)
  - [Verify Interface Compliance](#verify-interface-compliance)
  - [Receivers and Interfaces](#receivers-and-interfaces)
  - [Zero-value Mutexes are Valid](#zero-value-mutexes-are-valid)
  - [Copy Slices and Maps at Boundaries](#copy-slices-and-maps-at-boundaries)
  - [Defer to Clean Up](#defer-to-clean-up)
  - [Channel Size is One or None](#channel-size-is-one-or-none)
  - [Start Enums at One](#start-enums-at-one)
  - [Use `"time"` to handle time](#use-time-to-handle-time)
  - [Errors](#errors)
    - [Error Types](#error-types)
    - [Error Wrapping](#error-wrapping)
    - [Error Naming](#error-naming)
  - [Handle Type Assertion Failures](#handle-type-assertion-failures)
  - [Don't Panic](#dont-panic)
  - [Use go.uber.org/atomic](#use-gouberorgatomic)
  - [Avoid Mutable Globals](#avoid-mutable-globals)
  - [Avoid Embedding Types in Public Structs](#avoid-embedding-types-in-public-structs)
  - [Avoid Using Built-In Names](#avoid-using-built-in-names)
  - [Avoid `init()`](#avoid-init)
  - [Exit in Main](#exit-in-main)
    - [Exit Once](#exit-once)
  - [Use field tags in marshaled structs](#use-field-tags-in-marshaled-structs)
- [Performance](#performance)
  - [Prefer strconv over fmt](#prefer-strconv-over-fmt)
  - [Avoid string-to-byte conversion](#avoid-string-to-byte-conversion)
  - [Prefer Specifying Container Capacity](#prefer-specifying-container-capacity)
      - [Specifying Map Capacity Hints](#specifying-map-capacity-hints)
      - [Specifying Slice Capacity](#specifying-slice-capacity)
- [Style](#style)
  - [Avoid overly long lines](#avoid-overly-long-lines)
  - [Be Consistent](#be-consistent)
  - [Group Similar Declarations](#group-similar-declarations)
  - [Import Group Ordering](#import-group-ordering)
  - [Package Names](#package-names)
  - [Function Names](#function-names)
  - [Import Aliasing](#import-aliasing)
  - [Function Grouping and Ordering](#function-grouping-and-ordering)
  - [Reduce Nesting](#reduce-nesting)
  - [Unnecessary Else](#unnecessary-else)
  - [Top-level Variable Declarations](#top-level-variable-declarations)
  - [Prefix Unexported Globals with _](#prefix-unexported-globals-with-_)
  - [Embedding in Structs](#embedding-in-structs)
  - [Local Variable Declarations](#local-variable-declarations)
  - [nil is a valid slice](#nil-is-a-valid-slice)
  - [Reduce Scope of Variables](#reduce-scope-of-variables)
  - [Avoid Naked Parameters](#avoid-naked-parameters)
  - [Use Raw String Literals to Avoid Escaping](#use-raw-string-literals-to-avoid-escaping)
  - [Initializing Structs](#initializing-structs)
      - [Use Field Names to Initialize Structs](#use-field-names-to-initialize-structs)
      - [Omit Zero Value Fields in Structs](#omit-zero-value-fields-in-structs)
      - [Use `var` for Zero Value Structs](#use-var-for-zero-value-structs)
      - [Initializing Struct References](#initializing-struct-references)
  - [Initializing Maps](#initializing-maps)
  - [Format Strings outside Printf](#format-strings-outside-printf)
  - [Naming Printf-style Functions](#naming-printf-style-functions)
- [Patterns](#patterns)
  - [Test Tables](#test-tables)
  - [Functional Options](#functional-options)
- [Linting](#linting)

## Introduction to Clean Code

Clean code is the pragmatic concept of promoting readable and maintainable software. Clean code establishes trust in the codebase and helps minimize the chances of careless bugs being introduced. It also helps developers maintain their agility, which typically plummets as the codebase expands due to the increased risk of introducing bugs.

### Test-Driven Development

Test-driven development is the practice of testing your code frequently throughout short development cycles or sprints. It ultimately contributes to code cleanliness by inviting developers to question the functionality and purpose of their code. To make testing easier, developers are encouraged to write short functions that only do one thing. For example, it's arguably much easier to test (and understand) a function that's only 4 lines long than one that's 40.

Test-driven development consists of the following cycle:

1. Write (or execute) a test
2. If the test fails, make it pass
3. Refactor your code accordingly
4. Repeat

Testing and refactoring are intertwined in this process. As you refactor your code to make it more understandable or maintainable, you need to test your changes thoroughly to ensure that you haven't altered the behavior of your functions. This can be incredibly useful as the codebase grows.

###  Naming Conventions

#### Comments
I'd like to first address the topic of commenting code, which is an essential practice but tends to be misapplied. Unnecessary comments can indicate problems with the underlying code, such as the use of poor naming conventions. However, whether or not a particular comment is "necessary" is somewhat subjective and depends on how legibly the code was written. For example, the logic of well-written code may still be so complex that it requires a comment to clarify what is going on. In that case, one might argue that the comment is <em>helpful</em> and therefore necessary.

In Go, according to `gofmt`, <em>all</em> public variables and functions should be annotated. I think this is absolutely fine, as it gives us consistent rules for documenting our code. However, I always want to distinguish between comments that enable auto-generated documentation and <em>all other</em> comments. Annotation comments, for documentation, should be written like documentation&mdash;they should be at a high level of abstraction and concern the logical implementation of the code as little as possible.

I say this because there are other ways to explain code and ensure that it's being written comprehensibly and expressively. If the code is neither of those, some people find it acceptable to introduce a comment explaining the convoluted logic. Unfortunately, that doesn't really help. For one, most people simply won't read comments, as they tend to be very intrusive to the experience of reviewing code. Additionally, as you can imagine, a developer won't be too happy if they're forced to review unclear code that's been slathered with comments. The less that people have to read to understand what your code is doing, the better off they'll be.

Let's take a step back and look at some concrete examples. Here's how you <em>shouldn't</em> comment your code:

```go
// iterate over the range 0 to 9 
// and invoke the doSomething function
// for each iteration
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

This is what I like to call a <strong>tutorial comment</strong>; it's fairly common in tutorials, which often explain the low-level functionality of a language (or programming in general). While these comments may be helpful for beginners, they're absolutely useless in production code. Hopefully, we aren't collaborating with programmers who don't understand something as simple as a looping construct by the time they've begun working on a development team. As programmers, we shouldn't have to read the comment to understand what's going on&mdash;we know that we're iterating over the range 0 to 9 because we can simply read the code. Hence the proverb:

> <em>Document why, not how. &ndash; Venkat Subramaniam</em>

Following this logic, we can now change our comment to explain <em>why</em> we are iterating from the range 0 to 9:

```go
// instatiate 10 threads to handle upcoming work load
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

Now we understand <em>why</em> we have a loop and can tell <em>what</em> we're doing by simply reading the code... Sort of.

This still isn't what I'd consider clean code. The comment is worrying because it probably should not be necessary to express such an explanation in prose, assuming the code is well written (which it isn't). Technically, we're still saying what we're doing, not why we're doing it. We can easily express this "what" directly in our code by using more meaningful names:

```go
for workerID := 0; workerID < 10; workerID++ {
  instantiateThread(workerID)
}
```

With just a few changes to our variable and function names, we've managed to explain what we're doing directly in our code. This is much clearer for the reader because they won't have to read the comment and then map the prose to the code. Instead, they can simply read the code to understand what it's doing.

Of course, this was a relatively trivial example. Writing clear and expressive code is unfortunately not always so easy; it can become increasingly difficult as the codebase itself grows in complexity. The more you practice writing comments in this mindset and avoid explaining what you're doing, the cleaner your code will become.

#### Function Naming

Let's now move on to function naming conventions. The general rule here is really simple: the more specific the function, the more general its name. In other words, we want to start with a very broad and short function name, such as `Run` or `Parse`, that describes the general functionality. Let's imagine that we are creating a configuration parser. Following this naming convention, our top level of abstraction might look something like the following:

```go
func main() {
    configpath := flag.String("config-path", "", "configuration file path")
    flag.Parse()

    config, err := configuration.Parse(*configpath)
    
    ...
}
```

We'll focus on the naming of the `Parse` function. Despite this function's very short and general name, it's actually quite clear what it attempts to achieve.

When we go one layer deeper, our function naming will become slightly more specific:

```go
func Parse(filepath string) (Config, error) {
    switch fileExtension(filepath) {
    case "json":
        return parseJSON(filepath)
    case "yaml":
        return parseYAML(filepath)
    case "toml":
        return parseTOML(filepath)
    default:
        return Config{}, ErrUnknownFileExtension
    }
}
```

Here, we've clearly distinguished the nested function calls from their parent without being overly specific. This allows each nested function call to make sense on its own as well as within the context of the parent. On the other hand, if we had named the `parseJSON` function `json` instead, it couldn't possibly stand on its own. The functionality would become lost in the name, and we would no longer be able to tell whether this function is parsing, creating, or marshalling JSON. 

Notice that `fileExtension` is actually a little more specific. However, this is because its functionality is in fact quite specific in nature:

```go
func fileExtension(filepath string) string {
    segments := strings.Split(filepath, ".")
    return segments[len(segments)-1]
}
```

This kind of logical progression in our function names&mdash;from a high level of abstraction to a lower, more specific one&mdash;makes the code easier to follow and read. Consider the alternative: If our highest level of abstraction is too specific, then we'll end up with a name that attempts to cover all bases, like `DetermineFileExtensionAndParseConfigurationFile`. This is horrendously difficult to read; we are trying to be too specific too soon and end up confusing the reader, despite trying to be clear! 

#### Variable Naming
Rather interestingly, the opposite is true for variables. Unlike functions, our variables should be named from more to less specific the deeper we go into nested scopes.

> <em>You shouldn’t name your variables after their types for the same reason you wouldn’t name your pets 'dog' or 'cat'. &ndash; Dave Cheney</em>

Why should our variable names become less specific as we travel deeper into a function's scope? Simply put, as a variable's scope becomes smaller, it becomes increasingly clear for the reader what that variable represents, thereby eliminating the need for specific naming. In the example of the previous function `fileExtension`, we could even shorten the name of the variable `segments` to `s` if we wanted to. The context of the variable is so clear that it's unnecessary to explain it any further with longer variable names. Another good example of this is in nested for loops:

```go
func PrintBrandsInList(brands []BeerBrand) {
    for _, b := range brands { 
        fmt.Println(b)
    }
}
```

In the above example, the scope of the variable `b` is so small that we don't need to spend any additional brain power on remembering what exactly it represents. However, because the scope of `brands` is slightly larger, it helps for it to be more specific. When expanding the variable scope in the function below, this distinction becomes even more apparent:

```go
func BeerBrandListToBeerList(beerBrands []BeerBrand) []Beer {
    var beerList []Beer
    for _, brand := range beerBrands {
        for _, beer := range brand {
            beerList = append(beerList, beer)
        }
    }
    return beerList
}
```

Great! This function is easy to read. Now, let's apply the opposite (i.e., wrong) logic when naming our variables:

```go
func BeerBrandListToBeerList(b []BeerBrand) []Beer {
    var bl []Beer
    for _, beerBrand := range b {
        for _, beerBrandBeerName := range beerBrand {
            bl = append(bl, beerBrandBeerName)
        }
    }
    return bl
}
```

Even though it's possible to figure out what this function is doing, the excessive brevity of the variable names makes it difficult to follow the logic as we travel deeper. This could very well spiral into full-blown confusion because we're mixing short and long variable names inconsistently.

### Cleaning Functions

Now that we know some best practices for naming our variables and functions, as well as clarifying our code with comments, let's dive into some specifics of how we can refactor functions to make them cleaner.

#### Function Length

> <em>How small should a function be? Smaller than that! &ndash; Robert C. Martin</em>

When writing clean code, our primary goal is to make our code easily digestible. The most effective way to do this is to make our functions as short as possible. It's important to understand that we don't necessarily do this to avoid code duplication. The more important reason is to improve <em>code comprehension</em>.

It can help to look at a function's description at a very high level to understand this better:

```
fn GetItem:
    - parse json input for order id
    - get user from context
    - check user has appropriate role
    - get order from database
```

By writing short functions (which are typically 5&ndash;8 lines in Go), we can create code that reads almost as naturally as our description above:

```go
var (
    NullItem = Item{}
    ErrInsufficientPrivileges = errors.New("user does not have sufficient privileges")
)

func GetItem(ctx context.Context, json []bytes) (Item, error) {
    order, err := NewItemFromJSON(json)
    if err != nil {
        return NullItem, err
    }
    if !GetUserFromContext(ctx).IsAdmin() {
	      return NullItem, ErrInsufficientPrivileges
    }
    return db.GetItem(order.ItemID)
}
```

Using smaller functions also eliminates another horrible habit of writing code: indentation hell. <strong>Indentation hell</strong> typically occurs when a chain of `if` statements are carelessly nested in a function. This makes it <em>very</em> difficult for human beings to parse the code and should be eliminated whenever spotted. Indentation hell is particularly common when working with `interface{}` and using type casting:

```go
func GetItem(extension string) (Item, error) {
    if refIface, ok := db.ReferenceCache.Get(extension); ok {
        if ref, ok := refIface.(string); ok {
            if itemIface, ok := db.ItemCache.Get(ref); ok {
                if item, ok := itemIface.(Item); ok {
                    if item.Active {
                        return Item, nil
                    } else {
                      return EmptyItem, errors.New("no active item found in cache")
                    }
                } else {
                  return EmptyItem, errors.New("could not cast cache interface to Item")
                }
            } else {
              return EmptyItem, errors.New("extension was not found in cache reference")
            }
        } else {
          return EmptyItem, errors.New("could not cast cache reference interface to Item")
        }
    }
    return EmptyItem, errors.New("reference not found in cache")
}
```

First, indentation hell makes it difficult for other developers to understand the flow of your code. Second, if the logic in our `if` statements expands, it'll become exponentially more difficult to figure out which statement returns what (and to ensure that all paths return some value). Yet another problem is that this deep nesting of conditional statements forces the reader to frequently scroll and keep track of many logical states in their head. It also makes it more difficult to test the code and catch bugs because there are so many different nested possibilities that you have to account for.

Indentation hell can result in reader fatigue if a developer has to constantly parse unwieldy code like the sample above. Naturally, this is something we want to avoid at all costs.

So, how do we clean this function? Fortunately, it's actually quite simple. On our first iteration, we will try to ensure that we are returning an error as soon as possible. Instead of nesting the `if` and `else` statements, we want to "push our code to the left," so to speak. Take a look:

```go
func GetItem(extension string) (Item, error) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, errors.New("reference not found in cache")
    }

    ref, ok := refIface.(string)
    if !ok {
        // return cast error on reference 
    }

    itemIface, ok := db.ItemCache.Get(ref)
    if !ok {
        // return no item found in cache by reference
    }

    item, ok := itemIface.(Item)
    if !ok {
        // return cast error on item interface
    }

    if !item.Active {
        // return no item active
    }

    return Item, nil
}
```

Once we're done with our first attempt at refactoring the function, we can proceed to split up the function into smaller functions. Here's a good rule of thumb:  If the `value, err :=` pattern is repeated more than once in a function, this is an indication that we can split the logic of our code into smaller pieces:

```go
func GetItem(extension string) (Item, error) {
    ref, ok := getReference(extension)
    if !ok {
        return EmptyItem, ErrReferenceNotFound
    }
    return getItemByReference(ref)
}

func getReference(extension string) (string, bool) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, false
    }
    return refIface.(string)
}

func getItemByReference(reference string) (Item, error) {
    item, ok := getItemFromCache(reference)
    if !item.Active || !ok {
        return EmptyItem, ErrItemNotFound
    }
    return Item, nil
}

func getItemFromCache(reference string) (Item, bool) {
    if itemIface, ok := db.ItemCache.Get(ref); ok {
        return EmptyItem, false
    }
    return itemIface.(Item), true
}
```

As mentioned previously, indentation hell can make it difficult to test our code. When we split up our `GetItem` function into several helpers, we make it easier to track down bugs when testing our code. Unlike the original version, which consisted of several `if` statements in the same scope, the refactored version of `GetItem` has just two branching paths that we must consider. The helper functions are also short and digestible, making them easier to read.

> Note: For production code, one should elaborate on the code even further by returning errors instead of `bool` values. This makes it much easier to understand where the error is originating from. However, as these are just example functions, returning `bool` values will suffice for now. Examples of returning errors more explicitly will be explained in more detail later.

Notice that cleaning the `GetItem` function resulted in more lines of code overall. However, the code itself is now much easier to read. It's layered in an onion-style fashion, where we can ignore "layers" that we aren't interested in and simply peel back the ones that we do want to examine. This makes it easier to understand low-level functionality because we only have to read maybe 3&ndash;5 lines at a time.

This example illustrates that we cannot measure the cleanliness of our code by the number of lines it uses. The first version of the code was certainly much shorter. However, it was <em>artificially</em> short and very difficult to read. In most cases, cleaning code will initially expand the existing codebase in terms of the number of lines. But this is highly preferable to the alternative of having messy, convoluted logic. If you're ever in doubt about this, just consider how you feel about the following function, which does exactly the same thing as our code but only uses two lines:

```go
func GetItemIfActive(extension string) (Item, error) {
    if refIface,ok := db.ReferenceCache.Get(extension); ok {if ref,ok := refIface.(string); ok { if itemIface,ok := db.ItemCache.Get(ref); ok { if item,ok := itemIface.(Item); ok { if item.Active { return Item,nil }}}}} return EmptyItem, errors.New("reference not found in cache")
}
```

#### Function Signatures

Creating a good function naming structure makes it easier to read and understand the intent of the code. As we saw above, making our functions shorter helps us understand the function's logic. The last part of cleaning our functions involves understanding the context of the function input. With this comes another easy-to-follow rule: <strong>Function signatures should only contain one or two input parameters</strong>. In certain exceptional cases, three can be acceptable, but this is where we should start considering a refactor. Much like the rule that our functions should only be 5&ndash;8 lines long, this can seem quite extreme at first. However, I feel that this rule is much easier to justify.

Take the following function from [RabbitMQ's introduction tutorial to its Go library](https://www.rabbitmq.com/tutorials/tutorial-one-go.html):

```go
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
```

The function `QueueDeclare` takes six input parameters, which is quite a lot. With some effort, it's possible to understand what this code does thanks to the comments. However, the comments are actually part of the problem&mdash;as mentioned earlier, they should be substituted with descriptive code whenever possible. After all, there's nothing preventing us from invoking the `QueueDeclare` function <em>without</em> comments:

```go
q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
```

Now, without looking at the commented version, try to remember what the fourth and fifth `false` arguments represent. It's impossible, right? You will inevitably forget at some point. This can lead to costly mistakes and bugs that are difficult to correct. The mistakes might even occur through incorrect comments&mdash;imagine labeling the wrong input parameter. Correcting this mistake will be unbearably difficult to correct, especially when familiarity with the code has deteriorated over time or was low to begin with. Therefore, it is recommended to replace these input parameters with an 'Options' `struct` instead:

```go
type QueueOptions struct {
    Name string
    Durable bool
    DeleteOnExit bool
    Exclusive bool
    NoWait bool
    Arguments []interface{} 
}

q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
    Durable: false,
    DeleteOnExit: false,
    Exclusive: false,
    NoWait: false,
    Arguments: nil,
})
```

This solves two problems: misusing comments, and accidentally labeling the variables incorrectly. Of course, we can still confuse properties with the wrong value, but in these cases, it will be much easier to determine where our mistake lies within the code. The ordering of the properties also doesn't matter anymore, so incorrectly ordering the input values is no longer a concern. The last added bonus of this technique is that we can use our `QueueOptions` struct to infer the default values of our function's input parameters. When structures in Go are declared, all properties are initialised to their default value. This means that our `QueueDeclare` option can actually be invoked in the following way:

```go
q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
})
```

The rest of the values are initialised to their default value of `false` (except for `Arguments`, which as an interface has a default value of `nil`). Not only are we much safer with this approach, but we are also much clearer with our intentions. In this case, we could actually write less code. This is an all-around win for everyone on the project.

One final note on this: It's not always possible to change a function's signature. In this case, for example, we don't actually have control over our `QueueDeclare` function signature because it's from the RabbitMQ library. It's not our code, so we can't change it. However, we can wrap these functions to suit our purposes:

```go
type RMQChannel struct {
    channel *amqp.Channel
}

func (rmqch *RMQChannel) QueueDeclare(opts QueueOptions) (Queue, error) {
    return rmqch.channel.QueueDeclare(
        opts.Name,
        opts.Durable,
        opts.DeleteOnExit,
        opts.Exclusive,
        opts.NoWait,
        opts.Arguments, 
    )
} 
```

Basically, we create a new structure named `RMQChannel` that contains the `amqp.Channel` type, which has the `QueueDeclare` method. We then create our own version of this method, which essentially just calls the old version of the RabbitMQ library function. Our new method has all the advantages described before, and we achieved this without actually having to change any of the code in the RabbitMQ library.

We'll use this idea of wrapping functions to introduce more clean and safe code later when discussing `interface{}`.

### Variable Scope
Now, let's take a step back and revisit the idea of writing smaller functions. This has another nice side effect that we didn't cover in the previous chapter: Writing smaller functions can typically eliminate reliance on mutable variables that leak into the global scope.

Global variables are problematic and don't belong in clean code; they make it very difficult for programmers to understand the current state of a variable. If a variable is global and mutable, then by definition, its value can be changed by any part of the codebase. At no point can you guarantee that this variable is going to be a specific value... And that's a headache for everyone. This is yet another example of a trivial problem that's exacerbated when the codebase expands.

Let's look at a short example of how non-global variables with a large scope can cause problems. These variables also introduce the issue of <strong>variable shadowing</strong>, as demonstrated in the code taken from an article titled [Golang scope issue](https://idiallo.com/blog/golang-scopes):

```go
func doComplex() (string, error) {
    return "Success", nil
}

func main() {
    var val string
    num := 32

    switch num {
    case 16:
    // do nothing
    case 32:
        val, err := doComplex()
        if err != nil {
            panic(err)
        }
        if val == "" {
            // do something else
        }
    case 64:
        // do nothing
    }

    fmt.Println(val)
}
```

What's the problem with this code? From a quick skim, it seems the `var val string` value should be printed out as `Success` by the end of the `main` function. Unfortunately, this is not the case. The reason for this lies in the following line:

```go
val, err := doComplex()
```

This declares a new variable `val` in the switch's `case 32` scope and has nothing to do with the variable declared in the first line of `main`. Of course, it can be argued that Go syntax is a little tricky, which I don't necessarily disagree with, but there is a much worse issue at hand. The declaration of `var val string` as a mutable, largely scoped variable is completely unnecessary. If we do a <strong>very</strong> simple refactor, we will no longer have this issue:

```go
func getStringResult(num int) (string, error) {
    switch num {
    case 16:
    // do nothing
    case 32:
       return doComplex()
    case 64:
        // do nothing
    }
    return "", nil
}

func main() {
    val, err := getStringResult(32)
    if err != nil {
        panic(err)
    }
    if val == "" {
        // do something else
    }
    fmt.Println(val)
}
```

After our refactor, `val` is no longer modified, and the scope has been reduced. Again, keep in mind that these functions are very simple. Once this kind of code style becomes a part of larger, more complex systems, it can be impossible to figure out why errors are occurring. We don't want this to happen&mdash;not only because we generally dislike software errors but also because it's disrespectful to our colleagues, and ourselves; we are potentially wasting each other's time having to debug this type of code. Developers need to take responsibility for their own code rather than blaming these issues on the variable declaration syntax of a particular language like Go.

On a side note, if the `// do something else` part is another attempt to mutate the `val` variable, we should extract that logic out as its own self-contained function, as well as the previous part of it. This way, instead of expanding the mutable scope of our variables, we can just return a new value:

```go
func getVal(num int) (string, error) {
    val, err := getStringResult(num)
    if err != nil {
        return "", err
    }
    if val == "" {
        return NewValue() // pretend function
    }
}

func main() {
    val, err := getVal(32)
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}
```

### Variable Declaration 

Other than avoiding issues with variable scope and mutability, we can also improve readability by declaring variables as close to their usage as possible. In C programming, it's common to see the following approach to declaring variables:

```go
func main() {
  var err error
  var items []Item
  var sender, receiver chan Item
  
  items = store.GetItems()
  sender = make(chan Item)
  receiver = make(chan Item)
  
  for _, item := range items {
    ...
  }
}
```

This suffers from the same symptom as described in our discussion of variable scope. Even though these variables might not actually be reassigned at any point, this kind of coding style keeps the readers on their toes, in all the wrong ways. Much like computer memory, our brain's short-term memory has a limited capacity. Having to keep track of which variables are mutable and whether or not a particular fragment of code will mutate them makes it more difficult to understand what the code is doing. Figuring out the eventually returned value can be a nightmare. Therefore, to makes this easier for our readers (and our future selves), it's recommended that you declare variables as close to their usage as possible:

```go
func main() {
	var sender chan Item
	sender = make(chan Item)

	go func() {
		for {
			select {
			case item := <-sender:
				// do something
			}
		}
	}()
}
```

However, we can do even better by invoking the function directly after its declaration. This makes it much clearer that the function logic is associated with the declared variable:

```go
func main() {
  sender := func() chan Item {
    channel := make(chan Item)
    go func() {
      for {
        select { ... }
      }
    }()
    return channel
  }
}
```

And coming full circle, we can move the anonymous function to make it a named function instead:

```go
func main() {
  sender := NewSenderChannel()
}

func NewSenderChannel() chan Item {
  channel := make(chan Item)
  go func() {
    for {
      select { ... }
    }
  }()
  return channel
}
```

It is still clear that we are declaring a variable, and the logic associated with the returned channel is simple, unlike in the first example. This makes it easier to traverse the code and understand the role of each variable.

Of course, this doesn't actually prevent us from mutating our `sender` variable. There is nothing that we can do about this, as there is no way of declaring a `const struct` or `static` variables in Go. This means that we'll have to restrain ourselves from modifying this variable at a later point in the code.

> NOTE: The keyword `const` does exist but is limited in use to primitive types only.

One way of getting around this can at least limit the mutability of a variable to the package level. The trick involves creating a structure with the variable as a private property. This private property is thenceforth only accessible through other methods provided by this wrapping structure. Expanding on our channel example, this would look something like the following:

```go
type Sender struct {
  sender chan Item
}

func NewSender() *Sender {
  return &Sender{
    sender: NewSenderChannel(),
  }
}

func (s *Sender) Send(item Item) {
  s.sender <- item
}
```

We have now ensured that the `sender` property of our `Sender` struct is never mutated&mdash;at least not from outside of the package. As of writing this document, this is the only way of creating publicly immutable non-primitive variables. It's a little verbose, but it's truly worth the effort to ensure that we don't end up with strange bugs resulting from accidental variable modification. 

```go
func main() {
  sender := NewSender()
  sender.Send(&Item{})
}
```

Looking at the example above, it's clear how this also simplifies the usage of our package. This way of hiding the implementation is beneficial not only for the maintainers of the package but also for the users. Now, when initialising and using the `Sender` structure, there is no concern over its implementation. This opens up for a much looser architecture. Because our users aren't concerned with the implementation, we are free to change it at any point, since we have reduced the point of contact that users have with the package. If we no longer wish to use a channel implementation in our package, we can easily change this without breaking the usage of the `Send` method (as long as we adhere to its current function signature).

>  NOTE: There is a fantastic explanation of how to handle the abstraction in client libraries, taken from the talk [AWS re:Invent 2017: Embracing Change without Breaking the World (DEV319)](https://www.youtube.com/watch?v=kJq81Y7OEx4).

## Clean Go

This section focuses less on the generic aspects of writing clean Go code and more on the specifics, with an emphasis on the underlying clean code principles.

### Return Values

#### Returning Defined Errors

We'll start things off nice and easy by describing a cleaner way to return errors. As we discussed earlier, our main goal with writing clean code is to ensure readability, testability, and maintainability of the codebase. The technique for returning errors that we'll discuss here will achieve all three of those goals with very little effort.

Let's consider the normal way to return a custom error. This is a hypothetical example taken from a thread-safe map implementation that we've named `Store`:

```go
package smelly

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return Item{}, errors.New("item could not be found in the store") 
    }
    return item, nil
}
```

There is nothing inherently smelly about this function when we consider it in isolation. We look into the `items` map of our `Store` struct to see if we already have an item with the given `id`. If we do, we return it; otherwise, we return an error. Pretty standard. So, what is the issue with returning custom errors as string values? Well, let's look at what happens when we use this function inside another package:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := smelly.GetItem("123")
    if err != nil {
        if err.Error() == "item could not be found in the store" {
            http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, errr.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

This is actually not too bad. However, there is one glaring problem: An error in Go is simply an `interface` that implements a function (`Error()`) returning a string; thus, we are now hardcoding the expected error code into our codebase, which isn't ideal. This hardcoded string is known as a <strong>magic string</strong>. And its main problem is flexibility: If at some point we decide to change the string value used to represent an error, our code will break (softly) unless we update it in possibly many different places. Our code is tightly coupled&mdash;it relies on that specific magic string and the assumption that it will never change as the codebase grows.

An even worse situation would arise if a client were to use our package in their own code. Imagine that we decided to update our package and changed the string that represents an error&mdash;the client's software would now suddenly break. This is quite obviously something that we want to avoid. Fortunately, the fix is very simple:

```go
package clean

var (
    NullItem = Item{}

    ErrItemNotFound = errors.New("item could not be found in the store") 
)

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, ErrItemNotFound
    }
    return item, nil
}
```

By simply representing the error as a variable (`ErrItemNotFound`), we've ensured that anyone using this package can check against the variable rather than the actual string that it returns:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err, clean.ErrItemNotFound) {
           http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

This feels much nicer and is also much safer. Some would even say that it's easier to read as well. In the case of a more verbose error message, it certainly would be preferable for a developer to simply read `ErrItemNotFound` rather than a novel on why a certain error has been returned.

This approach is not limited to errors and can be used for other returned values. As an example, we are also returning a `NullItem` instead of `Item{}` as we did before. There are many different scenarios in which it might be preferable to return a defined object, rather than initialising it on return.

Returning default `NullItem` values like we did in the previous examples can also be safer in certain cases. As an example, a user of our package could forget to check for errors and end up initialising a variable that points to an empty struct containing a default value of `nil` as one or more property values. When attempting to access this `nil` value later in the code, the client software would panic. However, when we return our custom default value instead, we can ensure that all values that would otherwise default to `nil` are initialised. Thus, we'd ensure that we do not cause panics in our users' software.

This also benefits us. Consider this: If we wanted to achieve the same safety without returning a default value, we would have to change our code everywhere we return this type of empty value. However, with our default value approach, we now only have to change our code in a single place:

```go
var NullItem = Item{
    itemMap: map[string]Item{},
}
```
> NOTE: In many scenarios, invoking a panic will actually be preferable to indicate that there is an error check missing.

> NOTE: Every interface property in Go has a default value of `nil`. This means that this is useful for any struct that has an interface property. This is also true for structs that contain channels, maps, and slices, which could potentially also have a `nil` value.

#### Returning Dynamic Errors
There are certainly some scenarios where returning an error variable might not actually be viable. In cases where the information in customised errors is dynamic, if we want to describe error events more specifically, we can no longer define and return our static errors. Here's an example:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, fmt.Errorf("Could not find item with ID: %s", id)
    }
    return item, nil
}
```

So, what to do? There is no well-defined or standard method for handling and returning these kinds of dynamic errors. My personal preference is to return a new interface, with a bit of added functionality:

```go
type ErrorDetails interface {
    Error() string
    Type() string
}

type errDetails struct {
    errtype error
    details interface{}
}

func NewErrorDetails(err error, details ...interface{}) ErrorDetails {
    return &errDetails{
        errtype: err,
        details: details,
    }
}

func (err *errDetails) Error() string {
    return fmt.Sprintf("%v: %v", err.errtype, err.details)
}

func (err *errDetails) Type() error {
    return err.errtype
}
```

This new data structure still works as our standard error. We can still compare it to `nil` since it's an interface implementation, and we can still call `.Error()` on it, so it won't break any existing implementations. However, the advantage is that we can now check our error type as we could previously, despite our error now containing the <em>dynamic</em> details:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, NewErrorDetails(
            ErrItemNotFound,
            fmt.Sprintf("could not find item with id: %s", id))
    }
    return item, nil
}
```

And our HTTP handler function can then be refactored to check for a specific error again:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err.Type(), clean.ErrItemNotFound) {
            http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

### Nil Values 

A controversial aspect of Go is the addition of `nil`. This value corresponds to the value `NULL` in C and is essentially an uninitialised pointer. We've already seen some of the problems that `nil` can cause, but to sum up: Things break when you try to access methods or properties of a `nil` value. Thus, it's recommended to avoid returning a  `nil` value when possible. This way, the users of our code are less likely to accidentally access `nil` values. 

There are other scenarios in which it is common to find `nil` values that can cause some unnecessary pain. An example of this is incorrectly initialising a `struct` (as in the example below), which can lead to it containing `nil` properties. If accessed, those `nil`s will cause a panic.

```go
type App struct {
	Cache *KVCache
}

type KVCache struct {
  mtx sync.RWMutex
	store map[string]string
}

func (cache *KVCache) Add(key, value string) {
  cache.mtx.Lock()
  defer cache.mtx.Unlock()
  
	cache.store[key] = value
}
```

This code is absolutely fine. However, the danger is that our `App` can be initialised incorrectly, without initialising the `Cache` property within. Should the following code be invoked, our application will panic:

```go
	app := App{}
	app.Cache.Add("panic", "now")
```

The `Cache` property has never been initialised and is therefore a `nil` pointer. Thus, invoking the `Add` method like we did here will cause a panic, with the following message:

> panic: runtime error: invalid memory address or nil pointer dereference

Instead, we can turn the `Cache` property of our `App` structure into a private property and create a getter-like method to access it. This gives us more control over what we are returning; specifically, it ensures that we aren't returning a `nil` value:

```go
type App struct {
    cache *KVCache
}

func (app *App) Cache() *KVCache {
	if app.cache == nil {
        app.cache = NewKVCache()
	}
	return app.cache
}
```

The code that previously panicked will now be refactored to the following:

```go
app := App{}
app.Cache().Add("panic", "now")
```

This ensures that users of our package don't have to worry about the implementation and whether they're using our package in an unsafe manner. All they need to worry about is writing their own clean code.

> NOTE: There are other methods of achieving a similarly safe outcome. However, I believe this is the most straightforward approach.

### Pointers in Go
Pointers in Go are a rather extensive topic. They're a very big part of working with the language&mdash;so much so that it is essentially impossible to write Go without some knowledge of pointers and their workings in the language. Therefore, it is important to understand how to use pointers without adding unnecessary complexity (and thereby keeping your codebase clean). Note that we will not review the details of how pointers are implemented in Go. Instead, we will focus on the quirks of Go pointers and how we can handle them.

Pointers add complexity to code. If we aren't cautious, incorrectly using pointers can introduce nasty side effects or bugs that are particularly difficult to debug. By sticking to the basic principles of writing clean code that we covered in the first part of this document, we can at least reduce the chances of introducing unnecessary complexity to our code.

#### Pointer Mutability
We've already looked at the problem of mutability in the context of globally or largely scoped variables. However, mutability is not necessarily always a bad thing, and I am by no means an advocate for writing 100% pure functional programs. Mutability is a powerful tool, but we should really only ever use it when it's necessary. Let's have a look at a code example illustrating why:

```go
func (store *UserStore) Insert(user *User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = user
    return nil
}

func (store *UserStore) userExists(id int64) bool {
    _, ok := store.users[id]
    return ok
}
```

At first glance, this doesn't seem too bad. In fact, it might even seem like a rather simple insert function for a common list structure. We accept a pointer as input, and if no other users with this `id` exist, then we insert the provided user pointer into our list. Then, we use this functionality in our public API for creating new users:

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {
    user, err := parseUserFromRequest(r)
    if err != nil {
        http.Error(w, err, http.StatusBadRequest)
        return
    }
    if err := insertUser(w, user); err != nil {
      http.Error(w, err, http.StatusInternalServerError)
      return
    }
}

func insertUser(w http.ResponseWriter, user User) error {
  	if err := store.Insert(user); err != nil {
        return err
    }
  	user.Password = ""
	  return json.NewEncoder(w).Encode(user)
}
```

Once again, at first glance, everything looks fine. We parse the user from the received request and insert the user struct into our store. Once we have successfully inserted our user into the store, we then set the password to be an empty string before returning the user as a JSON object to our client. This is all quite common practice, typically when returning a user object whose password has been hashed, since we don't want to return the hashed password.

However, imagine that we are using an in-memory store based on a `map`. This code will produce some unexpected results. If we check our user store, we'll see that the change we made to the users password in the HTTP handler function also affected the object in our store. This is because the pointer address returned by `parseUserFromRequest` is what we populated our store with, rather than an actual value. Therefore, when making changes to the dereferenced password value, we end up changing the value of the object we are pointing to in our store.

This is a great example of why both mutability and variable scope can cause some serious issues and bugs when used incorrectly. When passing pointers as an input parameter of a function, we are expanding the scope of the variable whose data is being pointed to. Even more worrying is the fact that we are expanding the scope to an undefined level. We are *almost* expanding the scope of the variable to the global level. As demonstrated by the above example, this can lead to disastrous bugs that are particularly difficult to find and eradicate.

Fortunately, the fix for this is rather simple:

```go
func (store *UserStore) Insert(user User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = &user
    return nil
}
```

Instead of passing a pointer to a `User` struct, we are now passing in a copy of a `User`. We are still storing a pointer to our store; however, instead of storing the pointer from outside of the function, we are storing the pointer to the copied value, whose scope is inside the function. This fixes the immediate problem but might still cause issues further down the line if we aren't careful. Consider this code:

```go
func (store *UserStore) Get(id int64) (*User, error) {
    user, ok := store.users[id]
    if !ok {
        return EmptyUser, ErrUserNotFound
    }
    return store.users[id], nil
}
```

Again, this is a very standard implementation of a getter function for our store. However, it's still bad code because we are once again expanding the scope of our pointer, which may end up causing unexpected side effects. When returning the actual pointer value, which we are storing in our user store, we are essentially giving other parts of our application the ability to change our store values. This is bound to cause confusion. Our store should be the only entity allowed to make changes to its values. The easiest fix for this is to return a value of `User` rather than returning a pointer.

> NOTE: Consider the case where our application uses multiple threads. In this scenario, passing pointers to the same memory location can also potentially result in a race condition. In other words, we aren't only potentially corrupting our data&mdash;we could also cause a panic from a data race.

Please keep in mind that there is intrinsically nothing wrong with returning pointers. However, the expanded scope of variables (and the number of owners pointing to those variables) is the most important consideration when working with pointers. This is what categorises our previous example as a smelly operation. This is also why common Go constructors are absolutely fine:

```go
func AddName(user *User, name string) {
    user.Name = name
}
```

This is *okay* because the variable scope, which is defined by whoever invokes the function, remains the same after the function returns. Combined with the fact that the function invoker remains the sole owner of the variable, this means that the pointer cannot be manipulated in an unexpected manner.

### Closures Are Function Pointers

Before we get into the next topic of using interfaces in Go, I would like to introduce a common alternative. It's what C programmers know as "function pointers" and what most other programming languages call <strong>closures</strong>. A closure is simply an input parameter like any other, except it represents (points to) a function that can be invoked. In JavaScript, it's quite common to use closures as callbacks, which are just functions that are invoked after some asynchronous operation has finished. In Go, we don't really have this notion. We can, however, use closures to partially overcome a different hurdle: The lack of generics.

Consider the following function signature:

```go
func something(closure func(float64) float64) float64 { ... }
```

Here, `something` takes another function (a closure) as input and returns a `float64`. The input function takes a `float64` as input and also returns a `float64`. This pattern can be particularly useful for creating a loosely coupled architecture, making it easier to to add functionality without affecting other parts of the code. Suppose we have a struct containing data that we want to manipulate in some form. Through this structure's `Do()` method, we can perform operations on that data. If we know the operation ahead of time, we can obviously handle that logic directly in our `Do()` method:

```go
func (datastore *Datastore) Do(operation Operation, data []byte) error {
  switch(operation) {
  case COMPARE:
    return datastore.compare(data)
  case CONCAT:
    return datastore.add(data)
  default:
    return ErrUnknownOperation
  }
}
```

But as you can imagine, this function is quite rigid&mdash;it performs a predetermined operation on the data contained in the `Datastore` struct. If at some point we would like to introduce more operations, we'd end up bloating our `Do` method with quite a lot of irrelevant logic that would be hard to maintain. The function would have to always care about what operation it's performing and to cycle through a number of nested options for each operation. It might also be an issue for developers wanting to use our `Datastore` object who don't have access to edit our package code, since there is no way of extending structure methods in Go as there is in most OOP languages.

So instead, let's try a different approach using closures:

```go
func (datastore *Datastore) Do(operation func(data []byte, data []byte) ([]byte, error), data []byte) error {
  result, err := operation(datastore.data, data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func concat(a []byte, b []byte) ([]byte, error) {
  ...
}

func main() {
  ...
  datastore.Do(concat, data)
  ...
}
```

You'll notice immediately that the function signature for `Do` ends up being quite messy. We also have another issue: The closure isn't particularly generic. What happens if we find out that we actually want the `concat` to be able to take more than just two byte arrays as input? Or if we want to add some completely new functionality that may also need more or fewer input values than `(data []byte, data []byte)`?

One way to solve this issue is to change our `concat` function. In the example below, I have changed it to only take a single byte array as an input argument, but it could just as well have been the opposite case:

```go
func concat(data []byte) func(data []byte) ([]byte, error) {
  return func(concatting []byte) ([]byte, error) {
    return append(data, concatting), nil
  }
}

func (datastore *Datastore) Do(operation func(data []byte) ([]byte, error)) error {
  result, err := operation(datastore.data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func main() {
  ...
  datastore.Do(compare(data))
  ...
}
```

Notice how we've effectively moved some of the clutter out of the `Do` method signature and into the `concat` method signature. Here, the `concat` function returns yet another function. Within the returned function, we store the input values originally passed in to our `concat` function. The returned function can therefore now take a single input parameter; within our function logic, we will append it to our original input value. As a newly introduced concept, this may seem quite strange. However, it's good to get used to having this as an option; it can help loosen up logic coupling and get rid of bloated functions.

In the next section, we'll get into interfaces. Before we do so, let's take a short moment to discuss the difference between interfaces and closures. First, it's worth noting that interfaces and closures definitely solve some common problems. However, the way that interfaces are implemented in Go can sometimes make it tricky to decide whether to use interfaces or closures for a particular problem. Usually, whether an interface or a closure is used isn't really of importance; the right choice is whichever one solves the problem at hand. Typically, closures will be simpler to implement if the operation is simple by nature. However, as soon as the logic contained within a closure becomes complex, one should strongly consider using an interface instead.  

Dave Cheney has an excellent write-up on this topic, as well as a talk:

* https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
* https://www.youtube.com/watch?v=5buaPyJ0XeQ&t=9s

Jon Bodner also has a related talk:

* https://www.youtube.com/watch?v=5IKcPMJXkKs

### Interfaces in Go

In general, Go's approach to handling `interface`s is quite different from those of other languages. Interfaces aren't explicitly implemented like they would be in Java or C#; rather, they are implicitly created if they fulfill the contract of the interface. As an example, this means that any `struct` that has an `Error()` method implements (or "fulfills") the `Error` interface and can be returned as an `error`. This manner of implementing interfaces is extremely easy and makes Go feel more fast paced and dynamic.

However, there are certainly disadvantages with this approach. As the interface implementation is no longer explicit, it can be difficult to see which interfaces are implemented by a struct. Therefore, it's common to define interfaces with as few methods as possible; this makes it easier to understand whether a particular struct fulfills the contract of the interface.

An alternative is to create constructors that return an interface rather than the concrete type:

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}

type NullWriter struct {}

func (writer *NullWriter) Write(data []byte) (n int, err error) {
    // do nothing
    return len(data), nil
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

The above function ensures that the `NullWriter` struct implements the `Writer` interface. If we were to delete the `Write` method from `NullWriter`, we would get a compilation error. This is a good way of ensuring that our code behaves as expected and that we can rely on the compiler as a safety net in case we try to write invalid code.

In certain cases, it might not be desirable to write a constructor, or perhaps we would like for our constructor to return the concrete type, rather than the interface. As an example, the `NullWriter` struct has no properties to populate on initialisation, so writing a constructor is a little redundant. Therefore, we can use the less verbose method of checking interface compatibility:

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}

type NullWriter struct {}
var _ io.Writer = &NullWriter{}
```

In the above code, we are initialising a variable with the Go `blank identifier`, with the type assignment of `io.Writer`. This results in our variable being checked to fulfill the `io.Writer` interface contract, before being discarded. This method of checking interface fulfillment also makes it possible to check that several interface contracts are fulfilled:

```go
type NullReaderWriter struct{}
var _ io.Writer = &NullWriter{}
var _ io.Reader = &NullWriter{}
```

From the above code, it's very easy to understand which interfaces must be fulfilled; this ensures that the compiler will help us out during compile time. Therefore, this is generally the preferred solution for checking interface contract fulfillment.

There's yet another method of trying to be more explicit about which interfaces a given struct implements. However, this third method actually achieves the opposite of what we want. It involves using embedded interfaces as a struct property.

> <em>Wait what? &ndash; Presumably most people</em>

Let's rewind a bit before we dive deep into the forbidden forest of smelly Go. In Go, we can use embedded structs as a type of inheritance in our struct definitions. This is really nice, as we can decouple our code by defining reusable structs.

```go
type Metadata struct {
    CreatedBy types.User
}

type Document struct {
    *Metadata
    Title string
    Body string
}

type AudioFile struct {
    *Metadata
    Title string
    Body string
}
```

Above, we are defining a `Metadata` object that will provide us with property fields that we are likely to use on many different struct types. The neat thing about using the embedded struct, rather than explicitly defining the properties directly in our struct, is that it has decoupled the `Metadata` fields. Should we choose to update our `Metadata` object, we can change it in just a single place. As we've seen several times so far, we want to ensure that a change in one place in our code doesn't break other parts. Keeping these properties centralised makes it clear that structures with an embedded `Metadata` have the same properties&mdash;much like how structures that fulfill interfaces have the same methods.

Now, let's look at an example of how we can use a constructor to further prevent breaking our code when making changes to our `Metadata` struct:

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
    }
}

func NewDocument(title string, body string) Document {
    return Document{
        Metadata: NewMetadata(),
        Title: title,
        Body: body,
    }
}
```

Suppose that at a later point in time, we decide that we'd also like a `CreatedAt` field on our `Metadata` object. We can now easily achieve this by simply updating our `NewMetadata` constructor:

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
        CreatedAt: time.Now(),
    }
}
```

Now, both our `Document` and `AudioFile` structures are updated to also populate these fields on construction. This is the core principle behind decoupling and an excellent example of ensuring maintainability of code. We can also add new methods without breaking our existing code:

```go
type Metadata struct {
    CreatedBy types.User
    CreatedAt time.Time
    UpdatedBy types.User
    UpdatedAt time.Time
}

func (metadata *Metadata) AddUpdateInfo(user types.User) {
    metadata.UpdatedBy = user
    metadata.UpdatedAt = time.Now()
}
```

Again, without breaking the rest of our codebase, we've managed to introduce new functionality. This kind of programming makes implementing new features very quick and painless, which is exactly what we are trying to achieve by writing clean code.

Let's return to the topic of interface contract fulfillment using embedded interfaces. Consider the following code as an example:

```go
type NullWriter struct {
    Writer
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

The above code compiles. Technically, we are implementing the interface of `Writer` in our `NullWriter`, as `NullWriter` will inherit all the functions that are associated with this interface. Some see this as a clear way of showing that our `NullWriter` is implementing the `Writer` interface. However, we must be careful when using this technique.

```go
func main() {
    w := NewNullWriter()

    w.Write([]byte{1, 2, 3})
}
```

As mentioned before, the above code will compile. The `NewNullWriter` returns a `Writer`, and everything is hunky-dory according to the compiler because `NullWriter` fulfills the contract of `io.Writer`, via the embedded interface. However, running the code above will result in the following:

> panic: runtime error: invalid memory address or nil pointer dereference

What happened? An interface method in Go is essentially a function pointer. In this case, since we are pointing to the function of an interface, rather than an actual method implementation, we are trying to invoke a function that's actually a `nil` pointer. To prevent this from happening, we would have to provide the `NulllWriter` with a struct that fulfills the interface contract, with actual implemented methods.

```go
func main() {
  w := NullWriter{
    Writer: &bytes.Buffer{},
  }

    w.Write([]byte{1, 2, 3})
}
```

> NOTE: In the above example, `Writer` is referring to the embedded `io.Writer` interface. It is also possible to invoke the `Write` method by accessing this property with `w.Writer.Write()`.

We are no longer triggering a panic and can now use the `NullWriter` as a `Writer`. This initialisation process is not much different from having properties that are initialised as `nil`, as discussed previously. Therefore, logically, we should try to handle them in a similar way. However, this is where embedded interfaces become a little difficult to work with. In a previous section, it was explained that the best way to handle potential `nil` values is to make the property in question private and create a public *getter* method. This way, we could ensure that our property is, in fact, not `nil`. Unfortunately, this is simply not possible with embedded interfaces, as they are by nature always public.

Another concern raised by using embedded interfaces is the potential confusion caused by partially overwritten interface methods: 

```go
type MyReadCloser struct {
  io.ReadCloser
}

func (closer *ReadCloser) Read(data []byte) { ... }

func main() {
  closer := MyReadCloser{}
  
  closer.Read([]byte{1, 2, 3}) 	// works fine
  closer.Close() 		// causes panic
  closer.ReadCloser.Closer() 		// no panic 
}
```

Even though this might look like we're overriding methods, which is common in languages such as C# and Java, we actually aren't. Go doesn't support inheritance (and thus has no notion of a superclass). We can imitate the behaviour, but it is not a built-in part of the language. By using methods such as interface embedding without caution, we can create confusing and potentially buggy code, just to save a few more lines.

> NOTE: Some argue that using embedded interfaces is a good way of creating a mock structure for testing a subset of interface methods. Essentially, by using an embedded interface, you won't have to implement all of the methods of the interface; rather, you can choose to implement only the few methods that you'd like to test. Within the context of testing/mocking, I can see this argument, but I am still not a fan of this approach.

Let's quickly get back to clean code and proper usage of interfaces. It's time to discuss using interfaces as function parameters and return values. The most common proverb for interface usage with functions in Go is the following:

> <em>Be conservative in what you do; be liberal in what you accept from others &ndash; Jon Postel</em>

> FUN FACT: This proverb actually has nothing to do with Go. It's taken from an early specification of the TCP networking protocol.

In other words, you should write functions that accept an interface and return a concrete type. This is generally good practice and is especially useful when doing tests with mocking. As an example, we can create a function that takes a writer interface as its input and invokes the `Write` method of that interface:

```go
type Pipe struct {
    writer io.Writer
    buffer bytes.Buffer
}

func NewPipe(w io.Writer) *Pipe {
    return &Pipe{
        writer: w,
    }
} 

func (pipe *Pipe) Save() error {
    if _, err := pipe.writer.Write(pipe.FlushBuffer()); err != nil {
        return err
    }
    return nil
}
```

Let's assume that we are writing to a file when our application is running, but we don't want to write to a new file for all tests that invoke this function. We can implement a new mock type that will basically do nothing. Essentially, this is just basic dependency injection and mocking, but the point is that it is extremely easy to achieve in Go:

```go
type NullWriter struct {}

func (w *NullWriter) Write(data []byte) (int, error) {
    return len(data), nil
}

func TestFn(t *testing.T) {
    ...
    pipe := NewPipe(NullWriter{})
    ...
}
```

> NOTE: There is actually already a null writer implementation built into the `ioutil` package named `Discard`.

When constructing our `Pipe` struct with `NullWriter` (rather than a different writer), when invoking our `Save` function, nothing will happen. The only thing we had to do was add four lines of code. This is why it is encouraged to make interfaces as small as possible in idiomatic Go&mdash;it makes it especially easy to implement patterns like the one we just saw. However, this implementation of interfaces also comes with a <em>huge</em> downside.

### The Empty `interface{}`
Unlike other languages, Go does not have an implementation for generics. There have been many proposals for one, but all have been turned down by the Go language team. Unfortunately, without generics, developers must try to find creative alternatives, which very often involves using the empty `interface{}`. This section describes why these often <em>too</em> creative implementations should be considered bad practice and unclean code. There will also be examples of appropriate usage of the empty `interface{}` and how to avoid some pitfalls of writing code with it.

As mentioned in a previous section, Go determines whether a concrete type implements a particular interface by checking whether the type implements the <em>methods</em> of that interface. So what happens if our interface declares no methods, as is the case with the empty interface?

```go
type EmptyInterface interface {}
```

The above is equivalent to the built-in type `interface{}`. A natural consequence of this is that we can write generic functions that accept any type as arguments. This is extremely useful for certain kinds of functions, such as print helpers. Interestingly, this is actually what makes it possible to pass in any type to the `Println` function from the `fmt` package:

```go
func Println(v ...interface{}) {
    ...
}
```

In this case, `Println` isn't just accepting a single `interface{}`; rather, the function accepts a <em>slice</em> of types that implement the empty `interface{}`. As there are no methods associated with the empty `interface{}`, <em>all</em> types are accepted, even making it possible to feed `Println` with a slice of different types. This is a very common pattern when handling string conversion (both from and to a string). Good examples of this come from the `json` standard library package:

```go
func InsertItemHandler(w http.ResponseWriter, r *http.Request) {
    var item Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.InsertItem(item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatsOK)
}
```

All the less elegant code is contained within the `Decode` function. Thus, developers using this functionality won't have to worry about type reflection or type casting; we just have to worry about providing a pointer to a concrete type. This is good because the `Decode()` function is technically returning a concrete type. We are passing in our `Item` value, which will be populated from the body of the HTTP request. This means we won't have to deal with the potential risks of handling the `interface{}` value ourselves.

However, even when using the empty `interface{}` with good programming practices, we still have some issues. If we pass in a JSON string that has nothing to do with our `Item` type but is still valid JSON, we won't receive an error&mdash;our `item` variable will just be left with the default values. So, while we don't have to worry about reflection and casting errors, we will still have to make sure that the message sent from our client is a valid `Item` type. Unfortunately, as of writing this document, there is no simple or good way to implement these types of generic decoders without using the empty `interface{}` type.

The problem with using `interface{}` in this manner is that we are leaning towards using Go, a statically typed language, as a dynamically typed language. This becomes even clearer when looking at poor implementations of the `interface{}` type. The most common example of this comes from developers trying to implement a generic store or list of some sort.

Let's look at an example of trying to implement a generic HashMap package that can store any type using `interface{}`. 

```go
type HashMap struct {
    store map[string]interface{}
}

func (hashmap *HashMap) Insert(key string, value interface{}) {
    hashmap.store[key] = value
}

func (hashmap *HashMap) Get(key string) (interface{}, error) {
    value, ok := hashmap.store[key]
    if !ok {
        return nil, ErrKeyNotFoundInHashMap
    }
    return value
}
```

> NOTE: I have omitted thread safety from this example to keep it simple.

Please keep in mind that the implementation pattern shown above is actually used in quite a lot of Go packages. It is even used in the standard library `sync` package for the `sync.Map` type. So what's the problem with this implementation? Well, let's have a look at an example of using the package:

```go
func SomeFunction(id string) (Item, error) {
    itemIface, err := hashmap.Get(id)
    if err != nil {
        return EmptyItem, err
    }
    item, ok := itemIface.(Item)
    if !ok {
        return EmptyItem, ErrCastingItem
    }
    return item, nil
}
```

At first glance, this looks fine. However, we'll start getting into trouble if we add <em>different</em> types to our store, something that's currently allowed. There is nothing preventing us from adding something other than the `Item` type. So what happens when someone starts adding other types into our HashMap, like a pointer `*Item` instead of an `Item`? Our function now might return an error. Worst of all, this might not even be caught by our tests. Depending on the complexity of the system, this could introduce some bugs that are particularly difficult to debug.

This type of code should never reach production. Remember: Go does not (yet) support generics. That's just a fact that developers must accept for the time being. If we want to use generics, then we should use a different language that does support generics rather than relying on dangerous hacks.

So, how do we prevent this code from reaching production? The simplest solution is to just write the functions with concrete types instead of using `interface{}` values. Of course, this is not always the best approach, as there might be some functionality within the package that is not trivial to implement ourselves. Therefore, a better approach may be to create wrappers that expose the functionality we need but still ensure type safety:

```go
type ItemCache struct {
  kv tinykv.KV
} 

func (cache *ItemCache) Get(id string) (Item, error) {
  value, ok := cache.kv.Get(id)
  if !ok {
    return EmptyItem, ErrItemNotFound
  }
  return interfaceToItem(value)
}

func interfaceToItem(v interface{}) (Item, error) {
  item, ok := v.(Item)
  if !ok {
    return EmptyItem, ErrCouldNotCastItem
  }
  return item, nil
}

func (cache *ItemCache) Put(id string, item Item) error {
  return cache.kv.Put(id, item)
}
```

> NOTE: Implementations of other functionalities of the `tinykv.KV` cache have been omitted for the sake of brevity.

The wrapper above now ensures that we are using the actual types and that we are no longer passing in `interface{}` types. It is therefore no longer possible to accidentally populate our store with a wrong value type, and we have restricted our casting of types as much as possible. This is a very straightforward way of solving our issue, even if somewhat manually.

## Summary

First of all, thank you for making it all the way through this article! I hope it has provided some insight into clean code and how it helps ensure maintainability, readability, and stability in any codebase.

Let's briefly sum up all the topics we've covered:

- <strong>Functions</strong>&mdash;A function's name should reflect its scope; the smaller the scope of a function, the more specific its name. Ensure that all functions serve a single purpose in as few lines as possible. A good rule of thumb is to limit your functions to 5&ndash;8 lines and to only accept 2&ndash;3 arguments.

- <strong>Variables</strong>&mdash;Unlike functions, variables should assume more generic names as their scope becomes smaller. It's also recommended that you limit the scope of a variable as much as possible to prevent unintentional modification. On a similar note, you should keep the modification of variables to a minimum; this becomes an especially important consideration as the scope of a variable grows.

- <strong>Return Values</strong>&mdash;Concrete types should be returned whenever possible. Make it as difficult as possible for users of your package to make mistakes and as easy as possible for them to understand the values returned by your functions.

- <strong>Pointers</strong>&mdash;Use pointers with caution, and limit their scope and mutability to an absolute minimum. Remember: Garbage collection only assists with memory management; it does not assist with all of the other complexities associated with pointers.

- <strong>Interfaces</strong>&mdash;Use interfaces as much as possible to loosen the coupling of your code. Hide any code using the empty `interface{}` as much as possible from end users to prevent it from being exposed.

As a final note, it's worth mentioning that the notion of clean code is particularly subjective, and that likely won't ever change. However, much like my statement concerning `gofmt`, I think it's more important to find a common standard than something that everyone agrees with; the latter is extremely difficult to achieve.

It's also important to understand that fanaticism is never the goal with clean code. A codebase will most likely never be fully 'clean,' in the same way that your office desk probably isn't either. There's certainly room for you to step outside the rules and boundaries covered in this article. However, remember that the most important reason for writing clean code is to help yourself and other developers. We support engineers by ensuring stability in the software we produce and by making it easier to debug faulty code. We help our fellow developers by ensuring that our code is readable and easily digestible. We help <em>everyone</em> involved in the project by establishing a flexible codebase that allows us to quickly introduce new features without breaking our current platform. We move quickly by going slowly, and everyone is satisfied.

I hope you will join this discussion to help the Go community define (and refine) the concept of clean code. Let's establish a common ground so that we can improve software&mdash;not only for ourselves but for the sake of everyone.

<!--
Editing this document:
- Discuss all changes in GitHub issues first.
- Update the table of contents as new sections are added or removed.
- Use tables for side-by-side code samples. See below.
Code Samples:
Use 2 spaces to indent. Horizontal real estate is important in side-by-side
samples.
For side-by-side code samples, use the following snippet.
~~~
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
BAD CODE GOES HERE
```
</td><td>
```go
GOOD CODE GOES HERE
```
</td></tr>
</tbody></table>
~~~
(You need the empty lines between the <td> and code samples for it to be
treated as Markdown.)
If you need to add labels or descriptions below the code samples, add another
row before the </tbody></table> line.
~~~
<tr>
<td>DESCRIBE BAD CODE</td>
<td>DESCRIBE GOOD CODE</td>
</tr>
~~~
-->
# Uber Go Style Guide
## Table of Contents
- [Introduction](#introduction)
- [Guidelines](#guidelines)
  - [Pointers to Interfaces](#pointers-to-interfaces)
  - [Verify Interface Compliance](#verify-interface-compliance)
  - [Receivers and Interfaces](#receivers-and-interfaces)
  - [Zero-value Mutexes are Valid](#zero-value-mutexes-are-valid)
  - [Copy Slices and Maps at Boundaries](#copy-slices-and-maps-at-boundaries)
  - [Defer to Clean Up](#defer-to-clean-up)
  - [Channel Size is One or None](#channel-size-is-one-or-none)
  - [Start Enums at One](#start-enums-at-one)
  - [Use `"time"` to handle time](#use-time-to-handle-time)
  - [Errors](#errors)
    - [Error Types](#error-types)
    - [Error Wrapping](#error-wrapping)
    - [Error Naming](#error-naming)
  - [Handle Type Assertion Failures](#handle-type-assertion-failures)
  - [Don't Panic](#dont-panic)
  - [Use go.uber.org/atomic](#use-gouberorgatomic)
  - [Avoid Mutable Globals](#avoid-mutable-globals)
  - [Avoid Embedding Types in Public Structs](#avoid-embedding-types-in-public-structs)
  - [Avoid Using Built-In Names](#avoid-using-built-in-names)
  - [Avoid `init()`](#avoid-init)
  - [Exit in Main](#exit-in-main)
    - [Exit Once](#exit-once)
  - [Use field tags in marshaled structs](#use-field-tags-in-marshaled-structs)
- [Performance](#performance)
  - [Prefer strconv over fmt](#prefer-strconv-over-fmt)
  - [Avoid string-to-byte conversion](#avoid-string-to-byte-conversion)
  - [Prefer Specifying Container Capacity](#prefer-specifying-container-capacity)
      - [Specifying Map Capacity Hints](#specifying-map-capacity-hints)
      - [Specifying Slice Capacity](#specifying-slice-capacity)
- [Style](#style)
  - [Avoid overly long lines](#avoid-overly-long-lines)
  - [Be Consistent](#be-consistent)
  - [Group Similar Declarations](#group-similar-declarations)
  - [Import Group Ordering](#import-group-ordering)
  - [Package Names](#package-names)
  - [Function Names](#function-names)
  - [Import Aliasing](#import-aliasing)
  - [Function Grouping and Ordering](#function-grouping-and-ordering)
  - [Reduce Nesting](#reduce-nesting)
  - [Unnecessary Else](#unnecessary-else)
  - [Top-level Variable Declarations](#top-level-variable-declarations)
  - [Prefix Unexported Globals with _](#prefix-unexported-globals-with-_)
  - [Embedding in Structs](#embedding-in-structs)
  - [Local Variable Declarations](#local-variable-declarations)
  - [nil is a valid slice](#nil-is-a-valid-slice)
  - [Reduce Scope of Variables](#reduce-scope-of-variables)
  - [Avoid Naked Parameters](#avoid-naked-parameters)
  - [Use Raw String Literals to Avoid Escaping](#use-raw-string-literals-to-avoid-escaping)
  - [Initializing Structs](#initializing-structs)
      - [Use Field Names to Initialize Structs](#use-field-names-to-initialize-structs)
      - [Omit Zero Value Fields in Structs](#omit-zero-value-fields-in-structs)
      - [Use `var` for Zero Value Structs](#use-var-for-zero-value-structs)
      - [Initializing Struct References](#initializing-struct-references)
  - [Initializing Maps](#initializing-maps)
  - [Format Strings outside Printf](#format-strings-outside-printf)
  - [Naming Printf-style Functions](#naming-printf-style-functions)
- [Patterns](#patterns)
  - [Test Tables](#test-tables)
  - [Functional Options](#functional-options)
- [Linting](#linting)
## Introduction
Styles are the conventions that govern our code. The term style is a bit of a
misnomer, since these conventions cover far more than just source file
formatting—gofmt handles that for us.
The goal of this guide is to manage this complexity by describing in detail the
Dos and Don'ts of writing Go code at Uber. These rules exist to keep the code
base manageable while still allowing engineers to use Go language features
productively.
This guide was originally created by [Prashant Varanasi] and [Simon Newton] as
a way to bring some colleagues up to speed with using Go. Over the years it has
been amended based on feedback from others.
  [Prashant Varanasi]: https://github.com/prashantv
  [Simon Newton]: https://github.com/nomis52
This documents idiomatic conventions in Go code that we follow at Uber. A lot
of these are general guidelines for Go, while others extend upon external
resources:
1. [Effective Go](https://golang.org/doc/effective_go.html)
2. [Go Common Mistakes](https://github.com/golang/go/wiki/CommonMistakes)
3. [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
We aim for the code samples to be accurate for the two most recent minor versions
of Go [releases](https://go.dev/doc/devel/release).
All code should be error-free when run through `golint` and `go vet`. We
recommend setting up your editor to:
- Run `goimports` on save
- Run `golint` and `go vet` to check for errors
You can find information in editor support for Go tools here:
<https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins>
## Guidelines
### Pointers to Interfaces
You almost never need a pointer to an interface. You should be passing
interfaces as values—the underlying data can still be a pointer.
An interface is two fields:
1. A pointer to some type-specific information. You can think of this as
  "type."
2. Data pointer. If the data stored is a pointer, it’s stored directly. If
  the data stored is a value, then a pointer to the value is stored.
If you want interface methods to modify the underlying data, you must use a
pointer.
### Verify Interface Compliance
Verify interface compliance at compile time where appropriate. This includes:
- Exported types that are required to implement specific interfaces as part of
  their API contract
- Exported or unexported types that are part of a collection of types
  implementing the same interface
- Other cases where violating an interface would break users
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Handler struct {
  // ...
}
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
```
</td><td>
```go
type Handler struct {
  // ...
}
var _ http.Handler = (*Handler)(nil)
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```
</td></tr>
</tbody></table>
The statement `var _ http.Handler = (*Handler)(nil)` will fail to compile if
`*Handler` ever stops matching the `http.Handler` interface.
The right hand side of the assignment should be the zero value of the asserted
type. This is `nil` for pointer types (like `*Handler`), slices, and maps, and
an empty struct for struct types.
```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}
var _ http.Handler = LogHandler{}
func (h LogHandler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```
### Receivers and Interfaces
Methods with value receivers can be called on pointers as well as values.
Methods with pointer receivers can only be called on pointers or [addressable values].
  [addressable values]: https://golang.org/ref/spec#Method_values
For example,
```go
type S struct {
  data string
}
func (s S) Read() string {
  return s.data
}
func (s *S) Write(str string) {
  s.data = str
}
sVals := map[int]S{1: {"A"}}
// You can only call Read using a value
sVals[1].Read()
// This will not compile:
//  sVals[1].Write("test")
sPtrs := map[int]*S{1: {"A"}}
// You can call both Read and Write using a pointer
sPtrs[1].Read()
sPtrs[1].Write("test")
```
Similarly, an interface can be satisfied by a pointer, even if the method has a
value receiver.
```go
type F interface {
  f()
}
type S1 struct{}
func (s S1) f() {}
type S2 struct{}
func (s *S2) f() {}
s1Val := S1{}
s1Ptr := &S1{}
s2Val := S2{}
s2Ptr := &S2{}
var i F
i = s1Val
i = s1Ptr
i = s2Ptr
// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
//   i = s2Val
```
Effective Go has a good write up on [Pointers vs. Values].
  [Pointers vs. Values]: https://golang.org/doc/effective_go.html#pointers_vs_values
### Zero-value Mutexes are Valid
The zero-value of `sync.Mutex` and `sync.RWMutex` is valid, so you almost
never need a pointer to a mutex.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
mu := new(sync.Mutex)
mu.Lock()
```
</td><td>
```go
var mu sync.Mutex
mu.Lock()
```
</td></tr>
</tbody></table>
If you use a struct by pointer, then the mutex should be a non-pointer field on
it. Do not embed the mutex on the struct, even if the struct is not exported.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type SMap struct {
  sync.Mutex
  data map[string]string
}
func NewSMap() *SMap {
  return &SMap{
    data: make(map[string]string),
  }
}
func (m *SMap) Get(k string) string {
  m.Lock()
  defer m.Unlock()
  return m.data[k]
}
```
</td><td>
```go
type SMap struct {
  mu sync.Mutex
  data map[string]string
}
func NewSMap() *SMap {
  return &SMap{
    data: make(map[string]string),
  }
}
func (m *SMap) Get(k string) string {
  m.mu.Lock()
  defer m.mu.Unlock()
  return m.data[k]
}
```
</td></tr>
<tr><td>
The `Mutex` field, and the `Lock` and `Unlock` methods are unintentionally part
of the exported API of `SMap`.
</td><td>
The mutex and its methods are implementation details of `SMap` hidden from its
callers.
</td></tr>
</tbody></table>
### Copy Slices and Maps at Boundaries
Slices and maps contain pointers to the underlying data so be wary of scenarios
when they need to be copied.
#### Receiving Slices and Maps
Keep in mind that users can modify a map or slice you received as an argument
if you store a reference to it.
<table>
<thead><tr><th>Bad</th> <th>Good</th></tr></thead>
<tbody>
<tr>
<td>
```go
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = trips
}
trips := ...
d1.SetTrips(trips)
// Did you mean to modify d1.trips?
trips[0] = ...
```
</td>
<td>
```go
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = make([]Trip, len(trips))
  copy(d.trips, trips)
}
trips := ...
d1.SetTrips(trips)
// We can now modify trips[0] without affecting d1.trips.
trips[0] = ...
```
</td>
</tr>
</tbody>
</table>
#### Returning Slices and Maps
Similarly, be wary of user modifications to maps or slices exposing internal
state.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Stats struct {
  mu sync.Mutex
  counters map[string]int
}
// Snapshot returns the current stats.
func (s *Stats) Snapshot() map[string]int {
  s.mu.Lock()
  defer s.mu.Unlock()
  return s.counters
}
// snapshot is no longer protected by the mutex, so any
// access to the snapshot is subject to data races.
snapshot := stats.Snapshot()
```
</td><td>
```go
type Stats struct {
  mu sync.Mutex
  counters map[string]int
}
func (s *Stats) Snapshot() map[string]int {
  s.mu.Lock()
  defer s.mu.Unlock()
  result := make(map[string]int, len(s.counters))
  for k, v := range s.counters {
    result[k] = v
  }
  return result
}
// Snapshot is now a copy.
snapshot := stats.Snapshot()
```
</td></tr>
</tbody></table>
### Defer to Clean Up
Use defer to clean up resources such as files and locks.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
p.Lock()
if p.count < 10 {
  p.Unlock()
  return p.count
}
p.count++
newCount := p.count
p.Unlock()
return newCount
// easy to miss unlocks due to multiple returns
```
</td><td>
```go
p.Lock()
defer p.Unlock()
if p.count < 10 {
  return p.count
}
p.count++
return p.count
// more readable
```
</td></tr>
</tbody></table>
Defer has an extremely small overhead and should be avoided only if you can
prove that your function execution time is in the order of nanoseconds. The
readability win of using defers is worth the miniscule cost of using them. This
is especially true for larger methods that have more than simple memory
accesses, where the other computations are more significant than the `defer`.
### Channel Size is One or None
Channels should usually have a size of one or be unbuffered. By default,
channels are unbuffered and have a size of zero. Any other size
must be subject to a high level of scrutiny. Consider how the size is
determined, what prevents the channel from filling up under load and blocking
writers, and what happens when this occurs.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// Ought to be enough for anybody!
c := make(chan int, 64)
```
</td><td>
```go
// Size of one
c := make(chan int, 1) // or
// Unbuffered channel, size of zero
c := make(chan int)
```
</td></tr>
</tbody></table>
### Start Enums at One
The standard way of introducing enumerations in Go is to declare a custom type
and a `const` group with `iota`. Since variables have a 0 default value, you
should usually start your enums on a non-zero value.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Operation int
const (
  Add Operation = iota
  Subtract
  Multiply
)
// Add=0, Subtract=1, Multiply=2
```
</td><td>
```go
type Operation int
const (
  Add Operation = iota + 1
  Subtract
  Multiply
)
// Add=1, Subtract=2, Multiply=3
```
</td></tr>
</tbody></table>
There are cases where using the zero value makes sense, for example when the
zero value case is the desirable default behavior.
```go
type LogOutput int
const (
  LogToStdout LogOutput = iota
  LogToFile
  LogToRemote
)
// LogToStdout=0, LogToFile=1, LogToRemote=2
```
### Use `"time"` to handle time
Time is complicated. Incorrect assumptions often made about time include the
following.
1. A day has 24 hours
2. An hour has 60 minutes
3. A week has 7 days
4. A year has 365 days
5. [And a lot more](https://infiniteundo.com/post/25326999628/falsehoods-programmers-believe-about-time)
For example, *1* means that adding 24 hours to a time instant will not always
yield a new calendar day.
Therefore, always use the [`"time"`] package when dealing with time because it
helps deal with these incorrect assumptions in a safer, more accurate manner.
  [`"time"`]: https://golang.org/pkg/time/
#### Use `time.Time` for instants of time
Use [`time.Time`] when dealing with instants of time, and the methods on
`time.Time` when comparing, adding, or subtracting time.
  [`time.Time`]: https://golang.org/pkg/time/#Time
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func isActive(now, start, stop int) bool {
  return start <= now && now < stop
}
```
</td><td>
```go
func isActive(now, start, stop time.Time) bool {
  return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}
```
</td></tr>
</tbody></table>
#### Use `time.Duration` for periods of time
Use [`time.Duration`] when dealing with periods of time.
  [`time.Duration`]: https://golang.org/pkg/time/#Duration
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func poll(delay int) {
  for {
    // ...
    time.Sleep(time.Duration(delay) * time.Millisecond)
  }
}
poll(10) // was it seconds or milliseconds?
```
</td><td>
```go
func poll(delay time.Duration) {
  for {
    // ...
    time.Sleep(delay)
  }
}
poll(10*time.Second)
```
</td></tr>
</tbody></table>
Going back to the example of adding 24 hours to a time instant, the method we
use to add time depends on intent. If we want the same time of the day, but on
the next calendar day, we should use [`Time.AddDate`]. However, if we want an
instant of time guaranteed to be 24 hours after the previous time, we should
use [`Time.Add`].
  [`Time.AddDate`]: https://golang.org/pkg/time/#Time.AddDate
  [`Time.Add`]: https://golang.org/pkg/time/#Time.Add
```go
newDay := t.AddDate(0 /* years */, 0 /* months */, 1 /* days */)
maybeNewDay := t.Add(24 * time.Hour)
```
#### Use `time.Time` and `time.Duration` with external systems
Use `time.Duration` and `time.Time` in interactions with external systems when
possible. For example:
- Command-line flags: [`flag`] supports `time.Duration` via
  [`time.ParseDuration`]
- JSON: [`encoding/json`] supports encoding `time.Time` as an [RFC 3339]
  string via its [`UnmarshalJSON` method]
- SQL: [`database/sql`] supports converting `DATETIME` or `TIMESTAMP` columns
  into `time.Time` and back if the underlying driver supports it
- YAML: [`gopkg.in/yaml.v2`] supports `time.Time` as an [RFC 3339] string, and
  `time.Duration` via [`time.ParseDuration`].
  [`flag`]: https://golang.org/pkg/flag/
  [`time.ParseDuration`]: https://golang.org/pkg/time/#ParseDuration
  [`encoding/json`]: https://golang.org/pkg/encoding/json/
  [RFC 3339]: https://tools.ietf.org/html/rfc3339
  [`UnmarshalJSON` method]: https://golang.org/pkg/time/#Time.UnmarshalJSON
  [`database/sql`]: https://golang.org/pkg/database/sql/
  [`gopkg.in/yaml.v2`]: https://godoc.org/gopkg.in/yaml.v2
When it is not possible to use `time.Duration` in these interactions, use
`int` or `float64` and include the unit in the name of the field.
For example, since `encoding/json` does not support `time.Duration`, the unit
is included in the name of the field.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// {"interval": 2}
type Config struct {
  Interval int `json:"interval"`
}
```
</td><td>
```go
// {"intervalMillis": 2000}
type Config struct {
  IntervalMillis int `json:"intervalMillis"`
}
```
</td></tr>
</tbody></table>
When it is not possible to use `time.Time` in these interactions, unless an
alternative is agreed upon, use `string` and format timestamps as defined in
[RFC 3339]. This format is used by default by [`Time.UnmarshalText`] and is
available for use in `Time.Format` and `time.Parse` via [`time.RFC3339`].
  [`Time.UnmarshalText`]: https://golang.org/pkg/time/#Time.UnmarshalText
  [`time.RFC3339`]: https://golang.org/pkg/time/#RFC3339
Although this tends to not be a problem in practice, keep in mind that the
`"time"` package does not support parsing timestamps with leap seconds
([8728]), nor does it account for leap seconds in calculations ([15190]). If
you compare two instants of time, the difference will not include the leap
seconds that may have occurred between those two instants.
  [8728]: https://github.com/golang/go/issues/8728
  [15190]: https://github.com/golang/go/issues/15190
<!-- TODO: section on String methods for enums -->
### Errors
#### Error Types
There are few options for declaring errors.
Consider the following before picking the option best suited for your use case.
- Does the caller need to match the error so that they can handle it?
  If yes, we must support the [`errors.Is`] or [`errors.As`] functions
  by declaring a top-level error variable or a custom type.
- Is the error message a static string,
  or is it a dynamic string that requires contextual information?
  For the former, we can use [`errors.New`], but for the latter we must
  use [`fmt.Errorf`] or a custom error type.
- Are we propagating a new error returned by a downstream function?
  If so, see the [section on error wrapping](#error-wrapping).
[`errors.Is`]: https://golang.org/pkg/errors/#Is
[`errors.As`]: https://golang.org/pkg/errors/#As
| Error matching? | Error Message | Guidance                            |
|-----------------|---------------|-------------------------------------|
| No              | static        | [`errors.New`]                      |
| No              | dynamic       | [`fmt.Errorf`]                      |
| Yes             | static        | top-level `var` with [`errors.New`] |
| Yes             | dynamic       | custom `error` type                 |
[`errors.New`]: https://golang.org/pkg/errors/#New
[`fmt.Errorf`]: https://golang.org/pkg/fmt/#Errorf
For example,
use [`errors.New`] for an error with a static string.
Export this error as a variable to support matching it with `errors.Is`
if the caller needs to match and handle this error.
<table>
<thead><tr><th>No error matching</th><th>Error matching</th></tr></thead>
<tbody>
<tr><td>
```go
// package foo
func Open() error {
  return errors.New("could not open")
}
// package bar
if err := foo.Open(); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
```
</td><td>
```go
// package foo
var ErrCouldNotOpen = errors.New("could not open")
func Open() error {
  return ErrCouldNotOpen
}
// package bar
if err := foo.Open(); err != nil {
  if errors.Is(err, foo.ErrCouldNotOpen) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```
</td></tr>
</tbody></table>
For an error with a dynamic string,
use [`fmt.Errorf`] if the caller does not need to match it,
and a custom `error` if the caller does need to match it.
<table>
<thead><tr><th>No error matching</th><th>Error matching</th></tr></thead>
<tbody>
<tr><td>
```go
// package foo
func Open(file string) error {
  return fmt.Errorf("file %q not found", file)
}
// package bar
if err := foo.Open("testfile.txt"); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
```
</td><td>
```go
// package foo
type NotFoundError struct {
  File string
}
func (e *NotFoundError) Error() string {
  return fmt.Sprintf("file %q not found", e.File)
}
func Open(file string) error {
  return &NotFoundError{File: file}
}
// package bar
if err := foo.Open("testfile.txt"); err != nil {
  var notFound *NotFoundError
  if errors.As(err, &notFound) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```
</td></tr>
</tbody></table>
Note that if you export error variables or types from a package,
they will become part of the public API of the package.
#### Error Wrapping
There are three main options for propagating errors if a call fails:
- return the original error as-is
- add context with `fmt.Errorf` and the `%w` verb
- add context with `fmt.Errorf` and the `%v` verb
Return the original error as-is if there is no additional context to add.
This maintains the original error type and message.
This is well suited for cases when the underlying error message
has sufficient information to track down where it came from.
Otherwise, add context to the error message where possible
so that instead of a vague error such as "connection refused",
you get more useful errors such as "call service foo: connection refused".
Use `fmt.Errorf` to add context to your errors,
picking between the `%w` or `%v` verbs
based on whether the caller should be able to
match and extract the underlying cause.
- Use `%w` if the caller should have access to the underlying error.
  This is a good default for most wrapped errors,
  but be aware that callers may begin to rely on this behavior.
  So for cases where the wrapped error is a known `var` or type,
  document and test it as part of your function's contract.
- Use `%v` to obfuscate the underlying error.
  Callers will be unable to match it,
  but you can switch to `%w` in the future if needed.
When adding context to returned errors, keep the context succinct by avoiding
phrases like "failed to", which state the obvious and pile up as the error
percolates up through the stack:
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "failed to create new store: %w", err)
}
```
</td><td>
```go
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "new store: %w", err)
}
```
</td></tr><tr><td>
```
failed to x: failed to y: failed to create new store: the error
```
</td><td>
```
x: y: new store: the error
```
</td></tr>
</tbody></table>
However once the error is sent to another system, it should be clear the
message is an error (e.g. an `err` tag or "Failed" prefix in logs).
See also [Don't just check errors, handle them gracefully].
  [`"pkg/errors".Cause`]: https://godoc.org/github.com/pkg/errors#Cause
  [Don't just check errors, handle them gracefully]: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
#### Error Naming
For error values stored as global variables,
use the prefix `Err` or `err` depending on whether they're exported.
This guidance supersedes the [Prefix Unexported Globals with _](#prefix-unexported-globals-with-_).
```go
var (
  // The following two errors are exported
  // so that users of this package can match them
  // with errors.Is.
  ErrBrokenLink = errors.New("link is broken")
  ErrCouldNotOpen = errors.New("could not open")
  // This error is not exported because
  // we don't want to make it part of our public API.
  // We may still use it inside the package
  // with errors.Is.
  errNotFound = errors.New("not found")
)
```
For custom error types, use the suffix `Error` instead.
```go
// Similarly, this error is exported
// so that users of this package can match it
// with errors.As.
type NotFoundError struct {
  File string
}
func (e *NotFoundError) Error() string {
  return fmt.Sprintf("file %q not found", e.File)
}
// And this error is not exported because
// we don't want to make it part of the public API.
// We can still use it inside the package
// with errors.As.
type resolveError struct {
  Path string
}
func (e *resolveError) Error() string {
  return fmt.Sprintf("resolve %q", e.Path)
}
```
### Handle Type Assertion Failures
The single return value form of a [type assertion] will panic on an incorrect
type. Therefore, always use the "comma ok" idiom.
  [type assertion]: https://golang.org/ref/spec#Type_assertions
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
t := i.(string)
```
</td><td>
```go
t, ok := i.(string)
if !ok {
  // handle the error gracefully
}
```
</td></tr>
</tbody></table>
<!-- TODO: There are a few situations where the single assignment form is
fine. -->
### Don't Panic
Code running in production must avoid panics. Panics are a major source of
[cascading failures]. If an error occurs, the function must return an error and
allow the caller to decide how to handle it.
  [cascading failures]: https://en.wikipedia.org/wiki/Cascading_failure
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func run(args []string) {
  if len(args) == 0 {
    panic("an argument is required")
  }
  // ...
}
func main() {
  run(os.Args[1:])
}
```
</td><td>
```go
func run(args []string) error {
  if len(args) == 0 {
    return errors.New("an argument is required")
  }
  // ...
  return nil
}
func main() {
  if err := run(os.Args[1:]); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
```
</td></tr>
</tbody></table>
Panic/recover is not an error handling strategy. A program must panic only when
something irrecoverable happens such as a nil dereference. An exception to this is
program initialization: bad things at program startup that should abort the
program may cause panic.
```go
var _statusTemplate = template.Must(template.New("name").Parse("_statusHTML"))
```
Even in tests, prefer `t.Fatal` or `t.FailNow` over panics to ensure that the
test is marked as failed.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// func TestFoo(t *testing.T)
f, err := os.CreateTemp("", "test")
if err != nil {
  panic("failed to set up test")
}
```
</td><td>
```go
// func TestFoo(t *testing.T)
f, err := os.CreateTemp("", "test")
if err != nil {
  t.Fatal("failed to set up test")
}
```
</td></tr>
</tbody></table>
<!-- TODO: Explain how to use _test packages. -->
### Use go.uber.org/atomic
Atomic operations with the [sync/atomic] package operate on the raw types
(`int32`, `int64`, etc.) so it is easy to forget to use the atomic operation to
read or modify the variables.
[go.uber.org/atomic] adds type safety to these operations by hiding the
underlying type. Additionally, it includes a convenient `atomic.Bool` type.
  [go.uber.org/atomic]: https://godoc.org/go.uber.org/atomic
  [sync/atomic]: https://golang.org/pkg/sync/atomic/
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type foo struct {
  running int32  // atomic
}
func (f* foo) start() {
  if atomic.SwapInt32(&f.running, 1) == 1 {
     // already running…
     return
  }
  // start the Foo
}
func (f *foo) isRunning() bool {
  return f.running == 1  // race!
}
```
</td><td>
```go
type foo struct {
  running atomic.Bool
}
func (f *foo) start() {
  if f.running.Swap(true) {
     // already running…
     return
  }
  // start the Foo
}
func (f *foo) isRunning() bool {
  return f.running.Load()
}
```
</td></tr>
</tbody></table>
### Avoid Mutable Globals
Avoid mutating global variables, instead opting for dependency injection.
This applies to function pointers as well as other kinds of values.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// sign.go
var _timeNow = time.Now
func sign(msg string) string {
  now := _timeNow()
  return signWithTime(msg, now)
}
```
</td><td>
```go
// sign.go
type signer struct {
  now func() time.Time
}
func newSigner() *signer {
  return &signer{
    now: time.Now,
  }
}
func (s *signer) Sign(msg string) string {
  now := s.now()
  return signWithTime(msg, now)
}
```
</td></tr>
<tr><td>
```go
// sign_test.go
func TestSign(t *testing.T) {
  oldTimeNow := _timeNow
  _timeNow = func() time.Time {
    return someFixedTime
  }
  defer func() { _timeNow = oldTimeNow }()
  assert.Equal(t, want, sign(give))
}
```
</td><td>
```go
// sign_test.go
func TestSigner(t *testing.T) {
  s := newSigner()
  s.now = func() time.Time {
    return someFixedTime
  }
  assert.Equal(t, want, s.Sign(give))
}
```
</td></tr>
</tbody></table>
### Avoid Embedding Types in Public Structs
These embedded types leak implementation details, inhibit type evolution, and
obscure documentation.
Assuming you have implemented a variety of list types using a shared
`AbstractList`, avoid embedding the `AbstractList` in your concrete list
implementations.
Instead, hand-write only the methods to your concrete list that will delegate
to the abstract list.
```go
type AbstractList struct {}
// Add adds an entity to the list.
func (l *AbstractList) Add(e Entity) {
  // ...
}
// Remove removes an entity from the list.
func (l *AbstractList) Remove(e Entity) {
  // ...
}
```
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// ConcreteList is a list of entities.
type ConcreteList struct {
  *AbstractList
}
```
</td><td>
```go
// ConcreteList is a list of entities.
type ConcreteList struct {
  list *AbstractList
}
// Add adds an entity to the list.
func (l *ConcreteList) Add(e Entity) {
  l.list.Add(e)
}
// Remove removes an entity from the list.
func (l *ConcreteList) Remove(e Entity) {
  l.list.Remove(e)
}
```
</td></tr>
</tbody></table>
Go allows [type embedding] as a compromise between inheritance and composition.
The outer type gets implicit copies of the embedded type's methods.
These methods, by default, delegate to the same method of the embedded
instance.
  [type embedding]: https://golang.org/doc/effective_go.html#embedding
The struct also gains a field by the same name as the type.
So, if the embedded type is public, the field is public.
To maintain backward compatibility, every future version of the outer type must
keep the embedded type.
An embedded type is rarely necessary.
It is a convenience that helps you avoid writing tedious delegate methods.
Even embedding a compatible AbstractList *interface*, instead of the struct,
would offer the developer more flexibility to change in the future, but still
leak the detail that the concrete lists use an abstract implementation.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// AbstractList is a generalized implementation
// for various kinds of lists of entities.
type AbstractList interface {
  Add(Entity)
  Remove(Entity)
}
// ConcreteList is a list of entities.
type ConcreteList struct {
  AbstractList
}
```
</td><td>
```go
// ConcreteList is a list of entities.
type ConcreteList struct {
  list AbstractList
}
// Add adds an entity to the list.
func (l *ConcreteList) Add(e Entity) {
  l.list.Add(e)
}
// Remove removes an entity from the list.
func (l *ConcreteList) Remove(e Entity) {
  l.list.Remove(e)
}
```
</td></tr>
</tbody></table>
Either with an embedded struct or an embedded interface, the embedded type
places limits on the evolution of the type.
- Adding methods to an embedded interface is a breaking change.
- Removing methods from an embedded struct is a breaking change.
- Removing the embedded type is a breaking change.
- Replacing the embedded type, even with an alternative that satisfies the same
  interface, is a breaking change.
Although writing these delegate methods is tedious, the additional effort hides
an implementation detail, leaves more opportunities for change, and also
eliminates indirection for discovering the full List interface in
documentation.
### Avoid Using Built-In Names
The Go [language specification] outlines several built-in,
[predeclared identifiers] that should not be used as names within Go programs.
Depending on context, reusing these identifiers as names will either shadow
the original within the current lexical scope (and any nested scopes) or make
affected code confusing. In the best case, the compiler will complain; in the
worst case, such code may introduce latent, hard-to-grep bugs.
  [language specification]: https://golang.org/ref/spec
  [predeclared identifiers]: https://golang.org/ref/spec#Predeclared_identifiers
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
var error string
// `error` shadows the builtin
// or
func handleErrorMessage(error string) {
    // `error` shadows the builtin
}
```
</td><td>
```go
var errorMessage string
// `error` refers to the builtin
// or
func handleErrorMessage(msg string) {
    // `error` refers to the builtin
}
```
</td></tr>
<tr><td>
```go
type Foo struct {
    // While these fields technically don't
    // constitute shadowing, grepping for
    // `error` or `string` strings is now
    // ambiguous.
    error  error
    string string
}
func (f Foo) Error() error {
    // `error` and `f.error` are
    // visually similar
    return f.error
}
func (f Foo) String() string {
    // `string` and `f.string` are
    // visually similar
    return f.string
}
```
</td><td>
```go
type Foo struct {
    // `error` and `string` strings are
    // now unambiguous.
    err error
    str string
}
func (f Foo) Error() error {
    return f.err
}
func (f Foo) String() string {
    return f.str
}
```
</td></tr>
</tbody></table>
Note that the compiler will not generate errors when using predeclared
identifiers, but tools such as `go vet` should correctly point out these and
other cases of shadowing.
### Avoid `init()`
Avoid `init()` where possible. When `init()` is unavoidable or desirable, code
should attempt to:
1. Be completely deterministic, regardless of program environment or invocation.
2. Avoid depending on the ordering or side-effects of other `init()` functions.
   While `init()` ordering is well-known, code can change, and thus
   relationships between `init()` functions can make code brittle and
   error-prone.
3. Avoid accessing or manipulating global or environment state, such as machine
   information, environment variables, working directory, program
   arguments/inputs, etc.
4. Avoid I/O, including both filesystem, network, and system calls.
Code that cannot satisfy these requirements likely belongs as a helper to be
called as part of `main()` (or elsewhere in a program's lifecycle), or be
written as part of `main()` itself. In particular, libraries that are intended
to be used by other programs should take special care to be completely
deterministic and not perform "init magic".
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Foo struct {
    // ...
}
var _defaultFoo Foo
func init() {
    _defaultFoo = Foo{
        // ...
    }
}
```
</td><td>
```go
var _defaultFoo = Foo{
    // ...
}
// or, better, for testability:
var _defaultFoo = defaultFoo()
func defaultFoo() Foo {
    return Foo{
        // ...
    }
}
```
</td></tr>
<tr><td>
```go
type Config struct {
    // ...
}
var _config Config
func init() {
    // Bad: based on current directory
    cwd, _ := os.Getwd()
    // Bad: I/O
    raw, _ := os.ReadFile(
        path.Join(cwd, "config", "config.yaml"),
    )
    yaml.Unmarshal(raw, &_config)
}
```
</td><td>
```go
type Config struct {
    // ...
}
func loadConfig() Config {
    cwd, err := os.Getwd()
    // handle err
    raw, err := os.ReadFile(
        path.Join(cwd, "config", "config.yaml"),
    )
    // handle err
    var config Config
    yaml.Unmarshal(raw, &config)
    return config
}
```
</td></tr>
</tbody></table>
Considering the above, some situations in which `init()` may be preferable or
necessary might include:
- Complex expressions that cannot be represented as single assignments.
- Pluggable hooks, such as `database/sql` dialects, encoding type registries, etc.
- Optimizations to [Google Cloud Functions] and other forms of deterministic
  precomputation.
  [Google Cloud Functions]: https://cloud.google.com/functions/docs/bestpractices/tips#use_global_variables_to_reuse_objects_in_future_invocations
### Exit in Main
Go programs use [`os.Exit`] or [`log.Fatal*`] to exit immediately. (Panicking
is not a good way to exit programs, please [don't panic](#dont-panic).)
  [`os.Exit`]: https://golang.org/pkg/os/#Exit
  [`log.Fatal*`]: https://golang.org/pkg/log/#Fatal
Call one of `os.Exit` or `log.Fatal*` **only in `main()`**. All other
functions should return errors to signal failure.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func main() {
  body := readFile(path)
  fmt.Println(body)
}
func readFile(path string) string {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  b, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }
  return string(b)
}
```
</td><td>
```go
func main() {
  body, err := readFile(path)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(body)
}
func readFile(path string) (string, error) {
  f, err := os.Open(path)
  if err != nil {
    return "", err
  }
  b, err := io.ReadAll(f)
  if err != nil {
    return "", err
  }
  return string(b), nil
}
```
</td></tr>
</tbody></table>
Rationale: Programs with multiple functions that exit present a few issues:
- Non-obvious control flow: Any function can exit the program so it becomes
  difficult to reason about the control flow.
- Difficult to test: A function that exits the program will also exit the test
  calling it. This makes the function difficult to test and introduces risk of
  skipping other tests that have not yet been run by `go test`.
- Skipped cleanup: When a function exits the program, it skips function calls
  enqueued with `defer` statements. This adds risk of skipping important
  cleanup tasks.
#### Exit Once
If possible, prefer to call `os.Exit` or `log.Fatal` **at most once** in your
`main()`. If there are multiple error scenarios that halt program execution,
put that logic under a separate function and return errors from it.
This has the effect of shortening your `main()` function and putting all key
business logic into a separate, testable function.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
package main
func main() {
  args := os.Args[1:]
  if len(args) != 1 {
    log.Fatal("missing file")
  }
  name := args[0]
  f, err := os.Open(name)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  // If we call log.Fatal after this line,
  // f.Close will not be called.
  b, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }
  // ...
}
```
</td><td>
```go
package main
func main() {
  if err := run(); err != nil {
    log.Fatal(err)
  }
}
func run() error {
  args := os.Args[1:]
  if len(args) != 1 {
    return errors.New("missing file")
  }
  name := args[0]
  f, err := os.Open(name)
  if err != nil {
    return err
  }
  defer f.Close()
  b, err := io.ReadAll(f)
  if err != nil {
    return err
  }
  // ...
}
```
</td></tr>
</tbody></table>
### Use field tags in marshaled structs
Any struct field that is marshaled into JSON, YAML,
or other formats that support tag-based field naming
should be annotated with the relevant tag.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Stock struct {
  Price int
  Name  string
}
bytes, err := json.Marshal(Stock{
  Price: 137,
  Name:  "UBER",
})
```
</td><td>
```go
type Stock struct {
  Price int    `json:"price"`
  Name  string `json:"name"`
  // Safe to rename Name to Symbol.
}
bytes, err := json.Marshal(Stock{
  Price: 137,
  Name:  "UBER",
})
```
</td></tr>
</tbody></table>
Rationale:
The serialized form of the structure is a contract between different systems.
Changes to the structure of the serialized form--including field names--break
this contract. Specifying field names inside tags makes the contract explicit,
and it guards against accidentally breaking the contract by refactoring or
renaming fields.
## Performance
Performance-specific guidelines apply only to the hot path.
### Prefer strconv over fmt
When converting primitives to/from strings, `strconv` is faster than
`fmt`.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
for i := 0; i < b.N; i++ {
  s := fmt.Sprint(rand.Int())
}
```
</td><td>
```go
for i := 0; i < b.N; i++ {
  s := strconv.Itoa(rand.Int())
}
```
</td></tr>
<tr><td>
```
BenchmarkFmtSprint-4    143 ns/op    2 allocs/op
```
</td><td>
```
BenchmarkStrconv-4    64.2 ns/op    1 allocs/op
```
</td></tr>
</tbody></table>
### Avoid string-to-byte conversion
Do not create byte slices from a fixed string repeatedly. Instead, perform the
conversion once and capture the result.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
for i := 0; i < b.N; i++ {
  w.Write([]byte("Hello world"))
}
```
</td><td>
```go
data := []byte("Hello world")
for i := 0; i < b.N; i++ {
  w.Write(data)
}
```
</td></tr>
<tr><td>
```
BenchmarkBad-4   50000000   22.2 ns/op
```
</td><td>
```
BenchmarkGood-4  500000000   3.25 ns/op
```
</td></tr>
</tbody></table>
### Prefer Specifying Container Capacity
Specify container capacity where possible in order to allocate memory for the
container up front. This minimizes subsequent allocations (by copying and
resizing of the container) as elements are added.
#### Specifying Map Capacity Hints
Where possible, provide capacity hints when initializing
maps with `make()`.
```go
make(map[T1]T2, hint)
```
Providing a capacity hint to `make()` tries to right-size the
map at initialization time, which reduces the need for growing
the map and allocations as elements are added to the map.
Note that, unlike slices, map capacity hints do not guarantee complete,
preemptive allocation, but are used to approximate the number of hashmap buckets
required. Consequently, allocations may still occur when adding elements to the
map, even up to the specified capacity.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
m := make(map[string]os.FileInfo)
files, _ := os.ReadDir("./files")
for _, f := range files {
    m[f.Name()] = f
}
```
</td><td>
```go
files, _ := os.ReadDir("./files")
m := make(map[string]os.DirEntry, len(files))
for _, f := range files {
    m[f.Name()] = f
}
```
</td></tr>
<tr><td>
`m` is created without a size hint; there may be more
allocations at assignment time.
</td><td>
`m` is created with a size hint; there may be fewer
allocations at assignment time.
</td></tr>
</tbody></table>
#### Specifying Slice Capacity
Where possible, provide capacity hints when initializing slices with `make()`,
particularly when appending.
```go
make([]T, length, capacity)
```
Unlike maps, slice capacity is not a hint: the compiler will allocate enough
memory for the capacity of the slice as provided to `make()`, which means that
subsequent `append()` operations will incur zero allocations (until the length
of the slice matches the capacity, after which any appends will require a resize
to hold additional elements).
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
for n := 0; n < b.N; n++ {
  data := make([]int, 0)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}
```
</td><td>
```go
for n := 0; n < b.N; n++ {
  data := make([]int, 0, size)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}
```
</td></tr>
<tr><td>
```
BenchmarkBad-4    100000000    2.48s
```
</td><td>
```
BenchmarkGood-4   100000000    0.21s
```
</td></tr>
</tbody></table>
## Style
### Avoid overly long lines
Avoid lines of code that require readers to scroll horizontally
or turn their heads too much.
We recommend a soft line length limit of **99 characters**.
Authors should aim to wrap lines before hitting this limit,
but it is not a hard limit.
Code is allowed to exceed this limit.
### Be Consistent
Some of the guidelines outlined in this document can be evaluated objectively;
others are situational, contextual, or subjective.
Above all else, **be consistent**.
Consistent code is easier to maintain, is easier to rationalize, requires less
cognitive overhead, and is easier to migrate or update as new conventions emerge
or classes of bugs are fixed.
Conversely, having multiple disparate or conflicting styles within a single
codebase causes maintenance overhead, uncertainty, and cognitive dissonance,
all of which can directly contribute to lower velocity, painful code reviews,
and bugs.
When applying these guidelines to a codebase, it is recommended that changes
are made at a package (or larger) level: application at a sub-package level
violates the above concern by introducing multiple styles into the same code.
### Group Similar Declarations
Go supports grouping similar declarations.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
import "a"
import "b"
```
</td><td>
```go
import (
  "a"
  "b"
)
```
</td></tr>
</tbody></table>
This also applies to constants, variables, and type declarations.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
const a = 1
const b = 2
var a = 1
var b = 2
type Area float64
type Volume float64
```
</td><td>
```go
const (
  a = 1
  b = 2
)
var (
  a = 1
  b = 2
)
type (
  Area float64
  Volume float64
)
```
</td></tr>
</tbody></table>
Only group related declarations. Do not group declarations that are unrelated.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Operation int
const (
  Add Operation = iota + 1
  Subtract
  Multiply
  EnvVar = "MY_ENV"
)
```
</td><td>
```go
type Operation int
const (
  Add Operation = iota + 1
  Subtract
  Multiply
)
const EnvVar = "MY_ENV"
```
</td></tr>
</tbody></table>
Groups are not limited in where they can be used. For example, you can use them
inside of functions.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func f() string {
  red := color.New(0xff0000)
  green := color.New(0x00ff00)
  blue := color.New(0x0000ff)
  // ...
}
```
</td><td>
```go
func f() string {
  var (
    red   = color.New(0xff0000)
    green = color.New(0x00ff00)
    blue  = color.New(0x0000ff)
  )
  // ...
}
```
</td></tr>
</tbody></table>
Exception: Variable declarations, particularly inside functions, should be
grouped together if declared adjacent to other variables. Do this for variables
declared together even if they are unrelated.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func (c *client) request() {
  caller := c.name
  format := "json"
  timeout := 5*time.Second
  var err error
  // ...
}
```
</td><td>
```go
func (c *client) request() {
  var (
    caller  = c.name
    format  = "json"
    timeout = 5*time.Second
    err error
  )
  // ...
}
```
</td></tr>
</tbody></table>
### Import Group Ordering
There should be two import groups:
- Standard library
- Everything else
This is the grouping applied by goimports by default.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
import (
  "fmt"
  "os"
  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```
</td><td>
```go
import (
  "fmt"
  "os"
  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```
</td></tr>
</tbody></table>
### Package Names
When naming packages, choose a name that is:
- All lower-case. No capitals or underscores.
- Does not need to be renamed using named imports at most call sites.
- Short and succinct. Remember that the name is identified in full at every call
  site.
- Not plural. For example, `net/url`, not `net/urls`.
- Not "common", "util", "shared", or "lib". These are bad, uninformative names.
See also [Package Names] and [Style guideline for Go packages].
  [Package Names]: https://blog.golang.org/package-names
  [Style guideline for Go packages]: https://rakyll.org/style-packages/
### Function Names
We follow the Go community's convention of using [MixedCaps for function
names]. An exception is made for test functions, which may contain underscores
for the purpose of grouping related test cases, e.g.,
`TestMyFunction_WhatIsBeingTested`.
  [MixedCaps for function names]: https://golang.org/doc/effective_go.html#mixed-caps
### Import Aliasing
Import aliasing must be used if the package name does not match the last
element of the import path.
```go
import (
  "net/http"
  client "example.com/client-go"
  trace "example.com/trace/v2"
)
```
In all other scenarios, import aliases should be avoided unless there is a
direct conflict between imports.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
import (
  "fmt"
  "os"
  nettrace "golang.net/x/trace"
)
```
</td><td>
```go
import (
  "fmt"
  "os"
  "runtime/trace"
  nettrace "golang.net/x/trace"
)
```
</td></tr>
</tbody></table>
### Function Grouping and Ordering
- Functions should be sorted in rough call order.
- Functions in a file should be grouped by receiver.
Therefore, exported functions should appear first in a file, after
`struct`, `const`, `var` definitions.
A `newXYZ()`/`NewXYZ()` may appear after the type is defined, but before the
rest of the methods on the receiver.
Since functions are grouped by receiver, plain utility functions should appear
towards the end of the file.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func (s *something) Cost() {
  return calcCost(s.weights)
}
type something struct{ ... }
func calcCost(n []int) int {...}
func (s *something) Stop() {...}
func newSomething() *something {
    return &something{}
}
```
</td><td>
```go
type something struct{ ... }
func newSomething() *something {
    return &something{}
}
func (s *something) Cost() {
  return calcCost(s.weights)
}
func (s *something) Stop() {...}
func calcCost(n []int) int {...}
```
</td></tr>
</tbody></table>
### Reduce Nesting
Code should reduce nesting where possible by handling error cases/special
conditions first and returning early or continuing the loop. Reduce the amount
of code that is nested multiple levels.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
for _, v := range data {
  if v.F1 == 1 {
    v = process(v)
    if err := v.Call(); err == nil {
      v.Send()
    } else {
      return err
    }
  } else {
    log.Printf("Invalid v: %v", v)
  }
}
```
</td><td>
```go
for _, v := range data {
  if v.F1 != 1 {
    log.Printf("Invalid v: %v", v)
    continue
  }
  v = process(v)
  if err := v.Call(); err != nil {
    return err
  }
  v.Send()
}
```
</td></tr>
</tbody></table>
### Unnecessary Else
If a variable is set in both branches of an if, it can be replaced with a
single if.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
var a int
if b {
  a = 100
} else {
  a = 10
}
```
</td><td>
```go
a := 10
if b {
  a = 100
}
```
</td></tr>
</tbody></table>
### Top-level Variable Declarations
At the top level, use the standard `var` keyword. Do not specify the type,
unless it is not the same type as the expression.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
var _s string = F()
func F() string { return "A" }
```
</td><td>
```go
var _s = F()
// Since F already states that it returns a string, we don't need to specify
// the type again.
func F() string { return "A" }
```
</td></tr>
</tbody></table>
Specify the type if the type of the expression does not match the desired type
exactly.
```go
type myError struct{}
func (myError) Error() string { return "error" }
func F() myError { return myError{} }
var _e error = F()
// F returns an object of type myError but we want error.
```
### Prefix Unexported Globals with _
Prefix unexported top-level `var`s and `const`s with `_` to make it clear when
they are used that they are global symbols.
Rationale: Top-level variables and constants have a package scope. Using a
generic name makes it easy to accidentally use the wrong value in a different
file.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// foo.go
const (
  defaultPort = 8080
  defaultUser = "user"
)
// bar.go
func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)
  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```
</td><td>
```go
// foo.go
const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```
</td></tr>
</tbody></table>
**Exception**: Unexported error values may use the prefix `err` without the underscore.
See [Error Naming](#error-naming).
### Embedding in Structs
Embedded types should be at the top of the field list of a
struct, and there must be an empty line separating embedded fields from regular
fields.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type Client struct {
  version int
  http.Client
}
```
</td><td>
```go
type Client struct {
  http.Client
  version int
}
```
</td></tr>
</tbody></table>
Embedding should provide tangible benefit, like adding or augmenting
functionality in a semantically-appropriate way. It should do this with zero
adverse user-facing effects (see also: [Avoid Embedding Types in Public Structs]).
Exception: Mutexes should not be embedded, even on unexported types. See also: [Zero-value Mutexes are Valid].
  [Avoid Embedding Types in Public Structs]: #avoid-embedding-types-in-public-structs
  [Zero-value Mutexes are Valid]: #zero-value-mutexes-are-valid
Embedding **should not**:
- Be purely cosmetic or convenience-oriented.
- Make outer types more difficult to construct or use.
- Affect outer types' zero values. If the outer type has a useful zero value, it
  should still have a useful zero value after embedding the inner type.
- Expose unrelated functions or fields from the outer type as a side-effect of
  embedding the inner type.
- Expose unexported types.
- Affect outer types' copy semantics.
- Change the outer type's API or type semantics.
- Embed a non-canonical form of the inner type.
- Expose implementation details of the outer type.
- Allow users to observe or control type internals.
- Change the general behavior of inner functions through wrapping in a way that
  would reasonably surprise users.
Simply put, embed consciously and intentionally. A good litmus test is, "would
all of these exported inner methods/fields be added directly to the outer type";
if the answer is "some" or "no", don't embed the inner type - use a field
instead.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
type A struct {
    // Bad: A.Lock() and A.Unlock() are
    //      now available, provide no
    //      functional benefit, and allow
    //      users to control details about
    //      the internals of A.
    sync.Mutex
}
```
</td><td>
```go
type countingWriteCloser struct {
    // Good: Write() is provided at this
    //       outer layer for a specific
    //       purpose, and delegates work
    //       to the inner type's Write().
    io.WriteCloser
    count int
}
func (w *countingWriteCloser) Write(bs []byte) (int, error) {
    w.count += len(bs)
    return w.WriteCloser.Write(bs)
}
```
</td></tr>
<tr><td>
```go
type Book struct {
    // Bad: pointer changes zero value usefulness
    io.ReadWriter
    // other fields
}
// later
var b Book
b.Read(...)  // panic: nil pointer
b.String()   // panic: nil pointer
b.Write(...) // panic: nil pointer
```
</td><td>
```go
type Book struct {
    // Good: has useful zero value
    bytes.Buffer
    // other fields
}
// later
var b Book
b.Read(...)  // ok
b.String()   // ok
b.Write(...) // ok
```
</td></tr>
<tr><td>
```go
type Client struct {
    sync.Mutex
    sync.WaitGroup
    bytes.Buffer
    url.URL
}
```
</td><td>
```go
type Client struct {
    mtx sync.Mutex
    wg  sync.WaitGroup
    buf bytes.Buffer
    url url.URL
}
```
</td></tr>
</tbody></table>
### Local Variable Declarations
Short variable declarations (`:=`) should be used if a variable is being set to
some value explicitly.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
var s = "foo"
```
</td><td>
```go
s := "foo"
```
</td></tr>
</tbody></table>
However, there are cases where the default value is clearer when the `var`
keyword is used. [Declaring Empty Slices], for example.
  [Declaring Empty Slices]: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
func f(list []int) {
  filtered := []int{}
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
```
</td><td>
```go
func f(list []int) {
  var filtered []int
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
```
</td></tr>
</tbody></table>
### nil is a valid slice
`nil` is a valid slice of length 0. This means that,
- You should not return a slice of length zero explicitly. Return `nil`
  instead.
  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>
  ```go
  if x == "" {
    return []int{}
  }
  ```
  </td><td>
  ```go
  if x == "" {
    return nil
  }
  ```
  </td></tr>
  </tbody></table>
- To check if a slice is empty, always use `len(s) == 0`. Do not check for
  `nil`.
  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>
  ```go
  func isEmpty(s []string) bool {
    return s == nil
  }
  ```
  </td><td>
  ```go
  func isEmpty(s []string) bool {
    return len(s) == 0
  }
  ```
  </td></tr>
  </tbody></table>
- The zero value (a slice declared with `var`) is usable immediately without
  `make()`.
  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>
  ```go
  nums := []int{}
  // or, nums := make([]int)
  if add1 {
    nums = append(nums, 1)
  }
  if add2 {
    nums = append(nums, 2)
  }
  ```
  </td><td>
  ```go
  var nums []int
  if add1 {
    nums = append(nums, 1)
  }
  if add2 {
    nums = append(nums, 2)
  }
  ```
  </td></tr>
  </tbody></table>
Remember that, while it is a valid slice, a nil slice is not equivalent to an
allocated slice of length 0 - one is nil and the other is not - and the two may
be treated differently in different situations (such as serialization).
### Reduce Scope of Variables
Where possible, reduce scope of variables. Do not reduce the scope if it
conflicts with [Reduce Nesting](#reduce-nesting).
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
err := os.WriteFile(name, data, 0644)
if err != nil {
 return err
}
```
</td><td>
```go
if err := os.WriteFile(name, data, 0644); err != nil {
 return err
}
```
</td></tr>
</tbody></table>
If you need a result of a function call outside of the if, then you should not
try to reduce the scope.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
if data, err := os.ReadFile(name); err == nil {
  err = cfg.Decode(data)
  if err != nil {
    return err
  }
  fmt.Println(cfg)
  return nil
} else {
  return err
}
```
</td><td>
```go
data, err := os.ReadFile(name)
if err != nil {
   return err
}
if err := cfg.Decode(data); err != nil {
  return err
}
fmt.Println(cfg)
return nil
```
</td></tr>
</tbody></table>
### Avoid Naked Parameters
Naked parameters in function calls can hurt readability. Add C-style comments
(`/* ... */`) for parameter names when their meaning is not obvious.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// func printInfo(name string, isLocal, done bool)
printInfo("foo", true, true)
```
</td><td>
```go
// func printInfo(name string, isLocal, done bool)
printInfo("foo", true /* isLocal */, true /* done */)
```
</td></tr>
</tbody></table>
Better yet, replace naked `bool` types with custom types for more readable and
type-safe code. This allows more than just two states (true/false) for that
parameter in the future.
```go
type Region int
const (
  UnknownRegion Region = iota
  Local
)
type Status int
const (
  StatusReady Status = iota + 1
  StatusDone
  // Maybe we will have a StatusInProgress in the future.
)
func printInfo(name string, region Region, status Status)
```
### Use Raw String Literals to Avoid Escaping
Go supports [raw string literals](https://golang.org/ref/spec#raw_string_lit),
which can span multiple lines and include quotes. Use these to avoid
hand-escaped strings which are much harder to read.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
wantError := "unknown name:\"test\""
```
</td><td>
```go
wantError := `unknown error:"test"`
```
</td></tr>
</tbody></table>
### Initializing Structs
#### Use Field Names to Initialize Structs
You should almost always specify field names when initializing structs. This is
now enforced by [`go vet`].
  [`go vet`]: https://golang.org/cmd/vet/
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
k := User{"John", "Doe", true}
```
</td><td>
```go
k := User{
    FirstName: "John",
    LastName: "Doe",
    Admin: true,
}
```
</td></tr>
</tbody></table>
Exception: Field names *may* be omitted in test tables when there are 3 or
fewer fields.
```go
tests := []struct{
  op Operation
  want string
}{
  {Add, "add"},
  {Subtract, "subtract"},
}
```
#### Omit Zero Value Fields in Structs
When initializing structs with field names, omit fields that have zero values
unless they provide meaningful context. Otherwise, let Go set these to zero
values automatically.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
user := User{
  FirstName: "John",
  LastName: "Doe",
  MiddleName: "",
  Admin: false,
}
```
</td><td>
```go
user := User{
  FirstName: "John",
  LastName: "Doe",
}
```
</td></tr>
</tbody></table>
This helps reduce noise for readers by omitting values that are default in
that context. Only meaningful values are specified.
Include zero values where field names provide meaningful context. For example,
test cases in [Test Tables](#test-tables) can benefit from names of fields
even when they are zero-valued.
```go
tests := []struct{
  give string
  want int
}{
  {give: "0", want: 0},
  // ...
}
```
#### Use `var` for Zero Value Structs
When all the fields of a struct are omitted in a declaration, use the `var`
form to declare the struct.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
user := User{}
```
</td><td>
```go
var user User
```
</td></tr>
</tbody></table>
This differentiates zero valued structs from those with non-zero fields
similar to the distinction created for [map initialization], and matches how
we prefer to [declare empty slices][Declaring Empty Slices].
  [map initialization]: #initializing-maps
#### Initializing Struct References
Use `&T{}` instead of `new(T)` when initializing struct references so that it
is consistent with the struct initialization.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
sval := T{Name: "foo"}
// inconsistent
sptr := new(T)
sptr.Name = "bar"
```
</td><td>
```go
sval := T{Name: "foo"}
sptr := &T{Name: "bar"}
```
</td></tr>
</tbody></table>
### Initializing Maps
Prefer `make(..)` for empty maps, and maps populated
programmatically. This makes map initialization visually
distinct from declaration, and it makes it easy to add size
hints later if available.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
var (
  // m1 is safe to read and write;
  // m2 will panic on writes.
  m1 = map[T1]T2{}
  m2 map[T1]T2
)
```
</td><td>
```go
var (
  // m1 is safe to read and write;
  // m2 will panic on writes.
  m1 = make(map[T1]T2)
  m2 map[T1]T2
)
```
</td></tr>
<tr><td>
Declaration and initialization are visually similar.
</td><td>
Declaration and initialization are visually distinct.
</td></tr>
</tbody></table>
Where possible, provide capacity hints when initializing
maps with `make()`. See
[Specifying Map Capacity Hints](#specifying-map-capacity-hints)
for more information.
On the other hand, if the map holds a fixed list of elements,
use map literals to initialize the map.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
m := make(map[T1]T2, 3)
m[k1] = v1
m[k2] = v2
m[k3] = v3
```
</td><td>
```go
m := map[T1]T2{
  k1: v1,
  k2: v2,
  k3: v3,
}
```
</td></tr>
</tbody></table>
The basic rule of thumb is to use map literals when adding a fixed set of
elements at initialization time, otherwise use `make` (and specify a size hint
if available).
### Format Strings outside Printf
If you declare format strings for `Printf`-style functions outside a string
literal, make them `const` values.
This helps `go vet` perform static analysis of the format string.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
msg := "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
```
</td><td>
```go
const msg = "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
```
</td></tr>
</tbody></table>
### Naming Printf-style Functions
When you declare a `Printf`-style function, make sure that `go vet` can detect
it and check the format string.
This means that you should use predefined `Printf`-style function
names if possible. `go vet` will check these by default. See [Printf family]
for more information.
  [Printf family]: https://golang.org/cmd/vet/#hdr-Printf_family
If using the predefined names is not an option, end the name you choose with
f: `Wrapf`, not `Wrap`. `go vet` can be asked to check specific `Printf`-style
names but they must end with f.
```shell
$ go vet -printfuncs=wrapf,statusf
```
See also [go vet: Printf family check].
  [go vet: Printf family check]: https://kuzminva.wordpress.com/2017/11/07/go-vet-printf-family-check/
## Patterns
### Test Tables
Use table-driven tests with [subtests] to avoid duplicating code when the core
test logic is repetitive.
  [subtests]: https://blog.golang.org/subtests
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// func TestSplitHostPort(t *testing.T)
host, port, err := net.SplitHostPort("192.0.2.0:8000")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "8000", port)
host, port, err = net.SplitHostPort("192.0.2.0:http")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "http", port)
host, port, err = net.SplitHostPort(":8000")
require.NoError(t, err)
assert.Equal(t, "", host)
assert.Equal(t, "8000", port)
host, port, err = net.SplitHostPort("1:8")
require.NoError(t, err)
assert.Equal(t, "1", host)
assert.Equal(t, "8", port)
```
</td><td>
```go
// func TestSplitHostPort(t *testing.T)
tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  {
    give:     "192.0.2.0:8000",
    wantHost: "192.0.2.0",
    wantPort: "8000",
  },
  {
    give:     "192.0.2.0:http",
    wantHost: "192.0.2.0",
    wantPort: "http",
  },
  {
    give:     ":8000",
    wantHost: "",
    wantPort: "8000",
  },
  {
    give:     "1:8",
    wantHost: "1",
    wantPort: "8",
  },
}
for _, tt := range tests {
  t.Run(tt.give, func(t *testing.T) {
    host, port, err := net.SplitHostPort(tt.give)
    require.NoError(t, err)
    assert.Equal(t, tt.wantHost, host)
    assert.Equal(t, tt.wantPort, port)
  })
}
```
</td></tr>
</tbody></table>
Test tables make it easier to add context to error messages, reduce duplicate
logic, and add new test cases.
We follow the convention that the slice of structs is referred to as `tests`
and each test case `tt`. Further, we encourage explicating the input and output
values for each test case with `give` and `want` prefixes.
```go
tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  // ...
}
for _, tt := range tests {
  // ...
}
```
Parallel tests, like some specialized loops (for example, those that spawn
goroutines or capture references as part of the loop body),
must take care to explicitly assign loop variables within the loop's scope to
ensure that they hold the expected values.
```go
tests := []struct{
  give string
  // ...
}{
  // ...
}
for _, tt := range tests {
  tt := tt // for t.Parallel
  t.Run(tt.give, func(t *testing.T) {
    t.Parallel()
    // ...
  })
}
```
In the example above, we must declare a `tt` variable scoped to the loop
iteration because of the use of `t.Parallel()` below.
If we do not do that, most or all tests will receive an unexpected value for
`tt`, or a value that changes as they're running.
### Functional Options
Functional options is a pattern in which you declare an opaque `Option` type
that records information in some internal struct. You accept a variadic number
of these options and act upon the full information recorded by the options on
the internal struct.
Use this pattern for optional arguments in constructors and other public APIs
that you foresee needing to expand, especially if you already have three or
more arguments on those functions.
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>
```go
// package db
func Open(
  addr string,
  cache bool,
  logger *zap.Logger
) (*Connection, error) {
  // ...
}
```
</td><td>
```go
// package db
type Option interface {
  // ...
}
func WithCache(c bool) Option {
  // ...
}
func WithLogger(log *zap.Logger) Option {
  // ...
}
// Open creates a connection.
func Open(
  addr string,
  opts ...Option,
) (*Connection, error) {
  // ...
}
```
</td></tr>
<tr><td>
The cache and logger parameters must always be provided, even if the user
wants to use the default.
```go
db.Open(addr, db.DefaultCache, zap.NewNop())
db.Open(addr, db.DefaultCache, log)
db.Open(addr, false /* cache */, zap.NewNop())
db.Open(addr, false /* cache */, log)
```
</td><td>
Options are provided only if needed.
```go
db.Open(addr)
db.Open(addr, db.WithLogger(log))
db.Open(addr, db.WithCache(false))
db.Open(
  addr,
  db.WithCache(false),
  db.WithLogger(log),
)
```
</td></tr>
</tbody></table>
Our suggested way of implementing this pattern is with an `Option` interface
that holds an unexported method, recording options on an unexported `options`
struct.
```go
type options struct {
  cache  bool
  logger *zap.Logger
}
type Option interface {
  apply(*options)
}
type cacheOption bool
func (c cacheOption) apply(opts *options) {
  opts.cache = bool(c)
}
func WithCache(c bool) Option {
  return cacheOption(c)
}
type loggerOption struct {
  Log *zap.Logger
}
func (l loggerOption) apply(opts *options) {
  opts.logger = l.Log
}
func WithLogger(log *zap.Logger) Option {
  return loggerOption{Log: log}
}
// Open creates a connection.
func Open(
  addr string,
  opts ...Option,
) (*Connection, error) {
  options := options{
    cache:  defaultCache,
    logger: zap.NewNop(),
  }
  for _, o := range opts {
    o.apply(&options)
  }
  // ...
}
```
Note that there's a method of implementing this pattern with closures but we
believe that the pattern above provides more flexibility for authors and is
easier to debug and test for users. In particular, it allows options to be
compared against each other in tests and mocks, versus closures where this is
impossible. Further, it lets options implement other interfaces, including
`fmt.Stringer` which allows for user-readable string representations of the
options.
See also,
- [Self-referential functions and the design of options]
- [Functional options for friendly APIs]
  [Self-referential functions and the design of options]: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
  [Functional options for friendly APIs]: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
<!-- TODO: replace this with parameter structs and functional options, when to
use one vs other -->
## Linting
More importantly than any "blessed" set of linters, lint consistently across a
codebase.
We recommend using the following linters at a minimum, because we feel that they
help to catch the most common issues and also establish a high bar for code
quality without being unnecessarily prescriptive:
- [errcheck] to ensure that errors are handled
- [goimports] to format code and manage imports
- [golint] to point out common style mistakes
- [govet] to analyze code for common mistakes
- [staticcheck] to do various static analysis checks
  [errcheck]: https://github.com/kisielk/errcheck
  [goimports]: https://godoc.org/golang.org/x/tools/cmd/goimports
  [golint]: https://github.com/golang/lint
  [govet]: https://golang.org/cmd/vet/
  [staticcheck]: https://staticcheck.io/
### Lint Runners
We recommend [golangci-lint] as the go-to lint runner for Go code, largely due
to its performance in larger codebases and ability to configure and use many
canonical linters at once. This repo has an example [.golangci.yml] config file
with recommended linters and settings.
golangci-lint has [various linters] available for use. The above linters are
recommended as a base set, and we encourage teams to add any additional linters
that make sense for their projects.
  [golangci-lint]: https://github.com/golangci/golangci-lint
  [.golangci.yml]: https://github.com/uber-go/guide/blob/master/.golangci.yml
  [various linters]: https://golangci-lint.run/usage/linters/
