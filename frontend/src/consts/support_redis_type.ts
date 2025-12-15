export const types: Record<string, string> = {
    STRING: 'STRING',
    HASH: 'HASH',
    LIST: 'LIST',
    SET: 'SET',
    ZSET: 'ZSET',
} as const

export const typesColor: Record<string, string> = {
    [types.STRING]: '#5A96E3',
    [types.HASH]: '#9575DE',
    [types.LIST]: '#7A9D54',
    [types.SET]: '#F3AA60',
    [types.ZSET]: '#FF6666',
}

// export const typesName = Object.fromEntries(Object.entries(types).map(([key, value]) => [key, value.name]))

export const validType = (t: string):boolean => {
    return t in types
}