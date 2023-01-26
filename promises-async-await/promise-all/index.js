const longRunningTask = async () => {
    // simulate a workload
    sleep(3000);
    return Math.floor(Math.random() * Math.floor(100));
};

const [a, b, c] = await Promise.all(longRunningTask(), 
                                    longRunningTask(), 
                                    longRunningTask());
console.log(a, b, c);