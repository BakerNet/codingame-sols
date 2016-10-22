/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var n = parseInt(readline()); // the number of temperatures to analyse
var temps = readline(); // the n temperatures expressed as integers ranging from -273 to 5526
temps = temps.split(' ');
var low = 9999999999999999999999;
// Write an action using print()
// To debug: printErr('Debug messages...');
for(let i = 0; i<n; i++){
    let x = parseInt(temps[i]);
    if (Math.abs(x) < Math.abs(low)){
        low = x;   
    }else if(Math.abs(x) === Math.abs(low)){
        if (x > low){
            low = x;
        }
    }
}
if(low === 9999999999999999999999){
    print(0);
}else{
    print(low);
}