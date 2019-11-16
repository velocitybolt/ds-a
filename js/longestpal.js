/*

Longest Palindrome

cbbd --> bb
abba --> abba
a --> a

*/


function longestPalindrome(s) {
    /*
        The purpose of this function is to return the longest palindrone in the given string
        One we can go about this is splitting the given substring into two and matching both sides.
        A way we can approach this is using pointers to mark the left and right side of the string.

    */

    let startIndex = 0;
    // set to one at minimum because it will contain at least 1 char
    let maxLength = 1;

    function expandAroundMiddle(left, right) {
        // ensure its within bounds and the opposite sides have the same char meaning palindrome
        while (left >= 0 && right <= s.length && s[right] === s[left]) {
            // calculate current pal length
            // this is calculated the indexes of right and left. 
            // [cbbd] ==> first b is at index 1 , second b is and index 2 
            //==> 2 - 1 = 1 + 1 (arrays are 0 indexed) ==> length of currpal
            const currPal = right - left + 1;
            // if the current palindrome is longer than current stored max
            if (currPal > maxLength) {
                // update the maxlength to be currPal 
                maxLength = currPal;
                // set the start index to be left because the left is where the previous palindrome starts
                startIndex = left;
            }
            // decrement the left side
            left -= 1;
            // increment the right side
            right += 1;
        }

    }
    for (let i = 0; i < s.length; i++) {
        // cover odd length strs
        expandAroundMiddle(i - 1, i + 1);
        // cover even length strs 
        expandAroundMiddle(i, i + 1);
    }
    return s.slice(startIndex, startIndex + maxLength);


}

console.log(longestPalindrome("cbbd"));
console.log(longestPalindrome("abba"));
console.log(longestPalindrome("c"));