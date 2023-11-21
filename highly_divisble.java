class Solution{

    public static int findDivisorsNumber(int number){
        int divisor = 1;
        int limit = number;
        for(int i = 1; i < limit; i++){
            if(number % i == 0){
                limit = number/i;
                if(limit != i){
                    divisor++;
                }
                divisor++;
            }
        }
        return divisor;
    }

    public static void main (String[] args){
        int n = 1;
        int divisorsNum = 500;
        while(true){
            int number = n * (n + 1) / 2;
            int divisors = findDivisorsNumber(number);
            System.out.println(n);
            System.out.println(divisors);
            if(divisors > divisorsNum){
                System.out.println(number);
                return;
            }
            n++;
        }

    }
}