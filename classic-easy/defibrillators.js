/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var LON = parseFloat(readline().replace(/,/g, '.'));
var LAT = parseFloat(readline().replace(/,/g, '.'));
//print(LON);
//print(LAT);
var N = parseInt(readline());
var defibs = [];
var dist = 99999999999.9;
var answer = -1;
for (var i = 0; i < N; i++) {
    var DEFIB = readline();
    defibs[i] = DEFIB.replace(/,/g, '.').split(';');
    defibs[i][4] = parseFloat(defibs[i][4]);
    defibs[i][5] = parseFloat(defibs[i][5]);
    //print(defibs[i]);
    let x = (defibs[i][4] - LON) * Math.cos((defibs[i][5] + LAT)/2);
    //print(x);
    let y = (defibs[i][5] - LAT);
    //print(y);
    let d = Math.sqrt((x*x) + (y*y)) * 6371;
    //print(d);
    if(d < dist){
        dist = d;
        //print(d);
        answer = i;
        //print(i);
    }
}

// Write an action using print()
// To debug: printErr('Debug messages...');

print(defibs[answer][1]);