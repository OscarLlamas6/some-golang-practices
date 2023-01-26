
const longRunningTask = async () => {
    // simulate a workload
    sleep(3000);
    return Math.floor(Math.random() * Math.floor(100));
};

const r = await longRunningTask();
console.log(r);