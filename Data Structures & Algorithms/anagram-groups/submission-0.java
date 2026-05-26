class Solution {

    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> result = new HashMap<>();

        for (String str: strs){
            char[] ch = str.toCharArray();
            Arrays.sort(ch);
            String sorted_str = Arrays.toString(ch);
            result.putIfAbsent(sorted_str, new ArrayList<>());
            result.get(sorted_str).add(str);
        }
        return new ArrayList<>(result.values());
    }
    }