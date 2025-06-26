const $createService = function(objectId, properties) {
    const service = $createBareService(objectId, properties);
    
    let serviceProxy;
    serviceProxy = new Proxy(service, {
        get: function(target, property, receiver) {
            // Access to raw service object
            if (property === "$") {
                return target;
            }
            
            // Method access - return bound method with proxy as context
            if (service.hasMethod(property)) {
                const method = service.getMethod(property);
                return function(...args) {
                    return method.call(serviceProxy, ...args);
                };
            }
            
            // Property access
            if (service.hasProperty(property)) {
                return service.getProperty(property);
            }
            
            // Convenience methods
            if (property === 'on') {
                return function(event, callback) {
                    if (service.hasProperty(event)) {
                        return service.onProperty(event, callback);
                    }
                    return service.onSignal(event, callback);
                };
            }
            
            if (property === 'emit') {
                return function(signal, ...args) {
                    return service.emitSignal(signal, ...args);
                };
            }
            
            // Built-in service methods and properties
            if (property in target) {
                const value = target[property];
                if (typeof value === 'function') {
                    return value.bind(target);
                }
                return value;
            }
            
            // Undefined property - provide helpful error
            if (typeof property === 'string' && !property.startsWith('_')) {
                console.warn(`Property '${property}' not found on service '${service.objectId()}'. Available properties: [${Object.keys(service.getProperties()).join(', ')}]`);
            }
            
            return undefined;
        },
        
        set: function(target, property, value, receiver) {
            // Don't intercept internal properties
            if (property === "$" || property.startsWith('_')) {
                return Reflect.set(target, property, value, receiver);
            }
            
            // Function assignment = method registration
            if (typeof value === "function") {
                // Wrap function to provide 'this' context as proxy
                const wrappedFn = function(...args) {
                    return value.call(serviceProxy, ...args);
                };
                service.onMethod(property, wrappedFn);
                return true;
            }
            
            // Property assignment
            if (service.hasProperty(property)) {
                service.setProperty(property, value);
                return true;
            }
            
            // New property - add to service properties
            service.setProperty(property, value);
            return true;
        }
    });
    
    return serviceProxy;
}