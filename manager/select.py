import typing
class Select():
    def __init__(self, include:typing.List[str], exclude=[]):
        self.include = include
        self.exclude = exclude

class Task():
    def __init__(self,root:str,language:str, include:typing.List[str], exclude=[]):
        self.root = root
        self.language = language
        self.select = Select(include, exclude)
        

