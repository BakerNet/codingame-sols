/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var inputs = readline().split(' ');
var N = parseInt(inputs[0]); // the total number of nodes in the level, including the gateways
//print(N);
var L = parseInt(inputs[1]); // the number of links
//print(L);
var E = parseInt(inputs[2]); // the number of exit gateways
var nodes = new Array(N);
for(let i = 0; i< N; i++){
    nodes[i] = {isGateway: false, visited: false, links: []};
}

for (var i = 0; i < L; i++) {
    var inputs = readline().split(' ');
    var N1 = parseInt(inputs[0]); // N1 and N2 defines a link between these nodes
    var N2 = parseInt(inputs[1]);
    nodes[N1].links.push(N2);
    nodes[N2].links.push(N1);
    //print(JSON.stringify(nodes[N1]));
}
for (var i = 0; i < E; i++) {
    var EI = parseInt(readline()); // the index of a gateway node
    nodes[EI].isGateway = true;
    //print(EI + JSON.stringify(nodes[EI]));
}

Object.prototype.clone = function() {
    var newObj = (this instanceof Array) ? [] : {};
    for (i in this) {
        if (i == 'clone') 
            continue;
        if (this[i] && typeof this[i] == "object") {
            newObj[i] = this[i].clone();
        } 
        else 
            newObj[i] = this[i]
    } return newObj;
};

// game loop
while (true) {
    var SI = parseInt(readline()); // The index of the node on which the Skynet agent is positioned this turn

    // Write an action using print()
    // To debug: printErr('Debug messages...');
    let lnodes = nodes.clone();
    let queue = [];
    let kill = '';
    
    queue.push(SI);
    
    while(queue.length > 0){
        let currI = queue.shift();
        let currN = lnodes[currI];
        currN.visited = true;
        for(let i = 0; i<currN.links.length; i++){
            let N = currN.links[i];
            if(!lnodes[N].visited){
                if (lnodes[N].isGateway){
                    kill += currI + ' ' + N;
                    queue = [];
                    break;
                }else{
                    queue.push(N);
                }
            }
        }
    }

    // Example: 0 1 are the indices of the nodes you wish to sever the link between
    print(kill);
}