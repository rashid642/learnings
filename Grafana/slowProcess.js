function getRandomIntLessThan(x) {
    return Math.floor(Math.random() * x);
}

const timeRequired = [1000, 2000, 5000, 3000, 4000]
const slowProcess = () => {
    return new Promise((resolve, reject) => {
        const size = timeRequired.length
        const randomNumber = getRandomIntLessThan(2*size)
        if(randomNumber >= size) {
            reject("Some error happened")
        }
        console.log("Time to wait :", timeRequired[randomNumber])
        setTimeout(() => {
            resolve("Returning back");
        }, timeRequired[randomNumber])
    })
}

module.exports = slowProcess