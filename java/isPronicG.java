public class PronicNumber {
	public static boolean isHeteromecic(int n) {
		double square = Math.sqrt(n);
		double intBelow = Math.floor(square);
		double intAbove = Math.ceil(square);
		return n == (intAbove * intBelow);
	}
}