import java.util.HashMap;

public class Main {

    private static HashMap<char,int> m = new HashMap<char, int>();


    public static void main(String[] args) {

        m.put('I',1);
        m.put('V',5);
        m.put('X',10);
        m.put('L',50);
        m.put('C',100);
        m.put('D',500);
        m.put('M',1000);

        String value1 = args[0];
    }

    public int parseRomanNums(char[] nums) {

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