import java.util.HashMap;

public class Main {

    public static void main(String[] args) {

        String value1 = args[0];
        char[] asChars = value1.toCharArray();

        Parser testies = new Parser();

        System.out.println(testies.parseRomanNums(asChars));
    }
}

class Parser {

    public int parseRomanNums(char[] nums) {

        HashMap<Character, Integer> m = new HashMap<Character, Integer>();

        m.put('I', 1);
        m.put('V', 5);
        m.put('X', 10);
        m.put('L', 50);
        m.put('C', 100);
        m.put('D', 500);
        m.put('M', 1000);

        int tot = 0;

        if (nums.length == 1) {
            tot += m.get(nums[0]);
        } else {
            for (int i = nums.length-1; i >= 0; i-- ) {
                if ( i != 0 ) {
                    if (m.get(nums[i]) > m.get(nums[i-1])) {
                        tot += m.get(nums[i]);
                        tot -= m.get(nums[i-1]);
                        i--;
                    } else {
                        tot += m.get(nums[i]);
                    }
                } else {
                    if (m.get(nums[0]) < m.get(nums[1])) {
                        tot -= m.get(nums[0]);
                    } else {
                        tot += m.get(nums[0]);
                    }
                }
            }
        }
        return tot;
    }
}