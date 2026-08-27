class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        mydict = {}
        for word in strs:
            key ="".join(sorted(word))
            if key not in mydict:
                mydict[key] = []
            mydict[key].append(word)
        
        return list(mydict.values())

            