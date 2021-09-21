function IsPronic(x){
    const sqrt = Math.sqrt(x);
    const intBelow = Math.trunc(sqrt);
    const intAbove = Math.ceil(sqrt);
    return x === (intAbove * intBelow);
}