# Types in Go

Roughly, we have four different groups of useful primitive types, namely: **string, numeric type, bool, and error type**.
They are all made available by *builtin* package.

- **string**:
    - is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text
    - may be empty, but not nil. 
    - are immutable (like Java)
- **bool**: *true* or *false*
- numeric:
    - integers:
        - unsigned: **uint8** (unsigned 8-bit integers), **uint16**, **uint32**, **uint64** 
        - signed: **int8** (signed 8-bit integers), **int16**, **int32**, **int64**,
    - real (floating point):
        - **float32** (IEEE-754 32-bit floating-point numbers), **float64** 
    - complex numbers:
        - **complex64** (complex numbers with float32 real and imaginary parts)
        - **complex128** (the set of all complex numbers with float64 real and imaginary parts)
    - aliases:
        - **byte**: alias for uint8 
        - **rune**: alias for int32

- Other information:
    - Although **error** is a primitive type, it is an interface type, which wrappers around the string type.
    - Variable declarations may be "factored" into blocks, as with import statements.
    - The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. 
    - **When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.**

