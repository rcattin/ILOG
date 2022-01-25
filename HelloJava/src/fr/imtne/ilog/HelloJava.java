package fr.imtne.ilog;

public class HelloJava {
    public static void main(String[] args) {
    	Message hello = new Message();
		hello.display();
		hello.display(args[0]);
    }
}
