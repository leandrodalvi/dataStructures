
class AVLNode{
    constructor(key, value, height){
        this.key = key;
        this.value = value;
        this.right = null;
        this.left = null;
        this.height = height;
    }

    //getters

    getKey(){
        return this.key;
    }

    getValue(){
        return this.value;
    }

    getRight(){
        return this.right;
    }

    getLeft(){
        return this.left;
    }

    getHeight(){
        return this.height;
    }

    //setters

    setKey(value){
        this.key = value;
    }

    setValue(value){
        this.value = value;
    }

    setRight(value){
        this.right = value;
    }

    setLeft(value){
        this.left = value;
    }

    setHeight(value){
        this.height = value;
    }

}