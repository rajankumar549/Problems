/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.

*/

/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
    var d = {}
    var keys =[]
    var res = []
    strs.forEach(function(val){
    var word = val;
    var key = val.split("").sort().join("")
        if (!d[key]){
            keys.push(key)
            d[key]=[]
        }
        d[key].push(word)
    })
    console.log(d)
    keys.forEach(function(val){
        res.push(d[val])
    })
    return res
};
