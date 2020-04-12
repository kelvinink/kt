class Singleton{
    // Private instance
    private static Singleton instance = new Singleton();

    // Private default constructor
    private Singleton(){}

    // Public getInstance method
    public static Singleton getInstance(){
        return instance;
    }
}