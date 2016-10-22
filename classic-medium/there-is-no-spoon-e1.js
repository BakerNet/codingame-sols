/**
 * Don't let the machines win. You are humanity's last hope...
 **/

var width = parseInt(readline()); // the number of cells on the X axis
var height = parseInt(readline()); // the number of cells on the Y axis
var grid = new Array(width);
for(let i = 0; i < width; i++){
    grid[i] = new Array(height);
}
for (var i = 0; i < height; i++) {
    var line = readline(); // width characters, each either 0 or .
    //print(line);
    line = line.split('');
    //print(line);
    for (let j=0; j<width; j++){
        if (line[j] === '0'){
            grid[j][i] = 1;
        }else{
            grid[j][i] = -1;
        }
    }
}

var string = '';

for (let i = 0; i< height; i++){
    for (let j = 0; j<width; j++){
        string = '';
        if(grid[j][i] > 0){
            string += j + ' ' + i + ' ';
            let rflag = false;
            for(let k = j + 1; k < width; k++){
                if(grid[k][i] > 0){
                    string += k + ' ' + i + ' ';
                    rflag = true;
                    break;
                }
            }
            if(!rflag){
                string += '-1 -1 ';
            }
            let lflag = false;
            for(let k = i + 1; k < height; k++){
                if(grid[j][k] > 0){
                    string += j + ' ' + k + ' ';
                    lflag = true;
                    break;
                }
            }
            if(!lflag){
                string += '-1 -1 ';
            }
            
            print(string);
            
        }
    }
}