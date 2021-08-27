import N0d3 from './n0d3'

interface SearchResult{
    key: number|string,
    content: any
}

export default class Bee3{
    root: N0d3|null;
    

    constructor(key: number|string, content){
        this.root = new N0d3([key], [content])
    }

    search(x:N0d3, key:number|string): null|SearchResult{
        let i = 0;
        const height = x.n;
        while(i <= height && key > x.keys[i]){
            i++;
        }
        if(i <= height && key === x.keys[i]){
            return {key: x.keys[i], content:x.content[i]}
        }else if(x.isLeaf){
            return null
        }
        return this.search(x.content[i], key);
    }
}
