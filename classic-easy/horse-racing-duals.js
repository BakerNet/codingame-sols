/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var N = parseInt(readline());
var arry = []
for (var i = 0; i < N; i++) {
    var pi = parseInt(readline());
    arry.push(pi);
}
function comp(a, b){
    return a - b;
}
arry.sort(comp);
var mindiff = 999999999999999999999;
for(let i = 0; i < N; i++){
    if ((arry[i+1] - arry[i]) < mindiff){
        mindiff = arry[i+1] - arry[i];
    }
}

// Write an action using print()
// To debug: printErr('Debug messages...');

print(mindiff);