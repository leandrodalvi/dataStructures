
function floodFill(pixel = [0,0], newC, originalChar){
    let x = pixel[0];
    const y = pixel[1]
    if (x < 0 || x >= numeroDeLinhas || y < 0 || y >= numeroDeColunas){
        return;
    }

    if(matrix[x][y] !== originalChar){
        return
    }else{
        matrix[pixel[0]][pixel[1]] = newC;

        floodFill([x+1, y], originalChar, newC);
        floodFill([x-1, y], originalChar, newC);
        floodFill([x, y+1], originalChar, newC);
        floodFill([x, y-1], originalChar, newC);
    }


}