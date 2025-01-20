
class Pos {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
}


const ball = $createActor("ball", {
    pos: { x: 0, y: 0 },
    vel: { x: 1, y: 1 },
    acc: { x: 1, y: 1 },
});


ball.move = function () {
    console.log("moving", JSON.stringify(this.$getState()));
    this.pos = { x: this.pos.x + this.vel.x, y: this.pos.y + this.vel.y };
    this.vel = { x: this.vel.x + this.acc.x, y: this.vel.y + this.acc.y };
};


ball.$onProperty("pos", function (value) {
    console.log("pos changed", JSON.stringify(value));
});

ball.$onProperty("vel", function (value) {
    console.log("vel changed", JSON.stringify(value));
});

ball.$onProperty("acc", function (value) {
    console.log("acc changed", JSON.stringify(value));
});

function main() {
    console.log("running");
    for (let i = 0; i < 10; i++) {
        ball.move();
    }
    console.log("done", JSON.stringify(ball.$getState()));
}
