const actor = $world.createActor('counter', { count: 0 });

actor.setMethod('increment', function () {
    console.log('increment');
    actor.setProperty('count', this.getProperty('count') + 1);
});

actor.setMethod('decrement', function () {
    console.log('decrement');
    actor.setProperty('count', this.getProperty('count') - 1);
});

actor.onProperty('count', function (value) {
    console.log('count changed', value);
});


function main() {
    // setInterval(function() {
    //     for (let i = 0; i < 3; i++) {
    //         actor.callMethod('increment');
    //     }
    //     for (let i = 0; i < 3; i++) {
    //         actor.callMethod('decrement');
    //     }    
    //     console.log('final count', actor.getProperty('count'));
    // }, 1000);
}
