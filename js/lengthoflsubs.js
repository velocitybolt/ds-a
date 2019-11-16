/*

The length of the longest substring
abcabcbb -- > 3
bbbbbbbbbbbb -> 1

*/

function lengthOfLongestSubstring(s) {
    // create key/val dict to store char and where it was last seen
    let lastChar = {}
    // var to store start index of substring
    let startIndex = 0;
    let maxLen = 0;

    /* The approach we are following here is the sliding window 
    The plan is to basically have two pointers the startindex and endindex which represent the dimensions of the window
    then we loop through input and check if we have seen this character already
    This is done by checking whether the current keys value at the position of the end char is located within the window.
    It will be done by checking if its >= startIndex 
    */
    for (let i = 0; i < s.length; i++) {
        let endChar = s[i];

        if (lastChar[endChar] >= startIndex) {
            // this means duplicate has been found
            // reset the startindex to be one after the last location it was seen
            startIndex = lastChar[endChar] + 1;
        }
        // this means the char is new store its key and location in the dict 
        lastChar[endChar] = i;
        // check whether this new substring is the longest substring
        // i - Startindex is the length of current string because startIndex is start and i is the index location of the current ending char 
        // but we also add one because arrays are 0 indexed. 
        maxLen = Math.max(maxLen, i - startIndex + 1);
    }
    return maxLen;

}
console.log(lengthOfLongestSubstring("abcabcbb"));

console.log(lengthOfLongestSubstring("bbbbbbbbbbbb"));

console.log(lengthOfLongestSubstring("pwwkew"));