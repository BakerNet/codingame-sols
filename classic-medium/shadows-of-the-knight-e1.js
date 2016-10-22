/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var inputs = readline().split(' ');
var W = parseInt(inputs[0]); // width of the building.
var H = parseInt(inputs[1]); // height of the building.
var grid = new Array(W);

var N = parseInt(readline()); // maximum number of turns before game over.

var inputs = readline().split(' ');
var X0 = parseInt(inputs[0]);
var Y0 = parseInt(inputs[1]);

var pos = {
    x: X0,
    y: Y0
}

var maxX = W - 1;
var minX = 0;
var maxY = H - 1;
var minY = 0;

// game loop
while (true) {
    var bombDir = readline(); // the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
    // Write an action using print()
    // To debug: printErr('Debug messages...');
    //print(bombDir);
    let move = '';
    switch(bombDir){
        case "U":
            maxY = pos.y - 1;
            pos.y = Math.floor((pos.y + minY) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        case "UR":
            maxY = pos.y - 1;
            pos.y = Math.floor((pos.y + minY) / 2);
            minX = pos.x + 1;
            pos.x = Math.ceil((pos.x + maxX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
            
        case "UL":
            maxY = pos.y - 1;
            pos.y = Math.floor((pos.y + minY) / 2);
            maxX = pos.x - 1;
            pos.x = Math.floor((pos.x + minX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        
        case "R":
            minX = pos.x + 1;
            pos.x = Math.ceil((pos.x + maxX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        
        case "L":
            maxX = pos.x - 1;
            pos.x = Math.floor((pos.x + minX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        case "D":
            minY = pos.y + 1;
            pos.y = Math.ceil((pos.y + maxY) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        case "DR":
            minY = pos.y + 1;
            pos.y = Math.ceil((pos.y + maxY) / 2);
            minX = pos.x + 1;
            pos.x = Math.ceil((pos.x + maxX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
        case "DL":
            minY = pos.y + 1;
            pos.y = Math.ceil((pos.y + maxY) / 2);
            maxX = pos.x - 1;
            pos.x = Math.floor((pos.x + minX) / 2);
            move += pos.x + ' ' + pos.y;
            break;
    }

    // the location of the next window Batman should jump to.
    print(move);
}