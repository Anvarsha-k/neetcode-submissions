class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        mydict = {}
        if len(s) == len(t):
            for ch in s:
                if ch in mydict:
                    mydict[ch] += 1
                else:
                    mydict[ch] = 1
            
            for i in t:
                if i not in mydict or mydict[i] == 0:
                    return False
                mydict[i] -= 1
            return True
        else:
            return False
            
        
        


            
        