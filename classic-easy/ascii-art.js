/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var L = parseInt(readline());
var H = parseInt(readline());
var T = readline();
var rows = [];
for (var i = 0; i < H; i++) {
    var ROW = readline();
    rows[i] = ROW;
}
var chars = T.toLowerCase();
//print(chars);
var nums = []
for(let i = 0; i< T.length; i++){
    nums[i] = chars.charCodeAt(i);
}
//print(nums);
var string = '';

for(let i=0; i<H; i++){
    string = '';
    for(let j = 0; j< nums.length; j++){
        let num = (nums[j] - 97) * L;
        if(nums[j] > 122 || nums[j] < 97){
            for(let k = 0; k<L; k++){
                string += rows[i].charAt(26*4 + k);
            }
        }else{
            for(let k = 0; k<L; k++){
                string += rows[i].charAt(num + k);
            }
        }
    }
    print(string);
}