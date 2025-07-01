class File:
    def __init__(self, name, size):
        self.name = name
        self.size = size
        self.children = []
        self.is_directory = False if '.' in name else True
        self.extension = name.split('.')[-1] if '.' in name else ""

    def __repr__(self):
        return f"{self.name}"

class filter():
    def __init__(self):
        pass
    def apply(self, file: File):
        pass

class minSizeFilter(filter):
    def __init__(self, size: int):
        self.size = size

    def apply(self, file: File):
        return file.size > self.size
    
class extensionFilter(filter):
    def __init__(self, extension: str):
        self.extension = extension

    def apply(self, file: File):
        return self.extension == file.extension
    

class FileSystem:
    def __init__(self):
        self.filters = []

    def add_filter(self, new_filter: filter):
        if isinstance(new_filter, filter):
            self.filters.append(new_filter)

    def apply_or_filtering(self, root):
        

        def dfs(root, result):
            if root.isDirectory:
                for child in root.children:
                    dfs(child, result)
            else:
                for filter in self.filters:
                    if filter.apply(root):
                        result.append(root)
                        print(result)
                        return
        result = []
        dfs(root, result)
        return result
    
    def apply_and_filtering(self, root):

        def dfs(root, result):
            if root.isDirectory:
                for child in root.children:
                    dfs(child, result)
            else:
                for filter in self.filters:
                    if not filter.apply(root):
                        return
                result.append(root)
                print(result)
        
        result = []
        dfs(root, result)
        return result


    

"""
   f1
 /  / \
f2 f3 f4
//\
f5 f6 f7
"""
if __name__ == "__main__":

    ## construct file tree
    f1 = File("root_300", 300)

    f2 = File("fiction_100", 100)
    f3 = File("action_200", 200)
    f4 = File("comedy_150", 150)

    f1.children = [f2, f3, f4]

    f5 = File("StarTrek_100.mov", 100)
    f6 = File("StarWars_100.txt", 100)
    f7 = File("Avater_200.mp4", 200)
    
    f2.children = [f5, f6, f7]

    f8 = File("SNL.mp4", 111)
    f9 = File("SNL.txt", 111)
    f4.children = [f8]

    ## add filters
    greater_than_100 = minSizeFilter(100)
    mp4_filter = extensionFilter("mp4")

    my_file_system = FileSystem()
    my_file_system.add_filter(greater_than_100)
    my_file_system.add_filter(mp4_filter)

    ## apply filters
    result = my_file_system.apply_or_filtering(f1)
    print(result)

    ## apply and filtering
    result = my_file_system.apply_and_filtering(f1)
    print(result)
    
    
    

