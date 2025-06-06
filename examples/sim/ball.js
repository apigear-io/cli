
const ball = $createService("ball", {
    pos: { x: 0, y: 0 },
    vel: { x: 1, y: 1 },
    acc: { x: 1, y: 1 },
});


ball.move = () => {
    var acc = ball.acc;
    var vel = ball.vel;
    var pos = ball.pos;
    var newPos = { x: pos.x + vel.x, y: pos.y + vel.y };
    var newVel = { x: vel.x + acc.x, y: vel.y + acc.y };
    ball.pos += newPos;
    ball.vel = newVel;
};

    
ball.$.onProperty("pos", function (value) {
    console.log("pos changed", JSON.stringify(value));
});

ball.$.onProperty("vel", function (value) {
    console.log("vel changed", JSON.stringify(value));
});

ball.$.onProperty("acc", function (value) {
    console.log("acc changed", JSON.stringify(value));
});

function main() {
    console.log("start");
    for (let i = 0; i < 10; i++) {
        ball.move()
    }
    console.log("finish", JSON.stringify(ball.$.getProperties()));
    $quit();
}
