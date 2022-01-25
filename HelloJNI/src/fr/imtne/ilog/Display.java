package fr.imtne.ilog;

public class Display {
	
	static {
	    System.loadLibrary("HelloNative");
	}
	
	private native void display(String message);
	
	public static void main(String[] args) {
		Display hj = new Display();
		hj.display( "HelloJNI" );
	}

}
