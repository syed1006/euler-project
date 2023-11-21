class Solution{

    public double findRoutes(int n){
        double routes = 1;

        for(int i = 1; i <= n; i++){
            routes = (routes * (n + i))  / i;
        }

        return routes;
    }

    public static void main(String[] args){
        Solution sol = new Solution();
        double totalRoutes = sol.findRoutes(20);
        System.out.println((totalRoutes));
    }
}