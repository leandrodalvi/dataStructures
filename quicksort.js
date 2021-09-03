export default function quicksort(A = [], p = 0, r = A.length - 1){
    if(p<r){
        const q = partition(A, p, r);
        quicksort(A, p, q - 1);
        quicksort(A, q + 1, r);
    }else{
        return;
    }
}


function partition(A, p, r){
    const x = A[r];
    let i = p - 1;
    for(let j = p; j <= r - 1; j++){
        if(A[j] <= x){
            i++;
            const aux = A[j];
            A[j] = A[i];
            A[i] = aux;
        }
    }
    const aux = A[r];
    A[r] = A[i + 1];
    A[i + 1] = aux;
    return i + 1;
}
