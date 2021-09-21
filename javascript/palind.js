function palin(word){
    for(let i = 1; i < word.length; i++){
        if(word[i - 1] !== word[word.length - i])
            return false;
    }
    return true;
}
