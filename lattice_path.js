function countRoutes(n) {
    let result = 1;

    for (let i = 1; i <= n; i++) {
        result *= (n + i) / i;
    }

    return result;
}

// Example usage:
const n = 20; // Replace with the desired grid size
const numberOfRoutes = countRoutes(n);
console.log(`Number of routes in a ${n}x${n} grid: ${numberOfRoutes}`);