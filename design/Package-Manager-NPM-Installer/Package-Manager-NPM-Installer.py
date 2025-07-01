from collections import defaultdict
class PackageManager:
    def __init__(self):
        self.graph = defaultdict(list) # Maps package to its dependencies
        self.installed = set() # Maps package to its dependencies
    
    def add_dependency(self, package: str, dependency: str):
        self.graph[package].append(dependency)

    def install(self, package: str, visited: set = None):
        if visited is None:
            visited = set()

        if package in visited:
            raise Exception("Cycle detected: Cannot install due to circular dependency")
        
        if package in self.installed:
            return
        
        visited.add(package)

        for dep in self.graph[package]:
            self.install(dep, visited)

        visited.remove(package)
        
        self.installed.add(package)
        print(f"Installed: {package}")

if __name__ == "__main__":
    manager = PackageManager()
    manager.add_dependency("A", "B")
    manager.add_dependency("B", "C")
    manager.add_dependency("C", "A")
    manager.install("A")