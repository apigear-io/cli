function $createProxy(actor, that) {
    const proxy = new Proxy(that || actor, {
        get(target, prop) {
            if (prop.startsWith('$')) {
                return actor[prop.slice(1)];
            }
            if (actor.hasMethod(prop)) {
                return actor.getMethod(prop);
            }
            if (actor.hasProperty(prop)) {
                return actor.getProperty(prop);
            }
            if (that) {
                return Reflect.get(that, prop);
            }
        },
        set(target, prop, value) {
            if (value instanceof Function) {
                const bound = value.bind(proxy);
                return actor.setMethod(prop, bound);
            }
            actor.setProperty(prop, value);
            if (that) {
                return Reflect.set(that, prop, value);
            }
        },
        apply(target, thisArg, args) {
            console.log('apply', args);
            if (args.length > 0 && actor.hasMethod(args[0])) {
                return actor.callMethod(args[0], ...args.slice(1));
            }
            return Reflect.apply(actor, thisArg, args);
        }
    });
    return proxy;
}

function $createActor(name, state) {
    const actor = $world.createActor(name, state);
    console.log('actor', actor.id());
    return $createProxy(actor);
}

class ProxyActor {
    constructor(name, state) {
        this.$actor = $world.createActor(name, state);
        return $createProxy(this.$actor, this);
    }
}