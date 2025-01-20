const actor = $world.createActor('counter3', { count: 0 });

actor.setMethod('increment', function () {
    this.setProperty('count', this.getProperty('count') + 1);
});

actor.setMethod('decrement', function () {
    this.setProperty('count', this.getProperty('count') - 1);
});

actor.onProperty('count', function (value) {
    console.log('count changed', value);
});


const main = function() {
    for (let i = 0; i < 5; i++) {
        actor.callMethod('increment');
    }
    for (let i = 0; i < 3; i++) {
        actor.callMethod('decrement');
    }
    return actor.getProperty('count');
}

for (let i = 0; i < 5; i++) {
    actor.callMethod('increment');
}
for (let i = 0; i < 3; i++) {
    actor.callMethod('decrement');
}
actor.getProperty('count');
