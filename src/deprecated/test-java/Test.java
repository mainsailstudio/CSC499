import java.util.List;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

public class Test{

    private static void getSubsets(List<String> superSet, int k, int idx, Set<String> current,List<List<String>> solution) {
        //successful stop clause
        if (current.size() == k) {
            solution.add(new ArrayList<>(current));
            return;
        }

        //unsuccessful stop clause
        if (idx == superSet.size()) return;

        String x = superSet.get(idx);
        current.add(x);

        //"guess" x is in the subset
        getSubsets(superSet, k, idx+1, current, solution);
        current.remove(x);

        //"guess" x is not in the subset
        getSubsets(superSet, k, idx+1, current, solution);
    }

    public static List<List<String>> getSubsets(List<String> superSet, int k) {
        List<List<String>> res = new ArrayList<>();
        getSubsets(superSet, k, 0, new HashSet<String>(), res);
        return res;
    }

    //Generating permutation using Heap Algorithm
    public static void heapPermutation(List<String> a, int size, int n, Set<String> current, List<List<String>> solution)
    {
        // if size becomes 1 then prints the obtained
        // permutation
        if (size == 1) {
            solution.add(new ArrayList<>(a));
        }

        for (int i=0; i<size; i++)
        {
            heapPermutation(a, size-1, n, current, solution);

            // if size is odd, swap first and last
            // element
            if (size % 2 == 1)
            {
                String temp = a.get(0);
                a.set(0, a.get(size-1));
                a.set(size-1, temp);
            }

            // If size is even, swap ith and last
            // element
            else
            {
                String temp = a.get(i);
                a.set(i, a.get(size-1));
                a.set(size-1, temp);
            }
        }
    }

    public static List<List<String>> getPerms(List<List<String>> superSet) {
        List<List<String>> res = new ArrayList<>();
        for(int i = 0; i < superSet.size(); i++){
            List<String> list = new ArrayList<>();
            list = superSet.get(i);
            heapPermutation(list, list.size(), list.size(), new HashSet<String>(), res);
        }
        return res;
    }

    public static void main(String[] args){
        List<String> superSet = new ArrayList<>();
        superSet.add("1");
        superSet.add("2");
        superSet.add("3");
        superSet.add("4");
        superSet.add("5");
        superSet.add("6");
        superSet.add("7");
        superSet.add("8");
        superSet.add("9");
        superSet.add("10");
        long startTime = System.nanoTime();
        List<List<String>> perms = getSubsets(superSet,4);
        System.out.println(perms);
        System.out.println("Total amount of initial is " + perms.size());
        List<List<String>> fullPerm = new ArrayList<>();
        fullPerm = getPerms(perms);
        System.out.println(fullPerm);
        System.out.println("Total amount of final is " + fullPerm.size());
        long endTime = System.nanoTime();
        long duration = (endTime - startTime);  // nanoseconds
       // long duration = ((endTime - startTime)/1000000);  // milliseconds
        System.out.println("Time took was " + duration + " nanoseconds");


    }
}