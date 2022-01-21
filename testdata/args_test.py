
def test(*args):
    print(",".join(args))

if __name__ == '__main__':
    test(*("1","2","3","4"))
    test(("1","2","3","4"))