/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var inputs = readline().split(' ');
var nbFloors = parseInt(inputs[0]); // number of floors
var width = parseInt(inputs[1]); // width of the area
var nbRounds = parseInt(inputs[2]); // maximum number of rounds
var exitFloor = parseInt(inputs[3]); // floor on which the exit is found
var exitPos = parseInt(inputs[4]); // position of the exit on its floor
var nbTotalClones = parseInt(inputs[5]); // number of generated clones
var nbAdditionalElevators = parseInt(inputs[6]); // ignore (always zero)
var nbElevators = parseInt(inputs[7]); // number of elevators
var elevators = new Array(nbFloors);
for(let i = 0; i < nbFloors; i++){
    elevators[i] = [];
}
for (var i = 0; i < nbElevators; i++) {
    var inputs = readline().split(' ');
    var elevatorFloor = parseInt(inputs[0]); // floor on which this elevator is found
    var elevatorPos = parseInt(inputs[1]); // position of the elevator on its floor
    elevators[elevatorFloor].push(elevatorPos);
}
var lastBlock;
//var lastBlockF;

// game loop
while (true) {
    var action = 'WAIT';
    var inputs = readline().split(' ');
    var cloneFloor = parseInt(inputs[0]); // floor of the leading clone
    var clonePos = parseInt(inputs[1]); // position of the leading clone on its floor
    var direction = inputs[2]; // direction of the leading clone: LEFT or RIGHT
    if(cloneFloor === exitFloor){
        if ((clonePos > exitPos && direction === 'RIGHT') || (clonePos < exitPos && direction === 'LEFT') && (clonePos !== lastBlock)){
            action = 'BLOCK';
            lastBlock = clonePos;
        }
    }else if (direction !== 'NONE'){
        let minDist = 99999999;
        let closestEle;
        for (let i = 0; i < elevators[cloneFloor].length; i++){
            if (Math.abs(elevators[cloneFloor][i] - clonePos) < minDist){
                minDist = Math.abs(elevators[cloneFloor][i] - clonePos);
                closestEle = elevators[cloneFloor][i];
            }
        }
        if ((clonePos > closestEle && direction === 'RIGHT') || (clonePos < closestEle && direction === 'LEFT') && (clonePos !== lastBlock)){
            action = 'BLOCK';
            lastBlock = clonePos;
        }
        
    }

    // Write an action using print()
    // To debug: printErr('Debug messages...');

    print(action); // action: WAIT or BLOCK
}