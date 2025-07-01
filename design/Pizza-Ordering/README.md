# Pizza Ordering System

## Problem
Design a customizable pizza ordering system using the Decorator pattern.

## Concepts 
Design pattern (Decorator), composition over inheritance

## UML
```
             +-------------------+
             |     <<interface>> |
             |       Pizza       |
             +-------------------+
             | +get_description()|
             | +get_cost()       |
             +-------------------+
               /                 \
              /                   \
   +--------------------+          +--------------------+
   |    BasePizza       |          | ToppingDecorator   |
   +--------------------+          +--------------------+
   | +get_description() | <-has a- | -pizza: Pizza      |
   | +get_cost()        |          | +__init__(Pizza)   |
   +--------------------+          +--------------------+
                                           |
                                           |
                                  +--------------------+
                                  |      Cheese        |
                                  +--------------------+
                                  | +get_description() |
                                  | +get_cost()        |
                                  +--------------------+

```

## Follow-ups
1. How to handle additional items like Coke and Wings in the order system
2. How to implement a coupon system for the entire order
3. How to implement a buy-one-get-one-free coupon system
4. How to handle free toppings and where to place this logic
5. How to handle different store configurations with varying prices

## Resource
- [pizza OOD](https://www.1point3acres.com/bbs/thread-1107792-1-1.html)