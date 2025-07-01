"""
He expected basic functionalities like Pizza Base,Pizza Size and Pizza Toppings. 

Started explaining my approach and then started coding it out. 

After creating the main object class, he told me to add Beverage options and how will I modify the code. 

Told I will be adding new classes with different beverage options,sizes and started coding and modified the code. 

After this was told to add Discount and Coupons with a little variation like discount for bases, different toppings, etc. 

Told my approach and accordingly modified the code. In certain places just wrote the placeholder function and explained what I will do and didn't code fully. 
"""

# I have an interface called Pizza, anything that can be described as a pizza and has a cost
class Pizza:
    def get_description(self):
        return ""
    
    def get_cost(self):
        return 0

# BasePizza is a concrete implementation of Pizza
class BasePizza(Pizza):
    def get_description(self):
        return "Plain Pizza"
    
    def get_cost(self):
        return 5.0

# ToppingDecorator is a base (abstract) decorator class
# - It inherits from Pizza so it can be used in place of a Pizza
# - It holds a reference to another Pizza object (self.pizza), allowing wrapping.
# - It typically does not implement get_description() or get_cost() itself, but sets up the structure for decorators.
class ToppingDecorator(Pizza):
    def __init__(self, pizza: Pizza):
        self.pizza = pizza # <--- composition over inheritance!

# Cheese is a concrete decorator
# - It inherits from ToppingDecorator.
# - 
class Cheese(ToppingDecorator):
    def get_description(self):
        return self.pizza.get_description() + ", Cheese"
    
    def get_cost(self):
        return self.pizza.get_cost() + 1.0


if __name__ == "__main__":
    # Usage Example 
    # The beauty of this pattern is its composability
    # This design scales well for complex customization scenarios while maintaining clean, readable code and following SOLID principles.

    # Create a pizza with multiple toppings
    pizza = BasePizza()
    pizza = Cheese(pizza)
    # pizza = Pepperoni(pizza)
    # pizza = Mushroom(pizza)

    print(f"{pizza.get_description()} - ${pizza.get_cost():.2f}")
    # Output: "Plain Pizza, Cheese, Pepperoni, Mushroom - $8.50"