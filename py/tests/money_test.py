class Dollar:
    def __init__(self, amount: int):
        self.amount = amount

    def times(self, multiple: int):
        return Dollar(self.amount * multiple)


def test_money():
    fiver = Dollar(5)
    tenner = fiver.times(2)
    assert tenner.amount == 10
