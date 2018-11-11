public class Main {
  public static void main(String... args) {
    if (args.length < 2) {
      System.out.println("need 2 arguments.");
      System.exit(1);
    }

    int x = Integer.parseInt(args[0]);
    int y = Integer.parseInt(args[1]);
    int sum = x + y;

    System.out.println("sum:" + sum);
  }
}

