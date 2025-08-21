const demo = $createService("test.Demo", {})

demo.dynamicSignal = function(arg0, arg1) {
    console.log(`Dynamic signal called with args: ${arg0}, ${arg1}`);
    this.emit('signal', arg0, arg1);
}
demo.constSignal = function() {
    console.log("Const signal called");
    this.emit('signal', "arg0", "arg1");
};

demo.on('signal', function(arg0, arg1) {
    console.log(`Signal received with args: ${arg0}, ${arg1}`);
});

function main() {
    let v = 0
    setInterval(function() {
        console.log(`----`);
        v++
        if (v%2===0) {
            demo.dynamicSignal("dynamicArg0", "dynamicArg1");
        } else {
            demo.constSignal();
        }
    }, 1000);
}
