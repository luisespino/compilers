let x;

document.getElementById("x").addEventListener("keydown", function(e) {
    if (e.key === "Enter") {
        x = document.getElementById("x").value;
        var container = document.getElementById("mynetwork");
        var DOTstring = PEG.parse(x);
        var parsedData = vis.parseDOTNetwork(DOTstring);
        var data = {
            nodes: parsedData.nodes,
            edges: parsedData.edges
        }
        var options = {
            nodes: {
                widthConstraint: 20,
            },        
            layout: {
                hierarchical: {
                    levelSeparation: 60,
                    nodeSpacing: 80,
                    parentCentralization: true,
                    direction: 'UD',        // UD, DU, LR, RL
                    sortMethod: 'directed',  // hubsize, directed
                    shakeTowards: 'roots'  // roots, leaves                        
                },
            },                        
        };
        var network = new vis.Network(container, data, options);
    }
});


