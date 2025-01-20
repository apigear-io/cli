export interface IWorld {
    createActor(actorId: string, state: { [key: string]: any }): IActor;
    getActor(actorId: string): IActor | null;
}

export interface IActor {
    setMethod(key: string, method: (...args: any[]) => any): void;
    callMethod(key: string, ...args: any[]): any;
    getProperty(key: string): any;
    setProperty(key: string, value: any): void;
    onPropertyChanged(property: string, listener: (...args: any[]) => void): void;
    emitPropertyChanged(property: string, value: any): void;
    getState(): { [key: string]: any };
    setState(properties: { [key: string]: any }): void;
    onSignal(signal: string, listener: (...args: any[]) => void): void;
    emitSignal(signal: string, ...args: any[]): void;
}

export class ActorProxy {
    constructor(actorId: string, state: { [key: string]: any })
    $setMethod(key: string, method: (...args: any[]) => any): void;
    $callMethod(key: string, ...args: any[]): any;
    $getProperty(key: string): any;
    $setProperty(key: string, value: any): void;
    $onProperty(property: string, listener: (...args: any[]) => void): void;
    $emitProperty(property: string, value: any): void;
    $getState(): { [key: string]: any };
    $setState(properties: { [key: string]: any }): void;
    $onSignal(signal: string, listener: (...args: any[]) => void): void;
    $emitSignal(signal: string, ...args: any[]): void;
}

declare global {
    var $world: IWorld;
    function $createActor(name: string, state: { [key: string]: any }): ActorProxy;
}