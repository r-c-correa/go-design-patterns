<p align="center">
  <h1 align="center">
    Abstract Factory
  </h1>
</p>

## Concept
This design pattern lets you create **families** of related objects without specifying their concrete class.

We'll use clothes as an example, in the **family** of clothes we can have pants, shirt blouse, among others. Each outfit has its **variations**, a shirt can be size X, Y or Z have a color A, B or C, among other characteristics.

Now imagine two clothing factories, both of which make shirts and pants, but the way each makes them is different (the implementation), one makes a blue shirt and the other a green shirt. **Abstract Factory** provides **Make Shirt** and **Make Pants** functions, and **Concrete Factories** implement these functions in their own way.

The client who uses the **Abstract Factory** does not need to know the implementation, he just uses, through the **Abstract Factory**, the **Make Shirt** and **Make Pants** functions.

![Abstract Factory Concept](/imgs/abstract_factory_1.jpg?raw=true)

## Implementation
We created an interface that represents the **Abstract Factory**, called **GUIFactory** (Ghaphical User Interface Factory) to represent functions for creating visual components, having **CreateButton** and **CreateListbox** functions. We made two **Concrete Factory** implementations, one for **Windows** and one for **Linux** (as each one assembles the layout with its own style), and we have an object called **Application** that uses the two **GUIFactory** functions from the **CreateUI** function.

![Abstract Factory Example](/imgs/abstract_factory_2.jpg?raw=true)

## Executing
```
go run .\main.go abstract_factory
```