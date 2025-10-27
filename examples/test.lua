-- Example Lua programs to test the interpreter

-- Simple arithmetic
local x = 5 + 3 * 2
local y = (5 + 3) * 2

-- Variables and assignments
local name = "Lua"
local version = 5.4
local isAwesome = true

-- Functions
function add(a, b)
    return a + b
end

function factorial(n)
    if n <= 1 then
        return 1
    else
        return n * factorial(n - 1)
    end
end

-- Tables
local person = {
    name = "Alice",
    age = 30,
    city = "Wonderland"
}

local numbers = {1, 2, 3, 4, 5}

-- Control structures
if x > 10 then
    print("x is greater than 10")
elseif x > 5 then
    print("x is greater than 5 but less than or equal to 10")
else
    print("x is 5 or less")
end

-- Loops
for i = 1, 10 do
    print(i)
end

local i = 1
while i <= 10 do
    print(i)
    i = i + 1
end

repeat
    print(i)
    i = i + 1
until i > 20

-- String operations
local greeting = "Hello, " .. name .. "!"
local length = #greeting

-- Table operations
local tableLength = #numbers

-- Logical operators
local result = (x > 0) and (y < 100) or false

-- Closures
function makeCounter()
    local count = 0
    return function()
        count = count + 1
        return count
    end
end

local counter = makeCounter()
print(counter())  -- 1
print(counter())  -- 2
print(counter())  -- 3

-- Multiple return values
function minmax(a, b)
    if a < b then
        return a, b
    else
        return b, a
    end
end

local min, max = minmax(10, 5)

-- Variadic functions
function sum(...)
    local total = 0
    for _, v in ipairs({...}) do
        total = total + v
    end
    return total
end

print(sum(1, 2, 3, 4, 5))  -- 15

-- Table constructor patterns
local matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
}

local mixed = {
    10,
    20,
    name = "Mixed",
    [5] = "five",
    ["key"] = "value"
}

-- Function as first-class values
local operations = {
    add = function(a, b) return a + b end,
    subtract = function(a, b) return a - b end,
    multiply = function(a, b) return a * b end,
    divide = function(a, b) return a / b end
}

print(operations.add(5, 3))  -- 8

-- Comments
-- This is a single-line comment

--[[
    This is a
    multi-line
    comment
]]

-- Metatables (advanced feature for later)
local mt = {
    __add = function(a, b)
        return {value = a.value + b.value}
    end
}

local obj1 = {value = 10}
local obj2 = {value = 20}
setmetatable(obj1, mt)
setmetatable(obj2, mt)

-- local result = obj1 + obj2  -- Uses __add metamethod