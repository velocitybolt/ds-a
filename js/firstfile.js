/*

Javascript Intro Stuff

*/

let a = 333;
let b = 2;

const addNums = (a, b) => {
    if (typeof (a) != 'number' || typeof (b) != 'number') {
        return;
    }
    const sum = a + b;
    if (sum > 30) {
        console.log(sum, "is greater than 30!")
    } else {
        console.log(sum, "is less than 30!")
    }
}
addNums(a, b);