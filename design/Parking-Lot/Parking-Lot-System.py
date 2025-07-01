from enum import Enum

class VehicleSize(Enum):
    MOTORCYCLE = 1
    COMPACT = 2
    LARGE = 3

class Vehicle:
    def __init__(self, plate, size: VehicleSize):
        self.plate = plate
        self.size = size

class ParkingSpot:
    def __init__(self, size: VehicleSize):
        self.size = size
        self.vehicle = None

    def can_park(self, vehicle):
        return self.vehicle is None and self.size >= vehicle.size
    
    def park(self, vehicle):
        if self.can_park(vehicle):
            self.vehicle = vehicle
            return True
        return False
    
class ParkingLot:
    def __init__(self, capacity: int, spots: list[ParkingSpot]):
        self.capacity = capacity
        self.spots = spots

    def park_vehicle(self, vehicle):
        for spot in self.spots:
            if spot.can_park(vehicle):
                spot.park(vehicle)
                return spot
            
        return None # Lot is full
    
    def unpark_vehicle(self, plate):
        for spot in self.spots:
            if spot.vehicle and spot.vehicle.plate == plate:
                spot.vehicle = None
                return True
            
        return False # Vehicle not found

if __name__ == "__main__":
    # This would work
    large_spot = ParkingSpot(VehicleSize.LARGE)
    motorcycle = Vehicle("ABC123", VehicleSize.MOTORCYCLE)
    large_spot.can_park(motorcycle)  # True (3 >= 1)

    # This would fail
    motorcycle_spot = ParkingSpot(VehicleSize.MOTORCYCLE) 
    large_car = Vehicle("XYZ789", VehicleSize.LARGE)
    motorcycle_spot.can_park(large_car)  # False (1 >= 3 is False)