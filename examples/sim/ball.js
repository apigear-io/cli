
// Ball physics simulation using natural API
const ball = $createService("ball", {
    pos: { x: 0, y: 0 },
    vel: { x: 1, y: 1 },
    acc: { x: 1, y: 1 },
});

// Define move method using natural API with 'this'
ball.move = function() {
    const acc = this.acc;
    const vel = this.vel;
    const pos = this.pos;
    
    // Calculate new position and velocity
    const newPos = { x: pos.x + vel.x, y: pos.y + vel.y };
    const newVel = { x: vel.x + acc.x, y: vel.y + acc.y };
    
    // Update properties using natural assignment
    this.pos = newPos;  // Fixed: was using += incorrectly
    this.vel = newVel;
    
    // Emit movement signal
    this.emit('moved', newPos);
};

// Reset method
ball.reset = function() {
    this.pos = { x: 0, y: 0 };
    this.vel = { x: 1, y: 1 };
    this.emit('reset');
}

// Monitor property changes using natural API
ball.on("pos", function (value) {
    console.log("Position changed:", JSON.stringify(value));
});

ball.on("vel", function (value) {
    console.log("Velocity changed:", JSON.stringify(value));
});

ball.on("acc", function (value) {
    console.log("Acceleration changed:", JSON.stringify(value));
});

// Listen to custom signals
ball.on('moved', function(newPos) {
    console.log(`Ball moved to: (${newPos.x}, ${newPos.y})`);
});

function main() {
    console.log("=== Ball Physics Simulation ===");
    console.log("Initial state:", JSON.stringify(ball.$.getProperties()));
    
    // Run simulation
    for (let i = 0; i < 5; i++) {
        console.log(`\nStep ${i + 1}:`);
        ball.move();
    }
    
    console.log("\nFinal state:", JSON.stringify(ball.$.getProperties()));
    
    // Demonstrate reset
    console.log("\nResetting ball...");
    ball.reset();
    console.log("State after reset:", JSON.stringify(ball.$.getProperties()));
    
    if (typeof $quit === 'function') {
        $quit();
    }
}
