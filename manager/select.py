class Select():
    def __init__(self, include:list[str], exclude=[]):
        self.include = include
        self.exclude = exclude

class Task():
    def __init__(self,root:str,language:str, include:list[str], exclude=[]):
        self.root = root
        self.language = language
        self.select = Select(include, exclude)
        

