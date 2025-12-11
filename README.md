## gontract
design-by-contract-like pre- and postconditions  for golang

This project provides a golang module to enable writing golang code in a design-by-contract-like fashion.

Lets look at what Design by Contract is and how it applies to golang.


# Design by Contract

Around 1988, [Bertrand Meyer](https://en.wikipedia.org/wiki/Bertrand_Meyer) has introduced the  ["Design by Contract"  methodology](https://en.wikipedia.org/wiki/Design_by_contract) (DbC) for object oriented programming (OOP) along with his programming language Eiffel (see [wikipedia](https://en.wikipedia.org/wiki/Eiffel_%28programming_language%29) and [eiffel.org](https://www.eiffel.org/) ).

The core of DbC consists essentially of three concepts:

*  preconditions
* postconditions
* (class) invariants


# DbC for non-OOP languages?


while class invariants are by their name alone  quite evidently indeed specific to OOP,  the hypothesis that this project aims to support is that at least preconditions and postconditions
can be applied more generally to routines (functions/methods) in non-OOP languages like golang  as well.

## Preconditions and Postconditions

Let's have a closer look at pre- and postconditions.

**preconditions** are conditions that are  typically  imposed  on  the input of a funcion (parameters/arguments) and they  have to be satisfied in order for the function body to  run at all. It is the caller's responsibiliy to make sure that preconditions are satisfied.

**postconditions** are usually imposed on the result of a function  (return values)
and they have to be satisfied in order for the function to complete regularly.
It is the responsibility of the function implementation  to make sure postconditions are satisfied
if the function ran, i. e. if it was provided with valid input.

preconditions and ppostconditions form the essence of the **contract** between the
contractor (the function) and the contractee (the caller of the function).

Pre- and postconditions are essentially means for checking the input and result of a function, just a bit more intrinsically and idiomatically than with the usual conditional (if-else) 
constructs. -- if built into the programming language proper.

It is however an important point to note that preconditions and postconditions serve a very different purpose than usual input and result  validation:

While input validation and result validation are mostly meant to catch runtime errors of a program (like invalid user input), preconditions and postconditions are mostly  meant to catch bugs in the software:

*  preconditions catch bugs in the caller.
*  postconditions catch bugs in the function implementation.
  

summing up, preconditions and postconditions are meant to make it easy and natural to write reliable and corret software, especially when the conditions -- i. e. the contract -- are written before the actual implementation. with other words the contract should be considered an important part if an API rather than something separate or part of the implementation. 



As a conclusion, he above  description should  have made it evident that the concepts of preconditions and postconditions are largely applicable to non-OOP languages.

## Input Validation

The above description poses the question whether preconditions are a replacement for input validation.

This is however certainly not the case:

Input validation is still needed, but the contract moves the responsibility from the function(contractor) to the caller (contractee).

With other words, one kind of software bugs that DbC can reveal is  that a user/caller of a function has not done proper validation of user input.






## Unit Testing

In many projects with procedural (non-OOP) programming languages, unit testing is an important part of the development (test driven development).

Are unit tests not necessary or useful with DbC?

unit tests are certainly useful for DbC projects but one important aspect to bear in mind is that unit tests for functions with pre- and postconditions
have to be written in a different manner than for functions without:

qriting unit tests for functions using conditions requites mechanisms for catching and recovering from violations of conditions.








# This project 

This project offers mechanisms for writing  preconditions and postconditions in golang.

Since golang does not have any such mechanisms   built into the language, 
they are implemented as separate functions intended  to be called at the beginning
and at the end of a function definition, respectively.

gontract implements conditions in an assertion-like manner using `panic()`.

firstly, there is a general condition function:
```go
func Condition(predicate bool, kind Kind, msg string)
```

Here, the type  `Kind`is defined in gontract and can take values  `JindPre` or `KindPost`.

Secongly, there are special-purpose wrappers:

```go
func PreCondition(predicate bool, msg string)


func PostCondition(predicate bool, msg string)
```

The condition functions panic if the `predicate` is false and return normally otherwise.







This approach effectively prevents a function to run or complete at all when conditions are not satisfied.

The implementation of gontract was partly inspired by

[stone.code/assert](https://pkg.go.dev/gitlab.com/stone.code/assert)

A couple of examples -- both positive and negative --  are provided to  illustrate how this library can be used:

* https://github.com/obnoxxx/gontract/tree/main/cmd/example_sqrt_success
* https://github.com/obnoxxx/gontract/tree/main/cmd/example_sqrt_fail_pre
* https://github.com/obnoxxx/gontract/tree/main/cmd/example_sqrt_fail_post 















