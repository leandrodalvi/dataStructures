function binarySearch(name ,array, min = 0, max = array.length - 1){
    let middle = Math.floor((min + max)/2);
   
    if (array[middle].name ===name){
        return array[middle].phone
    }else{
        if(array[middle].name > name) 
            return binarySearch(name, array, min, middle-1);
        else
            return binarySearch(name, array, middle+1, max);
    }
}