/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var inputs = readline().split(' ');
var L = parseInt(inputs[0]);
var C = parseInt(inputs[1]);

var grid = new Array(L);
for (let i = 0; i < L; i++){
    grid[i] = new Array(C);
}

var currstate = [-1, -1, 'S'];
var teleporters = [];
var log = new Array();
Array.prototype.compare = function (array){
    //printErr(array +  " = " + this);
    for (let i = 0; i < this.length; i++){
        if (array[i] != this[i]){
            return false;
        }
    }
    return true;
}

var moves = [];

for (let i = 0; i < L; i++) {
    let row = readline();
    printErr(row);
    grid[i] = row.split('');
    for (let j = 0; j < C; j++){
        if (grid[i][j] === '@'){
            currstate[0] = i;
            currstate[1] = j;
            printErr(currstate);
            grid[i][j] = ' ';
        } else if (grid[i][j] === 'T') {
            teleporters.push([i, j]);
        }
    }
}
printErr(teleporters[0]);
printErr(teleporters[1]);

var inverter = 0;
var breaker = 0;

function moveOne(){
    switch(currstate[2]){
        case 'S':
            currstate[0] = currstate[0] + 1;
            moves.push('SOUTH');
            break;
        case 'E':
            currstate[1] = currstate[1] + 1;
            moves.push('EAST');
            break;
        case 'N':
            currstate[0] = currstate[0] - 1;
            moves.push('NORTH');
            break;
        case 'W':
            currstate[1] = currstate[1] - 1;
            moves.push('WEST');
            break;
        default: print("Error, invalid direction");
            break;
    }
}
function breakOne(){
    printErr("Break..............");
    switch(currstate[2]){
        case 'S':
            currstate[0] = currstate[0] + 1;
            printErr(grid[currstate[0]][currstate[1]]);
            grid[currstate[0]][currstate[1]] = ' ';
            moves.push('SOUTH');
            break;
        case 'E':
            currstate[1] = currstate[1] + 1;
            printErr(grid[currstate[0]][currstate[1]]);
            grid[currstate[0]][currstate[1]] = ' ';
            moves.push('EAST');
            break;
        case 'N':
            currstate[0] = currstate[0] - 1;
            printErr(grid[currstate[0]][currstate[1]]);
            grid[currstate[0]][currstate[1]] = ' ';
            moves.push('NORTH');
            break;
        case 'W':
            currstate[1] = currstate[1] - 1;
            printErr(grid[currstate[0]][currstate[1]]);
            grid[currstate[0]][currstate[1]] = ' ';
            moves.push('WEST');
            break;
        default: print("Error, invalid direction");
            break;
    }
}
function changeDir(){
    let i = currstate[0],
        j = currstate[1];
    switch(inverter){
        case 0:
            if (grid[i + 1][j] !== 'X' && grid[i + 1][j] !== '#'){
                currstate[2] = 'S';
                //print('SOUTH');
            }else if(grid[i][j + 1] !== 'X' && grid[i][j + 1] !== '#'){
                currstate[2] = 'E';
                //print('EAST');
            }else if(grid[i - 1][j] !== 'X' && grid[i - 1][j] !== '#'){
                currstate[2] = 'N';
                //print('NORTH');
            }else if(grid[i][j - 1] !== 'X' && grid[i][j - 1] !== '#'){
                currstate[2] = 'W';
                //print('WEST');
            }else{
                print('LOOP');
                currstate[0] = -1;
                currstate[1] = -1;
            }
            break;
        case 1:
            if(grid[i][j - 1] !== 'X' && grid[i][j - 1] !== '#'){
                currstate[2] = 'W';
                //print('WEST');
            }else if(grid[i - 1][j] !== 'X' && grid[i - 1][j] !== '#'){
                currstate[2] = 'N';
                //print('NORTH');
            }else if(grid[i][j + 1] !== 'X' && grid[i][j + 1] !== '#'){
                currstate[2] = 'E';
                //print('EAST');
            }else if (grid[i + 1][j] !== 'X' && grid[i + 1][j] !== '#'){
                currstate[2] = 'S';
                //print('SOUTH');
            }else{
                print('LOOP');
                currstate[0] = -1;
                currstate[1] = -1;
            }
            break;
        default: print("Error invalid inverter");
            break;
    }
}
function teleport(nexti, nextj){
    if (teleporters[0][0] === currstate[0] && teleporters[0][1] === currstate[1]){
        currstate[0] = teleporters[1][0];
        currstate[1] = teleporters[1][1];
    }else{
        currstate[0] = teleporters[0][0];
        currstate[1] = teleporters[0][1];
    }
}


while((currstate [0] !== -1) && (currstate[1] !== -1)){
    let broke = false;
    let nexti = currstate[0], 
        nextj = currstate[1];
    switch(currstate[2]){
        case 'S':
            nexti = currstate[0] + 1;
            break;
        case 'E':
            nextj = currstate[1] + 1;
            break;
        case 'N':
            nexti = currstate[0] - 1;
            break;
        case 'W':
            nextj = currstate[1] - 1;
            break;
        default: print("Error, invalid direction");
            break;
    }
    let next = grid[nexti][nextj];
    switch(next){
        case '#': 
            changeDir();
            break;
        case 'X':
            if(breaker){
                breakOne();
                log = [];
                broke = true;
            } 
            else changeDir();
            break;
        case 'B':
            breaker = breaker^1;
            moveOne();
            log = [];
            broke = true;
            break;
        case '$':
            moveOne();
            currstate[0] = -1;
            currstate[1] = -1;
            break;
        case 'I':
            inverter = inverter^1;
            moveOne();
            break;
        case 'T':
            moveOne();
            teleport();
            break;
        case 'S':
        case 'W':
        case 'N':
        case 'E':
            moveOne();
            currstate[2] = next;
            break;
        default: 
            moveOne();
            break;
    }
    for (let i = 0; i < log.length; i++){
        //printErr(log);
        if (log[i].compare(currstate)){
            print("LOOP");
            currstate[0] = -1;
            currstate[1] = -1;
            moves = []
        }
    }
    if (!broke){
        log.push(currstate.slice());
    }else{
        broke = false;
    }
}

// Write an action using print()
// To debug: printErr('Debug messages...');
for (let i = 0; i < moves.length; i++){
    print(moves[i]);
}