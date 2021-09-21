const fs = require('fs');

export default function unorderedGen(){
    const tenThousandArr = [];
    const hundredThousandArr = [];
    const miArr = [];
    const tenMiArr = []
    for (let index = 0; index < 10000; index++) {
        const aux = Math.floor(Math.random() * 10000)
        tenThousandArr.push(aux);
    }
    for (let index = 0; index < 100000; index++) {
        const aux = Math.floor(Math.random() * 100000)
        hundredThousandArr.push(aux);
    }
    for (let index = 0; index < 1000000; index++) {
        const aux = Math.floor(Math.random() * 1000000)
        miArr.push(aux);
    }
    for (let index = 0; index < 10000000; index++) {
        const aux = Math.floor(Math.random() * 10000000)
        tenMiArr.push(aux);
    }
    const jsonFile = {
        "tenThousand":JSON.stringify(tenThousandArr),
        "hundredThousand": JSON.stringify(hundredThousandArr),
        "mi": JSON.stringify(miArr),
        "tenMi": JSON.stringify(tenMiArr)
    }
    fs.writeFile('samples.json', jsonFile, function (err) {
        if (err) throw err;
            console.log('Saved!');
    });
}
