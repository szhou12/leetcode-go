class Locker:
    def __init__(self, locker_id, size):
        self.locker_id = locker_id
        self.size = size
        self.occupied = False


class LockerSystem:
    def __init__(self):
        self.lockers = []

    def add_locker(self, locker: Locker):
        self.lockers.append(locker)
    
    def assign_package(self, size):
        for locker in self.lockers:
            if not locker.occupied and locker.size >= size:
                locker.occupied = True
                return locker.locker_id
            
        return None