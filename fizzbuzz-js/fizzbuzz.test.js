const { fizzbuzz } = require("./fizzbuzz");

test("returns number for non-divisible", () => {
  expect(fizzbuzz(1)).toBe("1");
  expect(fizzbuzz(7)).toBe("7");
});

test("returns Fizz for multiples of 3", () => {
  expect(fizzbuzz(3)).toBe("Fizz");
  expect(fizzbuzz(9)).toBe("Fizz");
});

test("returns Buzz for multiples of 5", () => {
  expect(fizzbuzz(5)).toBe("Buzz");
  expect(fizzbuzz(10)).toBe("Buzz");
});

test("returns FizzBuzz for multiples of 15", () => {
  expect(fizzbuzz(15)).toBe("FizzBuzz");
  expect(fizzbuzz(30)).toBe("FizzBuzz");
});
