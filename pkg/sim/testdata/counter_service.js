const service = world.getService("counter", {count: 0 });
service.onPropertyChange("count", function (count) {
    console.log("count changed", count);
})
service.onMethod("increment", function () {
    const count = this.getProperty("count");
    this.setProperty("count", count + 1);
});