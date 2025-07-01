# Follow-up: How to handle additional items like Coke and Wings in the order system

# Step 1: Define a Generic Interface
from abc import ABC, abstractmethod

class OrderItem(ABC):
    @abstractmethod
    def get_description(self):
        pass
    @abstractmethod
    def get_cost(self):
        pass

# Step 2: Implement Pizza & Topping Using Decorator Pattern
class BasePizza(OrderItem):
    def get_description(self):
        return "Plain Pizza"
    def get_cost(self):
        return 5.0
    
class ToppingDecorator(OrderItem):
    def __init__(self, pizza: OrderItem):
        self.pizza = pizza

class Cheese(ToppingDecorator):
    def get_description(self):
        return self.pizza.get_description() + ", Cheese"
    
    def get_cost(self):
        return self.pizza.get_cost() + 1.0

# add more topping classes similarly

# Step 3: Implement Coke and Wings as Standalone Items
class Coke(OrderItem):
    def get_description(self):
        return "Coke"
    
    def get_cost(self):
        return 2.0

class Wings(OrderItem):
    def get_description(self):
        return "Wings"
    
    def get_cost(self):
        return 4.0
    
# Step 4: Composite Order Class
class Order:
    def __init__(self):
        self.items = []

    def add_item(self, item: OrderItem):
        self.items.append(item)

    def get_total_cost(self):
        res = 0
        for item in self.items:
            res += item.get_cost()
        return res
    
    def get_order_description(self):
        return "; ".join(item.get_description() for item in self.items)
    
if __name__ == "__main__":
    # Usage Example
    # Build items
    plain_pizza = BasePizza()
    cheese_pizza = Cheese(plain_pizza)
    coke = Coke()
    wings = Wings()

    # Build the order
    order = Order()
    order.add_item(cheese_pizza)
    order.add_item(coke)
    order.add_item(wings)

    # Print out
    print(order.get_order_description())  # Plain Pizza, Cheese; Coke; Wings
    print(order.get_total_cost())         # 5.0 + 1.0 + 2.0 + 4.0 = 12.0
