const one = async () => {
    // simulate a workload
    sleep(Math.floor(Math.random() * Math.floor(2000)));
    return 1;
};

const two = async () => {
    // simulate a workload
    sleep(Math.floor(Math.random() * Math.floor(1000)));
    sleep(Math.floor(Math.random() * Math.floor(1000)));
    return 2;
};

const r = await Promise.race(one(), two());
console.log(r);