/*
Given a string s consists of some words separated by spaces, return the length of the last word in the string. 
If the last word does not exist, return 0.
A word is a maximal substring consisting of non-space characters only.
*/


class Solution {
public:
    int lengthOfLastWord(string s) {
        int word_length = 0;
        if (s.size() == 0) return 0;
        for (int i = s.size() - 1; i >= 0; i--) {
            if (s[i] == ' ') {
                if (word_length == 0) {
                    continue;
                } else {
                    break;
                }
            }
            word_length++;
        }
        return word_length;
    }
};
