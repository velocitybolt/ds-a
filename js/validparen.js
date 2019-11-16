/*
    Given a string containing various parentheses chars

    Determine whether the given input string has balacned parentheses

    Input : "()" ==> True
    Input : "()[]{}" ==> True

*/

function validParen(s) {
    
    // Here we have two things a stack to checking paren balancing
    // And a javascript dictionary to check matching pairs
    let stack = [];
    const pairs = { "{":"}", "[":"]", "(":")" };

    for (let i = 0; i < s.length; i++) {
        // var to rep curr char for readability
        let currBrace = s[i]
        // loop and check if current brace is opening 
        if (pairs[currBrace]) {
            // then push onto the stack
            stack.push(currBrace);
        }
        // if the brace of the top of the stack is a right brace
        else if (pairs[stack.pop()] !== currBrace) {
            // return false because that is unbalanced ([{
            return false
        }
    }
    return stack.length === 0;

    
}

console.log(validParen("()[]{}"));