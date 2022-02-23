package fr.imtne.ilog;

public class Display {
	
	static {
	    System.loadLibrary("HelloNative");
	}
	
	private native void displayLine(int iLine, String message);
	
	public static void main(String[] args) {
		Display hj = new Display();
		hj.displayLine( 1, "Helo JNI !" );
	}

}
