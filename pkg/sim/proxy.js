const $createService = function(objectId, properties) {
    const service = $createBareService(objectId, properties);
    return new Proxy(service, {
        get: function(target, property, receiver) {
            if (property === "$") {
                return target;
            }
            if (service.hasMethod(property)) {
                return service.getMethod(property)
            }
            if (service.hasProperty(property)) {
                return service.getProperty(property);
            }
            return Reflect.get(target, property, receiver);
        },
        set: function(target, property, value, receiver) {
            target[property] = value;
            if (property === "$") {
                return
            }
            if (typeof value === "function") {
                return service.onMethod(property, value);
            }
            if (service.hasProperty(property)) {
                return service.setProperty(property, value);
            }
            return Reflect.set(target, property, value, receiver);
        }
    });
}