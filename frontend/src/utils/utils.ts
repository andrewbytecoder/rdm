let foo: string = "data"

foo.split(' ')

// 如果函数有return ts 会自动推到函数需要一个返回值
// 如果函数有多个返回值，那么TS会自动将返回值类型设置为 string | number, 具体视情况而定
export function bar() :string{
    if (foo === 'data') {
        return 1 + ""
    }

    return foo
}

