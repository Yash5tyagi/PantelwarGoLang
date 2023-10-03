# Interfaces

### What are Interfaces?
An interface is like a contract that defines a set of methods that a type must implement. It's a way to specify what behavior a certain type should have without specifying how it should be implemented.

#### Key Points

- Interface is a set of method signatures.
- When a type provides an definition for all the methods in the interface,it is said to implement the interface
- Eg: Washing Machine can be an interface with method signatures Cleaning() and Drying()
- Any type which provides definition for cleaning() and Drying() methods is said to implement the Washing machine Interface
- Interface is like a class like thing in which I will create a generic functions jisme whenever I will create a variable of that type then jab function call krunga interface ka then vo automatically detect krke type ussi type ka function chala dega.

### Type Assertion
- In Go, type assertions allow you to check and convert the type of a value from one type to another. Here's a simple explanation:

Imagine you have a box, and you're not sure what's inside it. Type assertions are like opening the box, checking what's inside, and making sure it's what you expect.

Here's how it works:

1. **Checking the Type**: You start by checking if the value inside the box is of a certain type.

2. **Conversion**: If it is indeed of that type, you can convert it to that type so that you can work with it more easily.

Here's some Go code to illustrate:

```go
var unknown interface{} // This is like a box with an unknown content

unknown = "Hello, Go!" // Now, let's say we put a string into the box

// We want to check if the content of the box is a string and then convert it to a string
if str, ok := unknown.(string); ok {
    // Now, 'str' is a string that we can work with
    fmt.Println("It's a string:", str)
} else {
    // The content is not a string
    fmt.Println("It's not a string")
}
```

In this example:

- We start with an `interface{}` type, which is like a box that can hold any type.
- We assign a string to this box.
- We use a type assertion to check if the content of the box is a string, and if it is, we convert it to a string variable called `str`.
- If the type assertion is successful, we print "It's a string" along with the string value. Otherwise, we print "It's not a string."

Type assertions are handy when you're working with interfaces or when you receive data of unknown types, and you need to figure out what's inside and handle it accordingly.

### Type Switch
- A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

