This is code example of Strategy pattern implementation with Go.

> The Strategy Pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

In this example, different ways of using a weapon, form our family of algorithms


_Read https://medium.com/@vansikagupta247/design-pattern-with-go-ft-strategy-pattern-a1efc58972e6 for explanation of the example._

*Key take away points:*
* Seperate what remains same from what keeps varying.
    Varying behavior in this example was functioning of a weapon. They were designed for the same purpose/behavior/action but with different implemenations or ways to function.
    Seperate it out into it's own package.
* Composition.
    Vigilante holds a reference to weapons. Weapon are interchangeable at run-time
