# from importlib import import_module
#
# module_a = import_module("a")

module_a = __import__("a")

def bb():
    module_a.aa()
    pass

if __name__ == '__main__':
    bb()

