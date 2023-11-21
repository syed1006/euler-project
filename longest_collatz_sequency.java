import java.math.BigInteger;
import java.util.HashMap;

public class longest_collatz_sequency {

    public static int findCollatzChainUntil(BigInteger num, HashMap<BigInteger, Integer>collMap){
        int chainCount = 1;
        if(collMap.containsKey((num))){
            return collMap.get((num));
        }
        while(!num.equals(BigInteger.ONE)){
            // System.out.println(num);
            if(collMap.containsKey((num))){
                chainCount += collMap.get(num);
                break;
            }
            if (num.testBit(0)) {  // If n is odd
                num = num.multiply(BigInteger.valueOf(3)).add(BigInteger.ONE);
            } else {  // If n is even
                num = num.divide(BigInteger.valueOf(2));
            }
            chainCount ++;
        }
        return chainCount;
    }
    
    public static void main(String[] args){
        int maxCount = 0;
        BigInteger maxCountNumber = new BigInteger("0");
        BigInteger decrementBy = new BigInteger("1");
        BigInteger number = new BigInteger("999999");
        HashMap<BigInteger, Integer> collMap = new HashMap<>();
        while(!number.equals(BigInteger.ONE)){
            int chainCount = findCollatzChainUntil(number, collMap);
            collMap.put((number), chainCount);
            // System.out.println(number);
            if(chainCount > maxCount){
                maxCount = chainCount;
                maxCountNumber = number;
                System.out.println("here");
                System.out.println((maxCount));
                System.out.println((number));
            }
            number = number.subtract(decrementBy);
        }
        System.out.println((maxCountNumber));
    }
}
