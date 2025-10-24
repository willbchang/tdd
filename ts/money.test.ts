import { expect, test } from 'vitest'

class Dollar {
  readonly amount: number

  constructor(amount: number) {
    this.amount = amount
  }

  times(i: number) {
    return new Dollar(this.amount * i)
  }

}

test('adds 1 + 2 to equal 3', () => {
  const fiver = new Dollar(5)
  const tenner = fiver.times(2)
  expect(tenner.amount).toBe(10)
})