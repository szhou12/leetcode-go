class Elevator:
    def __init__(self, id):
        self.id = id
        self.current_floor = 0
        self.direction = 0 # -1: down, 0: idle, 1: up
        self.requests = []

    def request(self, floor):
        self.requests.append(floor)
        self.requests.sort()

    def step(self):
        if self.requests:
            next_floor = self.requests[0]
            if next_floor > self.current_floor:
                self.current_floor += 1
            elif next_floor < self.current_floor:
                self.current_floor -= 1
            else:
                print(f"Stopping at floor {self.current_floor}")
                self.requests.pop(0)
    
