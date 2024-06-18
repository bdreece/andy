export default document.cookie
    .split(';')
    .map(v => v.split('='))
    .reduce(
        (cookies, [key, value]) => cookies.set(key, value),
        new Map<string, string>(),
    ) as ReadonlyMap<string, string>;
