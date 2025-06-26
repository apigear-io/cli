// Simple function example showing basic simulation script structure
function main() {
    console.log("main called");
    
    // This example demonstrates that simulations can be simple functions
    // without services if no state management is needed
    
    // Exit simulation
    if (typeof $quit === 'function') {
        $quit();
    }
}