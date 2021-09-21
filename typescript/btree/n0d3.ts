export default class N0d3 {
    keys: any[];
    content: any[];
    isLeaf: boolean;
    n: number;

    constructor(keys: any[], content: any[], isLeaf: boolean = true, n: number = 0){
        for(let i = 0; i<=keys.length; i++){
            if(typeof keys[i] !== 'string' && typeof keys[i] !== 'number'){
                throw new Error('Every key must be a number or a string')
            }
        }
        this.keys = keys;
        this.content = content;
        this.isLeaf = isLeaf;
        this.n = n;
    }

    setKey(keys: any[]): void{
        for(let i = 0; i<=keys.length; i++){
            if(typeof keys[i] !== 'string' && typeof keys[i] !== 'number'){
                throw new Error('Every key must be a number or a string')
            }
        }
        this.keys = keys;
    }

    setContent(content: any[]): void{
        this.content = content;
    }

    setIsLeaf(isLeaf: boolean): void{
        this.isLeaf = isLeaf;
    }

    setN(n: number): void{
        this.n = n;
    }

}
