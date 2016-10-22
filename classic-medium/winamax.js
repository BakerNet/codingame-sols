/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
var p1 = []
var n = parseInt(readline()); // the number of cards for player 1
for (var i = 0; i < n; i++) {
    var cardp1 = readline(); // the n cards of player 1
    p1.push(cardp1);
}
printErr(p1);
var p2 = []
var m = parseInt(readline()); // the number of cards for player 2
for (var i = 0; i < m; i++) {
    var cardp2 = readline(); // the m cards of player 2
    p2.push(cardp2);
}
printErr(p2);

function getVal(move){
    let val;
    if (move.length === 3){
        return 10;
    }else{
        val = move.charAt(0);
    }
    switch(val){
        case 'A':
            return 14;
        case 'K':
            return 13;
        case 'Q':
            return 12;
        case 'J':
            return 11;
        default:
            return parseInt(val);
    }
}
function getWinner(p1move, p2move, p1val, p2val, p1played, p2played){
    p1played.push(p1move);
    p2played.push(p2move);
    if (p1val === p2val){
        for (let i = 0; i < 3; i++){
            if (p1.length === 0 || p2.length === 0){
                print('PAT');
                return false;
            }
            p1played.push(p1.shift());
            p2played.push(p2.shift());
        }
        return getWinner(p1move = p1.shift(), p2move = p2.shift(), getVal(p1move), getVal(p2move), p1played, p2played);
    }else if(p1val < p2val){
        return "p2"
    }else{
        return "p1";
    }
}

var round = 0;
// Write an action using print()
// To debug: printErr('Debug messages...');
while(p1.length > 0 && p2.length > 0){
    let p1move = p1.shift();
    let p2move = p2.shift();
    let p1val = getVal(p1move);
    let p2val = getVal(p2move);
    let p1played = [];
    let p2played = [];
    let winner = getWinner(p1move, p2move, p1val, p2val, p1played, p2played);
    if (winner){
        if(winner === 'p1'){
            for (let i = 0; i < p1played.length; i++){
                p1.push(p1played[i]);
            }
            for (let i = 0; i< p2played.length; i++){
                p1.push(p2played[i]);
            }
        }else{
            for (let i = 0; i < p1played.length; i++){
                p2.push(p1played[i]);
            }
            for (let i = 0; i< p2played.length; i++){
                p2.push(p2played[i]);
            }
        }
    }else{
        exit(0);
    }
    round++;
}

if (p1.length === 0){
    print(`2 ${round}`);
}else{
    print(`1 ${round}`);
}